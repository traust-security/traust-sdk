// Code generated from traust-contracts 58803ea798ad6431e10da337eb3a28349c1c9992. DO NOT EDIT.

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
