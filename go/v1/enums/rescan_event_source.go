// Code generated from traust-contracts c2f6e6fb12270f8b3c8703da7bf39bd01776e4b1. DO NOT EDIT.

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
