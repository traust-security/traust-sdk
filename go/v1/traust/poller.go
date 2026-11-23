package traust

import (
	"context"
	"fmt"
	"time"
)

// Poller follows work core finishes later. Keep its ResumeToken to pick the work up again
// after a restart, from any process.
type Poller[T any] struct {
	token string
	poll  func(context.Context) (T, bool, error)
}

// NewPoller builds a Poller for work a resume token names; poll asks core once for the
// latest state and whether the work has finished. The domain pollers are built on it.
func NewPoller[T any](token string, poll func(context.Context) (T, bool, error)) *Poller[T] {
	return &Poller[T]{token: token, poll: poll}
}

// ResumeToken names the work; for an order it is the order id.
func (p *Poller[T]) ResumeToken() string { return p.token }

// Poll asks core once for the latest state and whether the work has finished.
func (p *Poller[T]) Poll(ctx context.Context) (T, bool, error) { return p.poll(ctx) }

// Wait polls every interval until the work finishes or ctx ends.
func (p *Poller[T]) Wait(ctx context.Context, every time.Duration) (T, error) {
	if every <= 0 {
		var zero T
		return zero, fmt.Errorf("traust: Wait interval must be positive, got %s", every)
	}
	ticker := time.NewTicker(every)
	defer ticker.Stop()
	for {
		latest, done, err := p.poll(ctx)
		if err != nil || done {
			return latest, err
		}
		select {
		case <-ctx.Done():
			return latest, ctx.Err()
		case <-ticker.C:
		}
	}
}
