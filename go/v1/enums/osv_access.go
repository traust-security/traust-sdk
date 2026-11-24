// Code generated from traust-contracts 58803ea798ad6431e10da337eb3a28349c1c9992. DO NOT EDIT.

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
