// Code generated from traust-contracts 31d0d884672149e606934c2844ffcbf3bdd4b925. DO NOT EDIT.

package enums

// Severity — Finding severity classification. Used across finding schemas (report, vuln-findings, cloud-config-audit, remediation, validation, etc.)
type Severity string

const (
	SeverityCritical      Severity = "critical"
	SeverityHigh          Severity = "high"
	SeverityMedium        Severity = "medium"
	SeverityLow           Severity = "low"
	SeverityInformational Severity = "informational"
)

// SeverityValues returns all valid Severity values.
func SeverityValues() []Severity {
	return []Severity{
		SeverityCritical,
		SeverityHigh,
		SeverityMedium,
		SeverityLow,
		SeverityInformational,
	}
}
