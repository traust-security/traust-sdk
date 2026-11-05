// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

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
