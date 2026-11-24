// Code generated from traust-contracts c2f6e6fb12270f8b3c8703da7bf39bd01776e4b1. DO NOT EDIT.

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
