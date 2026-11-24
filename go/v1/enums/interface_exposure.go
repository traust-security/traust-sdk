// Code generated from traust-contracts c2f6e6fb12270f8b3c8703da7bf39bd01776e4b1. DO NOT EDIT.

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
