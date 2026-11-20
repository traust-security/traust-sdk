// Code generated from traust-contracts 15bded2e3e4b6c70bd1e20ef8df50c7b66b5ebc1. DO NOT EDIT.

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
