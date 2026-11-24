// Code generated from traust-contracts 58803ea798ad6431e10da337eb3a28349c1c9992. DO NOT EDIT.

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
