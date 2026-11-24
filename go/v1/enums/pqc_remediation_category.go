// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

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
