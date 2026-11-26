// Code generated from traust-contracts 485578cbe584ad224d57b2c7c778174c0b66ba7c. DO NOT EDIT.

package enums

// RemediationCommitRelevance — Values of `relevance` in verification.schema.json.
type RemediationCommitRelevance string

const (
	RemediationCommitRelevanceDirect     RemediationCommitRelevance = "direct"
	RemediationCommitRelevanceSupporting RemediationCommitRelevance = "supporting"
	RemediationCommitRelevancePartial    RemediationCommitRelevance = "partial"
)

// RemediationCommitRelevanceValues returns all valid RemediationCommitRelevance values.
func RemediationCommitRelevanceValues() []RemediationCommitRelevance {
	return []RemediationCommitRelevance{
		RemediationCommitRelevanceDirect,
		RemediationCommitRelevanceSupporting,
		RemediationCommitRelevancePartial,
	}
}
