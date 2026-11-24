// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

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
