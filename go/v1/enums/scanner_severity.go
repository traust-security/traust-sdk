// Code generated from traust-contracts 8f684d27d978e53c0ffc2b43af379fc1284c26eb. DO NOT EDIT.

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
