// Code generated from traust-contracts 485578cbe584ad224d57b2c7c778174c0b66ba7c. DO NOT EDIT.

package enums

// PqcFactsAssessmentBasis — Values of `assessment_basis` in pqc-facts.schema.json.
type PqcFactsAssessmentBasis string

const (
	PqcFactsAssessmentBasisSource   PqcFactsAssessmentBasis = "source"
	PqcFactsAssessmentBasisSbomOnly PqcFactsAssessmentBasis = "sbom-only"
)

// PqcFactsAssessmentBasisValues returns all valid PqcFactsAssessmentBasis values.
func PqcFactsAssessmentBasisValues() []PqcFactsAssessmentBasis {
	return []PqcFactsAssessmentBasis{
		PqcFactsAssessmentBasisSource,
		PqcFactsAssessmentBasisSbomOnly,
	}
}
