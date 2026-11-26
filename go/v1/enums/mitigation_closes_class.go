// Code generated from traust-contracts 485578cbe584ad224d57b2c7c778174c0b66ba7c. DO NOT EDIT.

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
