// Code generated from traust-contracts 31d0d884672149e606934c2844ffcbf3bdd4b925. DO NOT EDIT.

package enums

// CloudConfigFindingStatus — Values of `status` in cloud-config-audit.schema.json.
type CloudConfigFindingStatus string

const (
	CloudConfigFindingStatusConfirmed   CloudConfigFindingStatus = "confirmed"
	CloudConfigFindingStatusSuppressed  CloudConfigFindingStatus = "suppressed"
	CloudConfigFindingStatusNeedsReview CloudConfigFindingStatus = "needs_review"
)

// CloudConfigFindingStatusValues returns all valid CloudConfigFindingStatus values.
func CloudConfigFindingStatusValues() []CloudConfigFindingStatus {
	return []CloudConfigFindingStatus{
		CloudConfigFindingStatusConfirmed,
		CloudConfigFindingStatusSuppressed,
		CloudConfigFindingStatusNeedsReview,
	}
}
