// Code generated from traust-contracts 369a2ab2c9b65225c6878eca75bada6a65df548f. DO NOT EDIT.

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
