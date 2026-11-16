// Code generated from traust-contracts 31d0d884672149e606934c2844ffcbf3bdd4b925. DO NOT EDIT.

package enums

// ComplianceTargetKind — Values of `kind` in compliance-assessment.schema.json.
type ComplianceTargetKind string

const (
	ComplianceTargetKindProduct     ComplianceTargetKind = "product"
	ComplianceTargetKindEnvironment ComplianceTargetKind = "environment"
	ComplianceTargetKindBoth        ComplianceTargetKind = "both"
)

// ComplianceTargetKindValues returns all valid ComplianceTargetKind values.
func ComplianceTargetKindValues() []ComplianceTargetKind {
	return []ComplianceTargetKind{
		ComplianceTargetKindProduct,
		ComplianceTargetKindEnvironment,
		ComplianceTargetKindBoth,
	}
}
