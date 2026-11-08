// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

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
