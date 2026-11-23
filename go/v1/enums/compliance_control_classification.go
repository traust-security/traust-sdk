// Code generated from traust-contracts 8f684d27d978e53c0ffc2b43af379fc1284c26eb. DO NOT EDIT.

package enums

// ComplianceControlClassification — Values of `classification` in compliance-assessment.schema.json.
type ComplianceControlClassification string

const (
	ComplianceControlClassificationDeterministic  ComplianceControlClassification = "deterministic"
	ComplianceControlClassificationEvidenceReview ComplianceControlClassification = "evidence_review"
	ComplianceControlClassificationOrganizational ComplianceControlClassification = "organizational"
)

// ComplianceControlClassificationValues returns all valid ComplianceControlClassification values.
func ComplianceControlClassificationValues() []ComplianceControlClassification {
	return []ComplianceControlClassification{
		ComplianceControlClassificationDeterministic,
		ComplianceControlClassificationEvidenceReview,
		ComplianceControlClassificationOrganizational,
	}
}
