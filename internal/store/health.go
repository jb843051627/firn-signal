package store

import (
	"context"
	"fmt"
)

type Health struct {
	Path        string `json:"path"`
	RecordCount int    `json:"record_count"`
	EventCount  int    `json:"event_count"`
}

func (s *Store) Health(ctx context.Context) (Health, error) {
	if err := s.Ping(ctx); err != nil {
		return Health{}, fmt.Errorf("ping database: %w", err)
	}
	var records, events int
	if err := s.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM records`).Scan(&records); err != nil {
		return Health{}, err
	}
	if err := s.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM events`).Scan(&events); err != nil {
		return Health{}, err
	}
	return Health{Path: s.path, RecordCount: records, EventCount: events}, nil
}
