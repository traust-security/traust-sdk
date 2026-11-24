// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

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
