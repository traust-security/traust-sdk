// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

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
