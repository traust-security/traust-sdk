// Code generated from traust-contracts 369a2ab2c9b65225c6878eca75bada6a65df548f. DO NOT EDIT.

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
