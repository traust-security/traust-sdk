// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

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
