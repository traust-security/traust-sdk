// Code generated from traust-contracts 8f684d27d978e53c0ffc2b43af379fc1284c26eb. DO NOT EDIT.

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
