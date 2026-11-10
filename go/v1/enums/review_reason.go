// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

package enums

// ReviewReason — ReviewItem.reason: the merged layer and triage/remediation review vocabulary. Existing codes retain separate meanings.
type ReviewReason string

const (
	ReviewReasonAmbiguousStatement    ReviewReason = "ambiguous_statement"
	ReviewReasonUnverifiedIdentity    ReviewReason = "unverified_identity"
	ReviewReasonInsufficientAuthority ReviewReason = "insufficient_authority"
	ReviewReasonNoFindingId           ReviewReason = "no_finding_id"
	ReviewReasonUndeterminedFinding   ReviewReason = "undetermined_finding"
	ReviewReasonFpAuditValve          ReviewReason = "fp_audit_valve"
	ReviewReasonNeedsManualTest       ReviewReason = "needs_manual_test"
	ReviewReasonStaleBaseline         ReviewReason = "stale_baseline"
	ReviewReasonRebaselineMapping     ReviewReason = "rebaseline_mapping"
	ReviewReasonRebaselineUnmatched   ReviewReason = "rebaseline_unmatched"
	ReviewReasonUnsoundRefutation     ReviewReason = "unsound_refutation"
	ReviewReasonWeakConfirmation      ReviewReason = "weak_confirmation"
	ReviewReasonUngradedConfirmation  ReviewReason = "ungraded_confirmation"
	ReviewReasonSeverityProposal      ReviewReason = "severity_proposal"
	ReviewReasonNeedsIdentity         ReviewReason = "needs_identity"
	ReviewReasonNewDiscovery          ReviewReason = "new_discovery"
	ReviewReasonSeverityUpgrade       ReviewReason = "severity_upgrade"
	ReviewReasonRegressionDetected    ReviewReason = "regression_detected"
	ReviewReasonEvidenceAvailable     ReviewReason = "evidence_available"
	ReviewReasonRemediationReady      ReviewReason = "remediation_ready"
	ReviewReasonRemediationDue        ReviewReason = "remediation_due"
	ReviewReasonFalsePositiveDispute  ReviewReason = "false_positive_dispute"
	ReviewReasonScopeChange           ReviewReason = "scope_change"
	ReviewReasonRevalidationRequired  ReviewReason = "revalidation_required"
	ReviewReasonSlaEscalation         ReviewReason = "sla_escalation"
	ReviewReasonManualQueue           ReviewReason = "manual_queue"
	ReviewReasonDependencyUpdate      ReviewReason = "dependency_update"
	ReviewReasonPatchAvailable        ReviewReason = "patch_available"
	ReviewReasonAnalystReview         ReviewReason = "analyst_review"
)

// ReviewReasonValues returns all valid ReviewReason values.
func ReviewReasonValues() []ReviewReason {
	return []ReviewReason{
		ReviewReasonAmbiguousStatement,
		ReviewReasonUnverifiedIdentity,
		ReviewReasonInsufficientAuthority,
		ReviewReasonNoFindingId,
		ReviewReasonUndeterminedFinding,
		ReviewReasonFpAuditValve,
		ReviewReasonNeedsManualTest,
		ReviewReasonStaleBaseline,
		ReviewReasonRebaselineMapping,
		ReviewReasonRebaselineUnmatched,
		ReviewReasonUnsoundRefutation,
		ReviewReasonWeakConfirmation,
		ReviewReasonUngradedConfirmation,
		ReviewReasonSeverityProposal,
		ReviewReasonNeedsIdentity,
		ReviewReasonNewDiscovery,
		ReviewReasonSeverityUpgrade,
		ReviewReasonRegressionDetected,
		ReviewReasonEvidenceAvailable,
		ReviewReasonRemediationReady,
		ReviewReasonRemediationDue,
		ReviewReasonFalsePositiveDispute,
		ReviewReasonScopeChange,
		ReviewReasonRevalidationRequired,
		ReviewReasonSlaEscalation,
		ReviewReasonManualQueue,
		ReviewReasonDependencyUpdate,
		ReviewReasonPatchAvailable,
		ReviewReasonAnalystReview,
	}
}
