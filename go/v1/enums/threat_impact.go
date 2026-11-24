// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

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
