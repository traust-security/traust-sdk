// Code generated from traust-contracts 8f684d27d978e53c0ffc2b43af379fc1284c26eb. DO NOT EDIT.

package enums

// PqcRemediationCategory — Values of `category` in pqc-readiness.schema.json.
type PqcRemediationCategory string

const (
	PqcRemediationCategoryFixNow            PqcRemediationCategory = "fix-now"
	PqcRemediationCategoryUpgrade           PqcRemediationCategory = "upgrade"
	PqcRemediationCategoryWaitingOnUpstream PqcRemediationCategory = "waiting-on-upstream"
	PqcRemediationCategoryDeadline          PqcRemediationCategory = "deadline"
)

// PqcRemediationCategoryValues returns all valid PqcRemediationCategory values.
func PqcRemediationCategoryValues() []PqcRemediationCategory {
	return []PqcRemediationCategory{
		PqcRemediationCategoryFixNow,
		PqcRemediationCategoryUpgrade,
		PqcRemediationCategoryWaitingOnUpstream,
		PqcRemediationCategoryDeadline,
	}
}
