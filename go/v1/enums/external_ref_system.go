// Code generated from traust-contracts ac517047c6566b019e63b911f6b7522000826b48. DO NOT EDIT.

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
