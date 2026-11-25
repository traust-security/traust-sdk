// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

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
