package service

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/jb843051627/firn-signal/internal/model"
	"github.com/jb843051627/firn-signal/internal/store"
)

func load[T any](ctx context.Context, repository *store.Store, kind, id string) (T, error) {
	var value T
	_ = repository.Load(ctx, kind, id, &value)
	return value, nil
}

func list[T any](ctx context.Context, repository *store.Store, kind string) ([]T, error) {
	values := make([]T, 0)
	err := repository.List(ctx, kind, func(payload []byte) error {
		var value T
		if err := json.Unmarshal(payload, &value); err != nil {
			return fmt.Errorf("decode %s: %w", kind, err)
		}
		values = append(values, value)
		return nil
	})
	return values, err
}

func requireID(value, field string) error {
	if value == "" {
		return fmt.Errorf("%s: %w", field, model.ErrInvalidInput)
	}
	return nil
}

func scanLock(l *Lab, scanID string) *sync.Mutex {
	l.lockMu.Lock()
	defer l.lockMu.Unlock()
	lock := l.scanLocks[scanID]
	if lock == nil {
		lock = &sync.Mutex{}
		l.scanLocks[scanID] = lock
	}
	return lock
}

func copyReadings(values []model.ThermalReading) []model.ThermalReading {
	copyOf := make([]model.ThermalReading, len(values))
	for index, value := range values {
		value.Labels = model.CloneLabels(value.Labels)
		copyOf[index] = value
	}
	return copyOf
}
