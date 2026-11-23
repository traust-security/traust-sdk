// Code generated from traust-contracts 8f684d27d978e53c0ffc2b43af379fc1284c26eb. DO NOT EDIT.

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
