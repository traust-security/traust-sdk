// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

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
