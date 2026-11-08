// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

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
