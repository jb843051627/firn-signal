package service

import (
	"context"
	"github.com/jb843051627/firn-signal/internal/store"
)

func (l *Lab) Health(ctx context.Context) (store.Health, error) { return l.store.Health(ctx) }
func (l *Lab) StorePath() string                                { return l.store.Path() }
