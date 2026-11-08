// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

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
