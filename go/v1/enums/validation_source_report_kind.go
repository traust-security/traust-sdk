// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

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
