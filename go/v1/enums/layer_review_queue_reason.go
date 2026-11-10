// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

package enums

// LayerReviewQueueReason — Historical layer review reasons, retained with their original codes. The merged review_reason vocabulary now owns ReviewItem.reason and layer queue_reason.
type LayerReviewQueueReason string

const (
	LayerReviewQueueReasonAmbiguousStatement    LayerReviewQueueReason = "ambiguous_statement"
	LayerReviewQueueReasonUnverifiedIdentity    LayerReviewQueueReason = "unverified_identity"
	LayerReviewQueueReasonInsufficientAuthority LayerReviewQueueReason = "insufficient_authority"
	LayerReviewQueueReasonNoFindingId           LayerReviewQueueReason = "no_finding_id"
	LayerReviewQueueReasonUndeterminedFinding   LayerReviewQueueReason = "undetermined_finding"
	LayerReviewQueueReasonFpAuditValve          LayerReviewQueueReason = "fp_audit_valve"
	LayerReviewQueueReasonNeedsManualTest       LayerReviewQueueReason = "needs_manual_test"
	LayerReviewQueueReasonStaleBaseline         LayerReviewQueueReason = "stale_baseline"
	LayerReviewQueueReasonRebaselineMapping     LayerReviewQueueReason = "rebaseline_mapping"
	LayerReviewQueueReasonRebaselineUnmatched   LayerReviewQueueReason = "rebaseline_unmatched"
	LayerReviewQueueReasonUnsoundRefutation     LayerReviewQueueReason = "unsound_refutation"
	LayerReviewQueueReasonWeakConfirmation      LayerReviewQueueReason = "weak_confirmation"
	LayerReviewQueueReasonUngradedConfirmation  LayerReviewQueueReason = "ungraded_confirmation"
	LayerReviewQueueReasonSeverityProposal      LayerReviewQueueReason = "severity_proposal"
	LayerReviewQueueReasonNeedsIdentity         LayerReviewQueueReason = "needs_identity"
)

// LayerReviewQueueReasonValues returns all valid LayerReviewQueueReason values.
func LayerReviewQueueReasonValues() []LayerReviewQueueReason {
	return []LayerReviewQueueReason{
		LayerReviewQueueReasonAmbiguousStatement,
		LayerReviewQueueReasonUnverifiedIdentity,
		LayerReviewQueueReasonInsufficientAuthority,
		LayerReviewQueueReasonNoFindingId,
		LayerReviewQueueReasonUndeterminedFinding,
		LayerReviewQueueReasonFpAuditValve,
		LayerReviewQueueReasonNeedsManualTest,
		LayerReviewQueueReasonStaleBaseline,
		LayerReviewQueueReasonRebaselineMapping,
		LayerReviewQueueReasonRebaselineUnmatched,
		LayerReviewQueueReasonUnsoundRefutation,
		LayerReviewQueueReasonWeakConfirmation,
		LayerReviewQueueReasonUngradedConfirmation,
		LayerReviewQueueReasonSeverityProposal,
		LayerReviewQueueReasonNeedsIdentity,
	}
}
