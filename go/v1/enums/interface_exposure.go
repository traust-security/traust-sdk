// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

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
