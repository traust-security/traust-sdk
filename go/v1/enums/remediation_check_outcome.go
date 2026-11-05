// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

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
