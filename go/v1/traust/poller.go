package traust

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Poller follows work core finishes later. Keep its ResumeToken to pick the work up again
// after a restart, from any process.
type Poller[T any] struct {
	token string
	poll  func(context.Context) (T, bool, error)
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

// BeginSubmit submits an order and follows it until core finishes it.
func (c *Client) BeginSubmit(ctx context.Context, body SubmitRequest) (*Poller[RunOrder], error) {
	order, err := c.Submit(ctx, body)
	if err != nil {
		return nil, err
	}
	if order.Id == nil {
		return nil, errors.New("traust: Submit answered an order without an id")
	}
	return c.ResumeOrder(order.Id.String())
}

// ResumeOrder follows the order a resume token names. Core's finished_at, not a local list of
// states, says when it has finished.
func (c *Client) ResumeOrder(token string) (*Poller[RunOrder], error) {
	id, err := uuid.Parse(token)
	if err != nil {
		return nil, fmt.Errorf("traust: resume token: %w", err)
	}
	return &Poller[RunOrder]{token: token, poll: func(ctx context.Context) (RunOrder, bool, error) {
		order, err := c.GetOrder(ctx, id)
		if err != nil {
			return RunOrder{}, false, err
		}
		return *order, order.FinishedAt != nil, nil
	}}, nil
}
