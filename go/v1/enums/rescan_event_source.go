// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

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
