// Code generated from traust-contracts 8f684d27d978e53c0ffc2b43af379fc1284c26eb. DO NOT EDIT.

package enums

// PqcVendorClass — What kind of vendor or dependency a PQC vendor claim is about.
type PqcVendorClass string

const (
	PqcVendorClassProvider  PqcVendorClass = "provider"
	PqcVendorClassLibrary   PqcVendorClass = "library"
	PqcVendorClassEcosystem PqcVendorClass = "ecosystem"
	PqcVendorClassService   PqcVendorClass = "service"
	PqcVendorClassInternal  PqcVendorClass = "internal"
)

// PqcVendorClassValues returns all valid PqcVendorClass values.
func PqcVendorClassValues() []PqcVendorClass {
	return []PqcVendorClass{
		PqcVendorClassProvider,
		PqcVendorClassLibrary,
		PqcVendorClassEcosystem,
		PqcVendorClassService,
		PqcVendorClassInternal,
	}
}
