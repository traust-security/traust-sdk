// Code generated from traust-contracts 15bded2e3e4b6c70bd1e20ef8df50c7b66b5ebc1. DO NOT EDIT.

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
