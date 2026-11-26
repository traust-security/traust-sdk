// Code generated from traust-contracts 485578cbe584ad224d57b2c7c778174c0b66ba7c. DO NOT EDIT.

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
