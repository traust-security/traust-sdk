// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

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
