// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

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
