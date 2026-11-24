// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

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
