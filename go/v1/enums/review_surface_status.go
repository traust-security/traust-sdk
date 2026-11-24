// Code generated from traust-contracts 369a2ab2c9b65225c6878eca75bada6a65df548f. DO NOT EDIT.

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
