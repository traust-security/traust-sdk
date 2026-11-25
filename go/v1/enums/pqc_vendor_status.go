// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

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
