// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

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
