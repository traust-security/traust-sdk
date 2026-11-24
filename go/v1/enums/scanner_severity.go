// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

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
