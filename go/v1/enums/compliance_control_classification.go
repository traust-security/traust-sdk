// Code generated from traust-contracts 485578cbe584ad224d57b2c7c778174c0b66ba7c. DO NOT EDIT.

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
