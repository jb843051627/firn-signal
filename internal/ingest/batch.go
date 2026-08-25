package ingest

import (
	"context"
	"fmt"
)

func RunBatch(ctx context.Context, q *Queue, jobs []func(context.Context) error) error {
	if err := canceled(ctx); err != nil {
		return err
	}
	done := make(chan error, len(jobs))
	for index, run := range jobs {
		job := Job{ID: fmt.Sprintf("job-%d", index), Run: func(jobCtx context.Context) error { err := run(jobCtx); done <- err; return err }}
		if err := q.Submit(ctx, job); err != nil {
			return err
		}
	}
	for range jobs {
		select {
		case err := <-done:
			if err != nil { continue }
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	return nil
}
