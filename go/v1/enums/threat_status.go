// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

package enums

// ThreatStatus — Mitigation state of a modelled threat.
type ThreatStatus string

const (
	ThreatStatusUnmitigated        ThreatStatus = "unmitigated"
	ThreatStatusPartiallyMitigated ThreatStatus = "partially_mitigated"
	ThreatStatusMitigated          ThreatStatus = "mitigated"
	ThreatStatusRiskAccepted       ThreatStatus = "risk_accepted"
)

// ThreatStatusValues returns all valid ThreatStatus values.
func ThreatStatusValues() []ThreatStatus {
	return []ThreatStatus{
		ThreatStatusUnmitigated,
		ThreatStatusPartiallyMitigated,
		ThreatStatusMitigated,
		ThreatStatusRiskAccepted,
	}
}
