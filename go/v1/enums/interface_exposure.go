// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

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
