// Code generated from traust-contracts c2f6e6fb12270f8b3c8703da7bf39bd01776e4b1. DO NOT EDIT.

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
