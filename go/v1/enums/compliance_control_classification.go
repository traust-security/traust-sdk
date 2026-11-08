// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

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
