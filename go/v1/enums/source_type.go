// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

package enums

// SourceType — Historical combined source vocabulary. Human channels now belong to Decision.source_channel, producing machine kinds to Assessment.kind. Rebaseline and restatement remain administrative operations, not evidence kinds. Source type never supplies a missing evidence grade.
type SourceType string

const (
	SourceTypeMrComment          SourceType = "mr_comment"
	SourceTypeCommit             SourceType = "commit"
	SourceTypeJira               SourceType = "jira"
	SourceTypeInteractive        SourceType = "interactive"
	SourceTypeValidationReport   SourceType = "validation_report"
	SourceTypeVerificationReport SourceType = "verification_report"
	SourceTypeTriageReport       SourceType = "triage_report"
	SourceTypeImpactReport       SourceType = "impact_report"
	SourceTypeVulnScanReport     SourceType = "vuln_scan_report"
	SourceTypeFuzzReport         SourceType = "fuzz_report"
	SourceTypeRebaseline         SourceType = "rebaseline"
	SourceTypeRestatement        SourceType = "restatement"
)

// SourceTypeValues returns all valid SourceType values.
func SourceTypeValues() []SourceType {
	return []SourceType{
		SourceTypeMrComment,
		SourceTypeCommit,
		SourceTypeJira,
		SourceTypeInteractive,
		SourceTypeValidationReport,
		SourceTypeVerificationReport,
		SourceTypeTriageReport,
		SourceTypeImpactReport,
		SourceTypeVulnScanReport,
		SourceTypeFuzzReport,
		SourceTypeRebaseline,
		SourceTypeRestatement,
	}
}
