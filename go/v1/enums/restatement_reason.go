// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

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
