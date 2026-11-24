// Code generated from traust-contracts 58803ea798ad6431e10da337eb3a28349c1c9992. DO NOT EDIT.

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
