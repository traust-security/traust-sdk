// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

package enums

// RemediationCheckOutcome — Values of `outcome` in remediation.schema.json.
type RemediationCheckOutcome string

const (
	RemediationCheckOutcomePass  RemediationCheckOutcome = "pass"
	RemediationCheckOutcomeFail  RemediationCheckOutcome = "fail"
	RemediationCheckOutcomeSkip  RemediationCheckOutcome = "skip"
	RemediationCheckOutcomeError RemediationCheckOutcome = "error"
)

// RemediationCheckOutcomeValues returns all valid RemediationCheckOutcome values.
func RemediationCheckOutcomeValues() []RemediationCheckOutcome {
	return []RemediationCheckOutcome{
		RemediationCheckOutcomePass,
		RemediationCheckOutcomeFail,
		RemediationCheckOutcomeSkip,
		RemediationCheckOutcomeError,
	}
}
