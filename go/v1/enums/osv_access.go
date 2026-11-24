// Code generated from traust-contracts 369a2ab2c9b65225c6878eca75bada6a65df548f. DO NOT EDIT.

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
