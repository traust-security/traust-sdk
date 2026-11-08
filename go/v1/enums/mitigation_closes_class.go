// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

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
