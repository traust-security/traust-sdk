// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

package enums

// InterfaceExposure — Values of `exposure` in isolation-review.schema.json.
type InterfaceExposure string

const (
	InterfaceExposurePublic   InterfaceExposure = "public"
	InterfaceExposureTenant   InterfaceExposure = "tenant"
	InterfaceExposurePartner  InterfaceExposure = "partner"
	InterfaceExposureInternal InterfaceExposure = "internal"
)

// InterfaceExposureValues returns all valid InterfaceExposure values.
func InterfaceExposureValues() []InterfaceExposure {
	return []InterfaceExposure{
		InterfaceExposurePublic,
		InterfaceExposureTenant,
		InterfaceExposurePartner,
		InterfaceExposureInternal,
	}
}
