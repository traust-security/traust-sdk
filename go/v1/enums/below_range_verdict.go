// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

package enums

// BelowRangeVerdict — A review's answer to: is the installed version affected (the vulnerable code present and the fix absent) despite the stated range? Only applicable becomes a finding (D1).
type BelowRangeVerdict string

const (
	BelowRangeVerdictApplicable    BelowRangeVerdict = "applicable"
	BelowRangeVerdictNotApplicable BelowRangeVerdict = "not_applicable"
	BelowRangeVerdictUndetermined  BelowRangeVerdict = "undetermined"
)

// BelowRangeVerdictValues returns all valid BelowRangeVerdict values.
func BelowRangeVerdictValues() []BelowRangeVerdict {
	return []BelowRangeVerdict{
		BelowRangeVerdictApplicable,
		BelowRangeVerdictNotApplicable,
		BelowRangeVerdictUndetermined,
	}
}
