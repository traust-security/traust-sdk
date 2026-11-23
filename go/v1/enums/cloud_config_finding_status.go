// Code generated from traust-contracts 8f684d27d978e53c0ffc2b43af379fc1284c26eb. DO NOT EDIT.

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
