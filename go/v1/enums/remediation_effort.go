// Code generated from traust-contracts 15bded2e3e4b6c70bd1e20ef8df50c7b66b5ebc1. DO NOT EDIT.

package enums

// RemediationEffort — Values of `remediation_effort` in pqc-blockers.schema.json.
type RemediationEffort string

const (
	RemediationEffortTrivial         RemediationEffort = "trivial"
	RemediationEffortModerate        RemediationEffort = "moderate"
	RemediationEffortSignificant     RemediationEffort = "significant"
	RemediationEffortBlockedExternal RemediationEffort = "blocked-external"
	RemediationEffortXs              RemediationEffort = "xs"
	RemediationEffortS               RemediationEffort = "s"
	RemediationEffortM               RemediationEffort = "m"
	RemediationEffortL               RemediationEffort = "l"
	RemediationEffortXl              RemediationEffort = "xl"
)

// RemediationEffortValues returns all valid RemediationEffort values.
func RemediationEffortValues() []RemediationEffort {
	return []RemediationEffort{
		RemediationEffortTrivial,
		RemediationEffortModerate,
		RemediationEffortSignificant,
		RemediationEffortBlockedExternal,
		RemediationEffortXs,
		RemediationEffortS,
		RemediationEffortM,
		RemediationEffortL,
		RemediationEffortXl,
	}
}
