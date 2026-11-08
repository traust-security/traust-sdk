// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

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
