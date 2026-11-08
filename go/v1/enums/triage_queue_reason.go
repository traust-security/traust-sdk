// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

package enums

// TriageQueueReason — Historical forward triage/remediation reasons. The merged review_reason vocabulary retains every code and meaning.
type TriageQueueReason string

const (
	TriageQueueReasonNewDiscovery         TriageQueueReason = "new_discovery"
	TriageQueueReasonSeverityUpgrade      TriageQueueReason = "severity_upgrade"
	TriageQueueReasonRegressionDetected   TriageQueueReason = "regression_detected"
	TriageQueueReasonEvidenceAvailable    TriageQueueReason = "evidence_available"
	TriageQueueReasonRemediationReady     TriageQueueReason = "remediation_ready"
	TriageQueueReasonRemediationDue       TriageQueueReason = "remediation_due"
	TriageQueueReasonFalsePositiveDispute TriageQueueReason = "false_positive_dispute"
	TriageQueueReasonScopeChange          TriageQueueReason = "scope_change"
	TriageQueueReasonRevalidationRequired TriageQueueReason = "revalidation_required"
	TriageQueueReasonSlaEscalation        TriageQueueReason = "sla_escalation"
	TriageQueueReasonManualQueue          TriageQueueReason = "manual_queue"
	TriageQueueReasonDependencyUpdate     TriageQueueReason = "dependency_update"
	TriageQueueReasonPatchAvailable       TriageQueueReason = "patch_available"
	TriageQueueReasonAnalystReview        TriageQueueReason = "analyst_review"
)

// TriageQueueReasonValues returns all valid TriageQueueReason values.
func TriageQueueReasonValues() []TriageQueueReason {
	return []TriageQueueReason{
		TriageQueueReasonNewDiscovery,
		TriageQueueReasonSeverityUpgrade,
		TriageQueueReasonRegressionDetected,
		TriageQueueReasonEvidenceAvailable,
		TriageQueueReasonRemediationReady,
		TriageQueueReasonRemediationDue,
		TriageQueueReasonFalsePositiveDispute,
		TriageQueueReasonScopeChange,
		TriageQueueReasonRevalidationRequired,
		TriageQueueReasonSlaEscalation,
		TriageQueueReasonManualQueue,
		TriageQueueReasonDependencyUpdate,
		TriageQueueReasonPatchAvailable,
		TriageQueueReasonAnalystReview,
	}
}
