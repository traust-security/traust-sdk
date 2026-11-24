// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

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
