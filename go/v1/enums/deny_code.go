// Code generated from traust-contracts 8f684d27d978e53c0ffc2b43af379fc1284c26eb. DO NOT EDIT.

package enums

// DenyCode — Machine-readable denial vocabulary for the platform API: why a policy gate (preflight, submit) refused an action, so consumers fail fast and route on the reason without parsing prose. Carried by the typed errors in-process and by the problem+json envelope's code and issues[].code over HTTP.
type DenyCode string

const (
	DenyCodeSubjectUnknown      DenyCode = "subject_unknown"
	DenyCodeSubjectUnreachable  DenyCode = "subject_unreachable"
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
		DenyCodeSubjectUnreachable,
		DenyCodeRoleDenied,
		DenyCodeOrderingViolated,
		DenyCodeBudgetExhausted,
		DenyCodeConfigStale,
		DenyCodeCountersignRequired,
		DenyCodeReviewRequired,
	}
}
