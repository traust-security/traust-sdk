// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

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
