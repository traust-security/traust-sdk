// Code generated from traust-contracts 485578cbe584ad224d57b2c7c778174c0b66ba7c. DO NOT EDIT.

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
