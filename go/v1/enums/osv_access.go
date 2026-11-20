// Code generated from traust-contracts 15bded2e3e4b6c70bd1e20ef8df50c7b66b5ebc1. DO NOT EDIT.

package enums

// OsvAccess — How OSV advisories are read: api queries the OSV API; air-gapped reads an adopter-provided osv-scanner and offline database.
type OsvAccess string

const (
	OsvAccessApi       OsvAccess = "api"
	OsvAccessAirGapped OsvAccess = "air-gapped"
)

// OsvAccessValues returns all valid OsvAccess values.
func OsvAccessValues() []OsvAccess {
	return []OsvAccess{
		OsvAccessApi,
		OsvAccessAirGapped,
	}
}
