// Code generated from traust-contracts 31d0d884672149e606934c2844ffcbf3bdd4b925. DO NOT EDIT.

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
