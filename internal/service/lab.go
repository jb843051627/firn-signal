package service

import (
	"context"
	"sync"

	"github.com/jb843051627/firn-signal/internal/clock"
	"github.com/jb843051627/firn-signal/internal/ingest"
	"github.com/jb843051627/firn-signal/internal/metrics"
	"github.com/jb843051627/firn-signal/internal/policy"
	"github.com/jb843051627/firn-signal/internal/store"
)

type Lab struct {
	store     *store.Store
	clock     clock.Clock
	policy    policy.Engine
	queue     *ingest.Queue
	metrics   *metrics.Registry
	lockMu    sync.Mutex
	scanLocks map[string]*sync.Mutex
}

func NewLab(repository *store.Store) *Lab {
	return &Lab{store: repository, clock: clock.System{}, policy: policy.New(), queue: ingest.New(128, 2), metrics: metrics.New(), scanLocks: make(map[string]*sync.Mutex)}
}

func NewLabWithClock(repository *store.Store, source clock.Clock) *Lab {
	lab := NewLab(repository)
	if source != nil {
		lab.clock = source
	}
	return lab
}

func (l *Lab) Close() error                                             { l.queue.Close(); return nil }
func (l *Lab) Metrics() map[string]int64                                { return l.metrics.Snapshot() }
func (l *Lab) DatabaseHealth(ctx context.Context) (store.Health, error) { return l.store.Health(ctx) }
