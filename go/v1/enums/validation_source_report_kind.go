// Code generated from traust-contracts c2f6e6fb12270f8b3c8703da7bf39bd01776e4b1. DO NOT EDIT.

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
