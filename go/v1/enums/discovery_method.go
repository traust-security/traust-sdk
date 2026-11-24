// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

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
