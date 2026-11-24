// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

package enums

// ReviewSurfaceStatus — How an audit or a follow-up scan answered one review surface.
type ReviewSurfaceStatus string

const (
	ReviewSurfaceStatusFinding       ReviewSurfaceStatus = "finding"
	ReviewSurfaceStatusNegative      ReviewSurfaceStatus = "negative"
	ReviewSurfaceStatusNotApplicable ReviewSurfaceStatus = "not_applicable"
	ReviewSurfaceStatusNotAssessed   ReviewSurfaceStatus = "not_assessed"
	ReviewSurfaceStatusOutOfScope    ReviewSurfaceStatus = "out_of_scope"
)

// ReviewSurfaceStatusValues returns all valid ReviewSurfaceStatus values.
func ReviewSurfaceStatusValues() []ReviewSurfaceStatus {
	return []ReviewSurfaceStatus{
		ReviewSurfaceStatusFinding,
		ReviewSurfaceStatusNegative,
		ReviewSurfaceStatusNotApplicable,
		ReviewSurfaceStatusNotAssessed,
		ReviewSurfaceStatusOutOfScope,
	}
}
