// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

package enums

// LivenessStatus — Forge-observed repository availability, distinct from recency of activity.
type LivenessStatus string

const (
	LivenessStatusActive   LivenessStatus = "active"
	LivenessStatusArchived LivenessStatus = "archived"
	LivenessStatusMoved    LivenessStatus = "moved"
	LivenessStatusMissing  LivenessStatus = "missing"
	LivenessStatusDisabled LivenessStatus = "disabled"
	LivenessStatusUnknown  LivenessStatus = "unknown"
)

// LivenessStatusValues returns all valid LivenessStatus values.
func LivenessStatusValues() []LivenessStatus {
	return []LivenessStatus{
		LivenessStatusActive,
		LivenessStatusArchived,
		LivenessStatusMoved,
		LivenessStatusMissing,
		LivenessStatusDisabled,
		LivenessStatusUnknown,
	}
}
