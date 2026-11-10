// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

package enums

// RescanEventSource — Operational vocabulary for traust_ops.run_order.event_source; Core enforces new writes, while stored historical text remains readable.
type RescanEventSource string

const (
	RescanEventSourceExternalReport RescanEventSource = "external-report"
	RescanEventSourceMethodology    RescanEventSource = "methodology"
	RescanEventSourceCve            RescanEventSource = "cve"
	RescanEventSourceRelease        RescanEventSource = "release"
)

// RescanEventSourceValues returns all valid RescanEventSource values.
func RescanEventSourceValues() []RescanEventSource {
	return []RescanEventSource{
		RescanEventSourceExternalReport,
		RescanEventSourceMethodology,
		RescanEventSourceCve,
		RescanEventSourceRelease,
	}
}
