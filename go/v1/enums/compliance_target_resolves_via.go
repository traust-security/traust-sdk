// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

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
