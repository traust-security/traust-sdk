package runs

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"

	"github.com/traust-security/traust-sdk/go/v1/traust"
)

// BeginSubmit submits an order and follows it until core finishes it.
func (c *Client) BeginSubmit(ctx context.Context, body traust.SubmitRequest) (*traust.Poller[traust.RunOrder], error) {
	order, err := c.Submit(ctx, body)
	if err != nil {
		return nil, err
	}
	if order.Id == nil {
		return nil, errors.New("runs: Submit answered an order without an id")
	}
	return c.ResumeOrder(order.Id.String())
}

// ResumeOrder follows the order a resume token names. Core's finished_at, not a local list of
// states, says when it has finished.
func (c *Client) ResumeOrder(token string) (*traust.Poller[traust.RunOrder], error) {
	id, err := uuid.Parse(token)
	if err != nil {
		return nil, fmt.Errorf("runs: resume token: %w", err)
	}
	return traust.NewPoller(token, func(ctx context.Context) (traust.RunOrder, bool, error) {
		order, err := c.GetOrder(ctx, id)
		if err != nil {
			return traust.RunOrder{}, false, err
		}
		return *order, order.FinishedAt != nil, nil
	}), nil
}
