// Code generated from traust-contracts 485578cbe584ad224d57b2c7c778174c0b66ba7c. DO NOT EDIT.

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
