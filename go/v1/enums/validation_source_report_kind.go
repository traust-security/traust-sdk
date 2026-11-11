// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

package enums

// ValidationSourceReportKind — Values of `kind` in validation.schema.json.
type ValidationSourceReportKind string

const (
	ValidationSourceReportKindSecurityAudit ValidationSourceReportKind = "security-audit"
	ValidationSourceReportKindThreatModel   ValidationSourceReportKind = "threat-model"
	ValidationSourceReportKindTriage        ValidationSourceReportKind = "triage"
)

// ValidationSourceReportKindValues returns all valid ValidationSourceReportKind values.
func ValidationSourceReportKindValues() []ValidationSourceReportKind {
	return []ValidationSourceReportKind{
		ValidationSourceReportKindSecurityAudit,
		ValidationSourceReportKindThreatModel,
		ValidationSourceReportKindTriage,
	}
}
