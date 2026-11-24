// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

package enums

// ExternalRefSystem — Values of `system` in layer.schema.json.
type ExternalRefSystem string

const (
	ExternalRefSystemCve      ExternalRefSystem = "cve"
	ExternalRefSystemBugzilla ExternalRefSystem = "bugzilla"
	ExternalRefSystemGhsa     ExternalRefSystem = "ghsa"
	ExternalRefSystemJira     ExternalRefSystem = "jira"
)

// ExternalRefSystemValues returns all valid ExternalRefSystem values.
func ExternalRefSystemValues() []ExternalRefSystem {
	return []ExternalRefSystem{
		ExternalRefSystemCve,
		ExternalRefSystemBugzilla,
		ExternalRefSystemGhsa,
		ExternalRefSystemJira,
	}
}
