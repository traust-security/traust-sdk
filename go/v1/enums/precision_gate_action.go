// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

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
