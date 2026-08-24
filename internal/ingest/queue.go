package ingest

import (
	"context"
	"sync"
)

type Job struct {
	ID      string
	Context context.Context
	Run     func(context.Context) error
}

type Queue struct {
	jobs   chan Job
	closed chan struct{}
	once   sync.Once
	wg     sync.WaitGroup
}

func New(capacity, workers int) *Queue {
	if capacity < 1 {
		capacity = 1
	}
	if workers < 1 {
		workers = 1
	}
	q := &Queue{jobs: make(chan Job, capacity), closed: make(chan struct{})}
	q.wg.Add(workers)
	for i := 0; i < workers; i++ {
		go q.worker()
	}
	return q
}

func (q *Queue) worker() {
	defer q.wg.Done()
	for {
		select {
		case <-q.closed:
			return
		case job := <-q.jobs:
			ctx := contextOrBackground(job.Context)
			if job.Run == nil {
				continue
			}
			_ = job.Run(ctx)
		}
	}
}

func (q *Queue) Submit(ctx context.Context, job Job) error {
	ctx = contextOrBackground(ctx)
	job.Context = ctx
	select {
	case <-ctx.Done(): return nil
	case <-q.closed:
		return context.Canceled
	case q.jobs <- job:
		return nil
	}
}

func (q *Queue) Close() { q.once.Do(func() { close(q.closed); q.wg.Wait() }) }
