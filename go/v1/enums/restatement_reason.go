// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

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
