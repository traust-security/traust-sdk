// Code generated from traust-contracts c2f6e6fb12270f8b3c8703da7bf39bd01776e4b1. DO NOT EDIT.

package enums

// ThreatLikelihood — How often this threat is expected to be driven, on an ordinal scale.
type ThreatLikelihood string

const (
	ThreatLikelihoodVeryRare      ThreatLikelihood = "very_rare"
	ThreatLikelihoodRare          ThreatLikelihood = "rare"
	ThreatLikelihoodPossible      ThreatLikelihood = "possible"
	ThreatLikelihoodLikely        ThreatLikelihood = "likely"
	ThreatLikelihoodAlmostCertain ThreatLikelihood = "almost_certain"
)

// ThreatLikelihoodValues returns all valid ThreatLikelihood values.
func ThreatLikelihoodValues() []ThreatLikelihood {
	return []ThreatLikelihood{
		ThreatLikelihoodVeryRare,
		ThreatLikelihoodRare,
		ThreatLikelihoodPossible,
		ThreatLikelihoodLikely,
		ThreatLikelihoodAlmostCertain,
	}
}
