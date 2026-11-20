// Code generated from traust-contracts 15bded2e3e4b6c70bd1e20ef8df50c7b66b5ebc1. DO NOT EDIT.

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
