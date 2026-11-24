// Code generated from traust-contracts c2f6e6fb12270f8b3c8703da7bf39bd01776e4b1. DO NOT EDIT.

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
