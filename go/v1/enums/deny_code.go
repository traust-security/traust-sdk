// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

package enums

// DenyCode — Machine-readable denial vocabulary for the platform API: why a policy gate (preflight, submit) refused an action, so consumers fail fast and route on the reason without parsing prose. Carried by the typed errors in-process and by the problem+json envelope's code and issues[].code over HTTP.
type DenyCode string

const (
	DenyCodeSubjectUnknown      DenyCode = "subject_unknown"
	DenyCodeRoleDenied          DenyCode = "role_denied"
	DenyCodeOrderingViolated    DenyCode = "ordering_violated"
	DenyCodeBudgetExhausted     DenyCode = "budget_exhausted"
	DenyCodeConfigStale         DenyCode = "config_stale"
	DenyCodeCountersignRequired DenyCode = "countersign_required"
	DenyCodeReviewRequired      DenyCode = "review_required"
)

// DenyCodeValues returns all valid DenyCode values.
func DenyCodeValues() []DenyCode {
	return []DenyCode{
		DenyCodeSubjectUnknown,
		DenyCodeRoleDenied,
		DenyCodeOrderingViolated,
		DenyCodeBudgetExhausted,
		DenyCodeConfigStale,
		DenyCodeCountersignRequired,
		DenyCodeReviewRequired,
	}
}
