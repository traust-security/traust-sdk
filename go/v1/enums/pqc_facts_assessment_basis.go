// Code generated from traust-contracts 58803ea798ad6431e10da337eb3a28349c1c9992. DO NOT EDIT.

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
