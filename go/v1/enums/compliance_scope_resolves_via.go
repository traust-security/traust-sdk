// Code generated from traust-contracts c2f6e6fb12270f8b3c8703da7bf39bd01776e4b1. DO NOT EDIT.

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
