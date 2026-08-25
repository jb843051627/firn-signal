package clock

import (
	"sync"
	"time"
)

type Fixed struct {
	mu    sync.RWMutex
	value time.Time
}

func NewFixed(value time.Time) *Fixed { return &Fixed{value: value.UTC()} }
func (f *Fixed) Now() time.Time       { f.mu.RLock(); defer f.mu.RUnlock(); return f.value }
func (f *Fixed) Set(value time.Time)  { f.mu.Lock(); f.value = value.UTC(); f.mu.Unlock() }
func (f *Fixed) Advance(delta time.Duration) {
	f.mu.Lock()
	f.value = f.value.Add(delta)
	f.mu.Unlock()
}
