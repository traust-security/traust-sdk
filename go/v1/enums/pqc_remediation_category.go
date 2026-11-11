// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

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
