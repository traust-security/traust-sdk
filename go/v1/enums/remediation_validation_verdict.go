// Code generated from traust-contracts 31d0d884672149e606934c2844ffcbf3bdd4b925. DO NOT EDIT.

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
