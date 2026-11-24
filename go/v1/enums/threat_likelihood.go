// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

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
