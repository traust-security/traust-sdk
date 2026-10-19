// Code generated from traust-contracts 3ed68158121322ce06525de170aa383e057b6a7b. DO NOT EDIT.

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
