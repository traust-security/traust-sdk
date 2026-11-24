// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

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
