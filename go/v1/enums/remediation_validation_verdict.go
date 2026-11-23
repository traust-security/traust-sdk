// Code generated from traust-contracts 8f684d27d978e53c0ffc2b43af379fc1284c26eb. DO NOT EDIT.

package enums

// RemediationValidationVerdict — Values of `before_verdict` in remediation.schema.json.
type RemediationValidationVerdict string

const (
	RemediationValidationVerdictConfirmed      RemediationValidationVerdict = "confirmed"
	RemediationValidationVerdictRefuted        RemediationValidationVerdict = "refuted"
	RemediationValidationVerdictInconclusive   RemediationValidationVerdict = "inconclusive"
	RemediationValidationVerdictBlockedByScope RemediationValidationVerdict = "blocked_by_scope"
	RemediationValidationVerdictNotAttempted   RemediationValidationVerdict = "not_attempted"
	RemediationValidationVerdictNotValidated   RemediationValidationVerdict = "not_validated"
)

// RemediationValidationVerdictValues returns all valid RemediationValidationVerdict values.
func RemediationValidationVerdictValues() []RemediationValidationVerdict {
	return []RemediationValidationVerdict{
		RemediationValidationVerdictConfirmed,
		RemediationValidationVerdictRefuted,
		RemediationValidationVerdictInconclusive,
		RemediationValidationVerdictBlockedByScope,
		RemediationValidationVerdictNotAttempted,
		RemediationValidationVerdictNotValidated,
	}
}
