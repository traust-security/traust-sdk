// Code generated from traust-contracts 2e521561f022fe839dd828ddf2185ba3a0a1045f. DO NOT EDIT.

package enums

// RestatementReason — Why the data was wrong.
type RestatementReason string

const (
	RestatementReasonDataError       RestatementReason = "data_error"
	RestatementReasonSchemaMigration RestatementReason = "schema_migration"
	RestatementReasonBaselineRewrite RestatementReason = "baseline_rewrite"
	RestatementReasonOperatorError   RestatementReason = "operator_error"
)

// RestatementReasonValues returns all valid RestatementReason values.
func RestatementReasonValues() []RestatementReason {
	return []RestatementReason{
		RestatementReasonDataError,
		RestatementReasonSchemaMigration,
		RestatementReasonBaselineRewrite,
		RestatementReasonOperatorError,
	}
}
