// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

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
