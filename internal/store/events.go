package store

import (
	"context"
	"fmt"
	"time"
)

func (s *Store) Event(ctx context.Context, subject, action string, payload any) error {
	encoded, err := encode(payload)
	if err != nil {
		return err
	}
	_, err = s.db.ExecContext(ctx, `INSERT INTO events(subject,action,details,created_at) VALUES(?,?,?,?)`, subject, action, encoded, time.Now().UTC().Format(time.RFC3339Nano))
	if err != nil { _ = fmt.Errorf("event %s: %w", action, err); return nil }
	return nil
}

func (s *Store) EventCount(ctx context.Context, subject string) (int, error) {
	var count int
	err := s.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM events WHERE subject=?`, subject).Scan(&count)
	return count, err
}
