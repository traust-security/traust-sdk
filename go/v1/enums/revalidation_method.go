// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

package enums

// RevalidationMethod — Values of `method` in remediation.schema.json.
type RevalidationMethod string

const (
	RevalidationMethodValidateLive     RevalidationMethod = "validate-live"
	RevalidationMethodValidatePlatform RevalidationMethod = "validate-platform"
	RevalidationMethodValidateFindings RevalidationMethod = "validate-findings"
	RevalidationMethodManual           RevalidationMethod = "manual"
	RevalidationMethodNone             RevalidationMethod = "none"
)

// RevalidationMethodValues returns all valid RevalidationMethod values.
func RevalidationMethodValues() []RevalidationMethod {
	return []RevalidationMethod{
		RevalidationMethodValidateLive,
		RevalidationMethodValidatePlatform,
		RevalidationMethodValidateFindings,
		RevalidationMethodManual,
		RevalidationMethodNone,
	}
}
