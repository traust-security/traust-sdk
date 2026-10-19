// Code generated from traust-contracts 3ed68158121322ce06525de170aa383e057b6a7b. DO NOT EDIT.

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
