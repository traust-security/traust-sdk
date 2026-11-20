// Code generated from traust-contracts 15bded2e3e4b6c70bd1e20ef8df50c7b66b5ebc1. DO NOT EDIT.

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
