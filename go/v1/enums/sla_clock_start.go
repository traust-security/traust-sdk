// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

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
