package traust

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

// RunOption tunes the order BeginSkill submits.
type RunOption func(*OrderRequest)

// ForRepository scopes the run to one repository.
func ForRepository(id uuid.UUID) RunOption {
	return func(order *OrderRequest) { order.RepositoryId = &id }
}

// BeginSkill submits a person's run of an asset by name and follows it to the end; the
// engine's claim worker resolves the asset in its own pack, so no config pin is sent.
func (c *Client) BeginSkill(ctx context.Context, asset string, opts ...RunOption) (*Poller[RunOrder], error) {
	if asset == "" {
		return nil, fmt.Errorf("traust: BeginSkill: asset is required")
	}
	// Core requires a priority; 1 is a person's normal ask (0 is core's own ask-now lanes).
	order := OrderRequest{Asset: asset, Priority: 1}
	for _, opt := range opts {
		opt(&order)
	}
	return c.BeginSubmit(ctx, SubmitRequest{Order: order})
}
