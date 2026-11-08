// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

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
