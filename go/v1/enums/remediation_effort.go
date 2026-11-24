// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

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
