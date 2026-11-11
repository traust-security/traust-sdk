// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

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
