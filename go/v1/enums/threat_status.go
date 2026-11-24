// Code generated from traust-contracts 58803ea798ad6431e10da337eb3a28349c1c9992. DO NOT EDIT.

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
