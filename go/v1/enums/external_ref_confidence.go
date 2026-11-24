// Code generated from traust-contracts c2f6e6fb12270f8b3c8703da7bf39bd01776e4b1. DO NOT EDIT.

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
