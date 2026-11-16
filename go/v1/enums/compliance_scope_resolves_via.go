// Code generated from traust-contracts 31d0d884672149e606934c2844ffcbf3bdd4b925. DO NOT EDIT.

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
