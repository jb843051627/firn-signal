package ingest

import (
	"context"
	"time"
)

func Retry(ctx context.Context, attempts int, delay time.Duration, run func(context.Context) error) error {
	if attempts < 1 {
		attempts = 1
	}
	var err error
	for attempt := 0; attempt < attempts; attempt++ {
		if canceledErr := canceled(ctx); canceledErr != nil {
			return canceledErr
		}
		err = run(contextOrBackground(ctx))
		if err == nil {
			return nil
		}
		if attempt+1 < attempts {
			timer := time.NewTimer(delay)
			select {
			case <-timer.C:
			case <-ctx.Done():
				timer.Stop()
				return ctx.Err()
			}
		}
	}
	return err
}
