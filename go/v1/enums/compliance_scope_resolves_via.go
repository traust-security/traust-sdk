// Code generated from traust-contracts 3ed68158121322ce06525de170aa383e057b6a7b. DO NOT EDIT.

package enums

// ComplianceScopeResolvesVia — repo-graph: membership = the product node's ships edges plus include minus exclude (organizational assertion).
type ComplianceScopeResolvesVia string

const (
	ComplianceScopeResolvesViaRepoGraph     ComplianceScopeResolvesVia = "repo-graph"
	ComplianceScopeResolvesViaDeploymentIac ComplianceScopeResolvesVia = "deployment-iac"
	ComplianceScopeResolvesViaExplicit      ComplianceScopeResolvesVia = "explicit"
)

// ComplianceScopeResolvesViaValues returns all valid ComplianceScopeResolvesVia values.
func ComplianceScopeResolvesViaValues() []ComplianceScopeResolvesVia {
	return []ComplianceScopeResolvesVia{
		ComplianceScopeResolvesViaRepoGraph,
		ComplianceScopeResolvesViaDeploymentIac,
		ComplianceScopeResolvesViaExplicit,
	}
}
