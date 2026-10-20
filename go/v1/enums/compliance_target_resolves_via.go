// Code generated from traust-contracts ac517047c6566b019e63b911f6b7522000826b48. DO NOT EDIT.

package enums

// ComplianceTargetResolvesVia — Values of `resolves_via` in compliance-assessment.schema.json.
type ComplianceTargetResolvesVia string

const (
	ComplianceTargetResolvesViaRepoGraph ComplianceTargetResolvesVia = "repo-graph"
	ComplianceTargetResolvesViaExplicit  ComplianceTargetResolvesVia = "explicit"
)

// ComplianceTargetResolvesViaValues returns all valid ComplianceTargetResolvesVia values.
func ComplianceTargetResolvesViaValues() []ComplianceTargetResolvesVia {
	return []ComplianceTargetResolvesVia{
		ComplianceTargetResolvesViaRepoGraph,
		ComplianceTargetResolvesViaExplicit,
	}
}
