// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

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
