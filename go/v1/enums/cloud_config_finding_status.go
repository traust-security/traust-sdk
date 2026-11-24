// Code generated from traust-contracts c2f6e6fb12270f8b3c8703da7bf39bd01776e4b1. DO NOT EDIT.

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
