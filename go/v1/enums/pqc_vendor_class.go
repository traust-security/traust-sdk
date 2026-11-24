// Code generated from traust-contracts 369a2ab2c9b65225c6878eca75bada6a65df548f. DO NOT EDIT.

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
