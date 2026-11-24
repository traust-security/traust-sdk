// Code generated from traust-contracts c2f6e6fb12270f8b3c8703da7bf39bd01776e4b1. DO NOT EDIT.

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
