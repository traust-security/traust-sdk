// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

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
