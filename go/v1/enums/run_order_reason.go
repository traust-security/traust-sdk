// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

package enums

// RunOrderReason — Operational vocabulary for traust_ops.run_order.reason; Core enforces new writes, while stored historical text remains readable.
type RunOrderReason string

const (
	RunOrderReasonBudgetOver           RunOrderReason = "budget-over"
	RunOrderReasonAssetNotInPack       RunOrderReason = "asset-not-in-pack"
	RunOrderReasonBudgetExhausted      RunOrderReason = "budget-exhausted"
	RunOrderReasonDiffTooLarge         RunOrderReason = "diff-too-large"
	RunOrderReasonRunnerCrashed        RunOrderReason = "runner-crashed"
	RunOrderReasonSessionOutputInvalid RunOrderReason = "session-output-invalid"
	RunOrderReasonBadRequest           RunOrderReason = "bad_request"
	RunOrderReasonNotFound             RunOrderReason = "not_found"
	RunOrderReasonSnapshotUnavailable  RunOrderReason = "snapshot_unavailable"
	RunOrderReasonSourceUnavailable    RunOrderReason = "source_unavailable"
	RunOrderReasonLedgerUnconfigured   RunOrderReason = "ledger_unconfigured"
	RunOrderReasonUpstreamStale        RunOrderReason = "upstream-stale"
	RunOrderReasonWriteRejected        RunOrderReason = "write-rejected"
)

// RunOrderReasonValues returns all valid RunOrderReason values.
func RunOrderReasonValues() []RunOrderReason {
	return []RunOrderReason{
		RunOrderReasonBudgetOver,
		RunOrderReasonAssetNotInPack,
		RunOrderReasonBudgetExhausted,
		RunOrderReasonDiffTooLarge,
		RunOrderReasonRunnerCrashed,
		RunOrderReasonSessionOutputInvalid,
		RunOrderReasonBadRequest,
		RunOrderReasonNotFound,
		RunOrderReasonSnapshotUnavailable,
		RunOrderReasonSourceUnavailable,
		RunOrderReasonLedgerUnconfigured,
		RunOrderReasonUpstreamStale,
		RunOrderReasonWriteRejected,
	}
}
