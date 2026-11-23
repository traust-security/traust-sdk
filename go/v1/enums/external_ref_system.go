// Code generated from traust-contracts 8f684d27d978e53c0ffc2b43af379fc1284c26eb. DO NOT EDIT.

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
