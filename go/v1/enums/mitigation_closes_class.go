// Code generated from traust-contracts ac517047c6566b019e63b911f6b7522000826b48. DO NOT EDIT.

package enums

// MitigationClosesClass — `no` is worthwhile defence in depth that does not close the class -- recorded rather than dropped.
type MitigationClosesClass string

const (
	MitigationClosesClassYes     MitigationClosesClass = "yes"
	MitigationClosesClassPartial MitigationClosesClass = "partial"
	MitigationClosesClassNo      MitigationClosesClass = "no"
)

// MitigationClosesClassValues returns all valid MitigationClosesClass values.
func MitigationClosesClassValues() []MitigationClosesClass {
	return []MitigationClosesClass{
		MitigationClosesClassYes,
		MitigationClosesClassPartial,
		MitigationClosesClassNo,
	}
}
