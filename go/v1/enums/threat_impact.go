// Code generated from traust-contracts c2f6e6fb12270f8b3c8703da7bf39bd01776e4b1. DO NOT EDIT.

package enums

// ThreatImpact — How bad this threat is if realised, modelled BEFORE reachability is known. Deliberately not the finding severity enum.
type ThreatImpact string

const (
	ThreatImpactLow         ThreatImpact = "low"
	ThreatImpactMedium      ThreatImpact = "medium"
	ThreatImpactHigh        ThreatImpact = "high"
	ThreatImpactCritical    ThreatImpact = "critical"
	ThreatImpactExistential ThreatImpact = "existential"
)

// ThreatImpactValues returns all valid ThreatImpact values.
func ThreatImpactValues() []ThreatImpact {
	return []ThreatImpact{
		ThreatImpactLow,
		ThreatImpactMedium,
		ThreatImpactHigh,
		ThreatImpactCritical,
		ThreatImpactExistential,
	}
}
