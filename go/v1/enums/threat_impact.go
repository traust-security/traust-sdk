// Code generated from traust-contracts ac517047c6566b019e63b911f6b7522000826b48. DO NOT EDIT.

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
