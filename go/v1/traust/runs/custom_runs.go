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

// BeginSkill submits a person's run of an asset by name at the priority the caller chooses (core
// claims the lowest first) and follows it to the end; the engine's claim worker resolves the
// asset in its own pack, so no config pin is sent.
func (c *Client) BeginSkill(ctx context.Context, asset string, priority int, opts ...RunOption) (*traust.Poller[traust.RunOrder], error) {
	if asset == "" {
		return nil, fmt.Errorf("runs: BeginSkill: asset is required")
	}
	order := traust.OrderRequest{Asset: asset, Priority: priority}
	for _, opt := range opts {
		opt(&order)
	}
	return c.BeginSubmit(ctx, traust.SubmitRequest{Order: order})
}
