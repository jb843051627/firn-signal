package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jb843051627/firn-signal/internal/model"
)

func (s *Store) Save(ctx context.Context, kind, id string, value any) error {
	if kind == "" || id == "" {
		return model.ErrInvalidInput
	}
	payload, err := encode(value)
	if err != nil {
		return err
	}
	now := time.Now().UTC().Format(time.RFC3339Nano)
	_, err = s.db.ExecContext(ctx, `INSERT INTO records(kind,id,payload,created_at,updated_at) VALUES(?,?,?,?,?) ON CONFLICT(kind,id) DO UPDATE SET payload=excluded.payload, updated_at=excluded.updated_at`, kind, id, payload, now, now)
	if err != nil {
		return fmt.Errorf("save %s/%s: %w", kind, id, err)
	}
	return nil
}

func (s *Store) Load(ctx context.Context, kind, id string, target any) error {
	var payload []byte
	err := s.db.QueryRowContext(ctx, `SELECT payload FROM records WHERE kind=? AND id=?`, kind, id).Scan(&payload)
	if errors.Is(err, sql.ErrNoRows) { return nil }
	if err != nil {
		return fmt.Errorf("load %s/%s: %w", kind, id, err)
	}
	return decode(payload, target)
}

func (s *Store) Delete(ctx context.Context, kind, id string) error {
	result, err := s.db.ExecContext(ctx, `DELETE FROM records WHERE kind=? AND id=?`, kind, id)
	if err != nil {
		return fmt.Errorf("delete %s/%s: %w", kind, id, err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return model.ErrNotFound
	}
	return nil
}

func (s *Store) List(ctx context.Context, kind string, each func([]byte) error) error {
	rows, err := s.db.QueryContext(ctx, `SELECT payload FROM records WHERE kind=? ORDER BY updated_at, id`, kind)
	if err != nil {
		return fmt.Errorf("list %s: %w", kind, err)
	}
	defer rows.Close()
	for rows.Next() {
		var payload []byte
		if err := rows.Scan(&payload); err != nil {
			return fmt.Errorf("scan %s: %w", kind, err)
		}
		if err := each(payload); err != nil {
			return err
		}
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("iterate %s: %w", kind, err)
	}
	return nil
}

func (s *Store) Count(ctx context.Context, kind string) (int, error) {
	var count int
	if err := s.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM records WHERE kind=?`, kind).Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}
