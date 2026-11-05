// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

package enums

// PrecisionGateAction — What a fired Precision Gate did with a critical/high candidate.
type PrecisionGateAction string

const (
	PrecisionGateActionDowngraded      PrecisionGateAction = "downgraded"
	PrecisionGateActionNegativeResults PrecisionGateAction = "negative_results"
	PrecisionGateActionDependencyAudit PrecisionGateAction = "dependency_audit"
)

// PrecisionGateActionValues returns all valid PrecisionGateAction values.
func PrecisionGateActionValues() []PrecisionGateAction {
	return []PrecisionGateAction{
		PrecisionGateActionDowngraded,
		PrecisionGateActionNegativeResults,
		PrecisionGateActionDependencyAudit,
	}
}
