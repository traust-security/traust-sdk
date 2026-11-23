package runs

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/traust-security/traust-sdk/go/v1/traust"
)

// RunOption tunes the order BeginSkill submits.
type RunOption func(*traust.OrderRequest)

// ForRepository scopes the run to one repository.
func ForRepository(id uuid.UUID) RunOption {
	return func(order *traust.OrderRequest) { order.RepositoryId = &id }
}

// BeginSkill submits a person's run of an asset by name and follows it to the end; the
// engine's claim worker resolves the asset in its own pack, so no config pin is sent.
func (c *Client) BeginSkill(ctx context.Context, asset string, opts ...RunOption) (*traust.Poller[traust.RunOrder], error) {
	if asset == "" {
		return nil, fmt.Errorf("runs: BeginSkill: asset is required")
	}
	// Core requires a priority and claims the lowest first; 1 is what a person's submit has used.
	order := traust.OrderRequest{Asset: asset, Priority: 1}
	for _, opt := range opts {
		opt(&order)
	}
	return c.BeginSubmit(ctx, traust.SubmitRequest{Order: order})
}
