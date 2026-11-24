// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

package enums

// RunOrderState — Operational vocabulary for traust_ops.run_order.state; Core enforces new writes, while stored historical text remains readable.
type RunOrderState string

const (
	RunOrderStatePlanned          RunOrderState = "planned"
	RunOrderStateAwaitingApproval RunOrderState = "awaiting_approval"
	RunOrderStateQueued           RunOrderState = "queued"
	RunOrderStateLeased           RunOrderState = "leased"
	RunOrderStateSucceeded        RunOrderState = "succeeded"
	RunOrderStateDegraded         RunOrderState = "degraded"
	RunOrderStateRefused          RunOrderState = "refused"
	RunOrderStateWithheld         RunOrderState = "withheld"
	RunOrderStateFailed           RunOrderState = "failed"
)

// RunOrderStateValues returns all valid RunOrderState values.
func RunOrderStateValues() []RunOrderState {
	return []RunOrderState{
		RunOrderStatePlanned,
		RunOrderStateAwaitingApproval,
		RunOrderStateQueued,
		RunOrderStateLeased,
		RunOrderStateSucceeded,
		RunOrderStateDegraded,
		RunOrderStateRefused,
		RunOrderStateWithheld,
		RunOrderStateFailed,
	}
}
