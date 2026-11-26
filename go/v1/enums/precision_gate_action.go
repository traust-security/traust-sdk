// Code generated from traust-contracts 485578cbe584ad224d57b2c7c778174c0b66ba7c. DO NOT EDIT.

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
