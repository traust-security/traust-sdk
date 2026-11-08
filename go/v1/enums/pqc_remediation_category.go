// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

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
