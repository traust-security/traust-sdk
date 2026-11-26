// Code generated from traust-contracts 485578cbe584ad224d57b2c7c778174c0b66ba7c. DO NOT EDIT.

package enums

// PqcVendorStatus — A vendor's post-quantum status as reviewed.
type PqcVendorStatus string

const (
	PqcVendorStatusAvailable  PqcVendorStatus = "available"
	PqcVendorStatusPartial    PqcVendorStatus = "partial"
	PqcVendorStatusInProgress PqcVendorStatus = "in_progress"
	PqcVendorStatusNone       PqcVendorStatus = "none"
)

// PqcVendorStatusValues returns all valid PqcVendorStatus values.
func PqcVendorStatusValues() []PqcVendorStatus {
	return []PqcVendorStatus{
		PqcVendorStatusAvailable,
		PqcVendorStatusPartial,
		PqcVendorStatusInProgress,
		PqcVendorStatusNone,
	}
}
