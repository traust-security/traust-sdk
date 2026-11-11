// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

package enums

// DecisionReason — Decision.reason: why a judgement changed or which existing remediation/verification step produced the statement. Operational reason labels preserve source intent and statement identity; they do not supply evidence or assurance.
type DecisionReason string

const (
	DecisionReasonRegression                 DecisionReason = "regression"
	DecisionReasonDataError                  DecisionReason = "data_error"
	DecisionReasonSchemaMigration            DecisionReason = "schema_migration"
	DecisionReasonBaselineRewrite            DecisionReason = "baseline_rewrite"
	DecisionReasonOperatorError              DecisionReason = "operator_error"
	DecisionReasonCandidate                  DecisionReason = "candidate"
	DecisionReasonForked                     DecisionReason = "forked"
	DecisionReasonPatched                    DecisionReason = "patched"
	DecisionReasonChecksPassed               DecisionReason = "checks_passed"
	DecisionReasonChecksFailed               DecisionReason = "checks_failed"
	DecisionReasonRevalidatedFixed           DecisionReason = "revalidated_fixed"
	DecisionReasonRevalidatedStillVulnerable DecisionReason = "revalidated_still_vulnerable"
	DecisionReasonPrOpened                   DecisionReason = "pr_opened"
	DecisionReasonMerged                     DecisionReason = "merged"
	DecisionReasonAbandoned                  DecisionReason = "abandoned"
	DecisionReasonResolved                   DecisionReason = "resolved"
	DecisionReasonPartiallyResolved          DecisionReason = "partially_resolved"
	DecisionReasonUnresolved                 DecisionReason = "unresolved"
	DecisionReasonNewApproach                DecisionReason = "new_approach"
	DecisionReasonFalsePositive              DecisionReason = "false_positive"
	DecisionReasonRiskAccepted               DecisionReason = "risk_accepted"
)

// DecisionReasonValues returns all valid DecisionReason values.
func DecisionReasonValues() []DecisionReason {
	return []DecisionReason{
		DecisionReasonRegression,
		DecisionReasonDataError,
		DecisionReasonSchemaMigration,
		DecisionReasonBaselineRewrite,
		DecisionReasonOperatorError,
		DecisionReasonCandidate,
		DecisionReasonForked,
		DecisionReasonPatched,
		DecisionReasonChecksPassed,
		DecisionReasonChecksFailed,
		DecisionReasonRevalidatedFixed,
		DecisionReasonRevalidatedStillVulnerable,
		DecisionReasonPrOpened,
		DecisionReasonMerged,
		DecisionReasonAbandoned,
		DecisionReasonResolved,
		DecisionReasonPartiallyResolved,
		DecisionReasonUnresolved,
		DecisionReasonNewApproach,
		DecisionReasonFalsePositive,
		DecisionReasonRiskAccepted,
	}
}
