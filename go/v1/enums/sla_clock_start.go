// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

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
