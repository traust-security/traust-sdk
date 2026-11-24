// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

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
