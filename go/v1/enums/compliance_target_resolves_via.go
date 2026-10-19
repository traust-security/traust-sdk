// Code generated from traust-contracts 3ed68158121322ce06525de170aa383e057b6a7b. DO NOT EDIT.

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
