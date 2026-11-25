// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

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
