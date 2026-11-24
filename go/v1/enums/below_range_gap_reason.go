// Code generated from traust-contracts c2f6e6fb12270f8b3c8703da7bf39bd01776e4b1. DO NOT EDIT.

package enums

// BelowRangeGapReason — Why a below-range candidate or lookup couldn't be checked. A coverage gap is never clean.
type BelowRangeGapReason string

const (
	BelowRangeGapReasonNugetIdUnresolved  BelowRangeGapReason = "nuget_id_unresolved"
	BelowRangeGapReasonOsvUnanswered      BelowRangeGapReason = "osv_unanswered"
	BelowRangeGapReasonAdvisoryUnreadable BelowRangeGapReason = "advisory_unreadable"
	BelowRangeGapReasonOverMaxCandidates  BelowRangeGapReason = "over_max_candidates"
	BelowRangeGapReasonSourceUnavailable  BelowRangeGapReason = "source_unavailable"
	BelowRangeGapReasonNoUsableFixPatch   BelowRangeGapReason = "no_usable_fix_patch"
)

// BelowRangeGapReasonValues returns all valid BelowRangeGapReason values.
func BelowRangeGapReasonValues() []BelowRangeGapReason {
	return []BelowRangeGapReason{
		BelowRangeGapReasonNugetIdUnresolved,
		BelowRangeGapReasonOsvUnanswered,
		BelowRangeGapReasonAdvisoryUnreadable,
		BelowRangeGapReasonOverMaxCandidates,
		BelowRangeGapReasonSourceUnavailable,
		BelowRangeGapReasonNoUsableFixPatch,
	}
}
