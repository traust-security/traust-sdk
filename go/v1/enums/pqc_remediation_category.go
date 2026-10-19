// Code generated from traust-contracts 3ed68158121322ce06525de170aa383e057b6a7b. DO NOT EDIT.

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
