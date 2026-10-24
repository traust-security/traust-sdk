// Code generated from traust-contracts 2e521561f022fe839dd828ddf2185ba3a0a1045f. DO NOT EDIT.

package enums

// RoadmapPriority — Priority labels for explicitly selected future-write roadmap contracts.
type RoadmapPriority string

const (
	RoadmapPriorityP0 RoadmapPriority = "p0"
	RoadmapPriorityP1 RoadmapPriority = "p1"
	RoadmapPriorityP2 RoadmapPriority = "p2"
	RoadmapPriorityP3 RoadmapPriority = "p3"
	RoadmapPriorityP4 RoadmapPriority = "p4"
)

// RoadmapPriorityValues returns all valid RoadmapPriority values.
func RoadmapPriorityValues() []RoadmapPriority {
	return []RoadmapPriority{
		RoadmapPriorityP0,
		RoadmapPriorityP1,
		RoadmapPriorityP2,
		RoadmapPriorityP3,
		RoadmapPriorityP4,
	}
}
