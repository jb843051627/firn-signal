package store

import (
	"context"
	"fmt"
)

func (s *Store) schema(ctx context.Context) error {
	statements := []string{
		`CREATE TABLE IF NOT EXISTS records (kind TEXT NOT NULL, id TEXT NOT NULL, payload BLOB NOT NULL, created_at TEXT NOT NULL, updated_at TEXT NOT NULL, PRIMARY KEY(kind, id))`,
		`CREATE INDEX IF NOT EXISTS records_kind_index ON records(kind, updated_at)`,
		`CREATE TABLE IF NOT EXISTS events (id INTEGER PRIMARY KEY AUTOINCREMENT, subject TEXT NOT NULL, action TEXT NOT NULL, details BLOB NOT NULL, created_at TEXT NOT NULL)`,
		`CREATE INDEX IF NOT EXISTS events_subject_index ON events(subject, created_at)`,
	}
	for _, statement := range statements {
		if _, err := s.db.ExecContext(ctx, statement); err != nil {
			return fmt.Errorf("create schema: %w", err)
		}
	}
	return nil
}
