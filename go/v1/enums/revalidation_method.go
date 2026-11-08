// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

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
