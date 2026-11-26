// Code generated from traust-contracts 485578cbe584ad224d57b2c7c778174c0b66ba7c. DO NOT EDIT.

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
