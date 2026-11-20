// Code generated from traust-contracts 15bded2e3e4b6c70bd1e20ef8df50c7b66b5ebc1. DO NOT EDIT.

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
