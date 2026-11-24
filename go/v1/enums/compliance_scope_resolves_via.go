// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

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
