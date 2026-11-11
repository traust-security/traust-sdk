// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

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
