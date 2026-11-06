// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

package enums

// ScannerSeverity — from the facts artifact; OSS Checkov emits none, so this is normally 'unrated' and the gate then requires the rationale to name the SKILL.md rubric row applied
type ScannerSeverity string

const (
	ScannerSeverityCritical      ScannerSeverity = "critical"
	ScannerSeverityHigh          ScannerSeverity = "high"
	ScannerSeverityMedium        ScannerSeverity = "medium"
	ScannerSeverityLow           ScannerSeverity = "low"
	ScannerSeverityInformational ScannerSeverity = "informational"
	ScannerSeverityUnrated       ScannerSeverity = "unrated"
)

// ScannerSeverityValues returns all valid ScannerSeverity values.
func ScannerSeverityValues() []ScannerSeverity {
	return []ScannerSeverity{
		ScannerSeverityCritical,
		ScannerSeverityHigh,
		ScannerSeverityMedium,
		ScannerSeverityLow,
		ScannerSeverityInformational,
		ScannerSeverityUnrated,
	}
}
