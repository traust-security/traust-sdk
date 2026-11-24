// Code generated from traust-contracts 58803ea798ad6431e10da337eb3a28349c1c9992. DO NOT EDIT.

package enums

// ExternalRefConfidence — 'confirmed' = title match at or above the auto-stamp threshold, or corroborated by an overlapping CWE.
type ExternalRefConfidence string

const (
	ExternalRefConfidenceConfirmed ExternalRefConfidence = "confirmed"
	ExternalRefConfidenceProbable  ExternalRefConfidence = "probable"
)

// ExternalRefConfidenceValues returns all valid ExternalRefConfidence values.
func ExternalRefConfidenceValues() []ExternalRefConfidence {
	return []ExternalRefConfidence{
		ExternalRefConfidenceConfirmed,
		ExternalRefConfidenceProbable,
	}
}
