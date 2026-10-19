// Code generated from traust-contracts 3ed68158121322ce06525de170aa383e057b6a7b. DO NOT EDIT.

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
