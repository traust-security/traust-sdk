// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

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
