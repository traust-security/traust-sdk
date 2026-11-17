// Code generated from traust-contracts 31d0d884672149e606934c2844ffcbf3bdd4b925. DO NOT EDIT.

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
