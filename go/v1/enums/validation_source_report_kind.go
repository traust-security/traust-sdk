// Code generated from traust-contracts 369a2ab2c9b65225c6878eca75bada6a65df548f. DO NOT EDIT.

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
