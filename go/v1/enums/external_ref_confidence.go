// Code generated from traust-contracts 2e521561f022fe839dd828ddf2185ba3a0a1045f. DO NOT EDIT.

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
