// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

package enums

// DiscoveryMethod — Values of `discovery_method` in validation.schema.json.
type DiscoveryMethod string

const (
	DiscoveryMethodReconDiff     DiscoveryMethod = "recon-diff"
	DiscoveryMethodPatternProbe  DiscoveryMethod = "pattern-probe"
	DiscoveryMethodChainEmergent DiscoveryMethod = "chain-emergent"
	DiscoveryMethodAiHypothesis  DiscoveryMethod = "ai-hypothesis"
)

// DiscoveryMethodValues returns all valid DiscoveryMethod values.
func DiscoveryMethodValues() []DiscoveryMethod {
	return []DiscoveryMethod{
		DiscoveryMethodReconDiff,
		DiscoveryMethodPatternProbe,
		DiscoveryMethodChainEmergent,
		DiscoveryMethodAiHypothesis,
	}
}
