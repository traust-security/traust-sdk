// Code generated from traust-contracts 15bded2e3e4b6c70bd1e20ef8df50c7b66b5ebc1. DO NOT EDIT.

package enums

// SlaClockStart — Which timestamp starts a finding's SLA clock.
type SlaClockStart string

const (
	SlaClockStartFirstRoutedOrFiled SlaClockStart = "first_routed_or_filed"
	SlaClockStartFirstEvent         SlaClockStart = "first_event"
	SlaClockStartAuditDate          SlaClockStart = "audit_date"
)

// SlaClockStartValues returns all valid SlaClockStart values.
func SlaClockStartValues() []SlaClockStart {
	return []SlaClockStart{
		SlaClockStartFirstRoutedOrFiled,
		SlaClockStartFirstEvent,
		SlaClockStartAuditDate,
	}
}
