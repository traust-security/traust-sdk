// Code generated from traust-contracts 15bded2e3e4b6c70bd1e20ef8df50c7b66b5ebc1. DO NOT EDIT.

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
