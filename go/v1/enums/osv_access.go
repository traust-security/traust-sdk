// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

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
