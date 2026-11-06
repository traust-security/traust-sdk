// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

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
