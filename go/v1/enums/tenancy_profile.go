// Code generated from traust-contracts ac517047c6566b019e63b911f6b7522000826b48. DO NOT EDIT.

package enums

// TenancyProfile — Values of `tenancy_profile` in layer.schema.json.
type TenancyProfile string

const (
	TenancyProfileMultiTenant  TenancyProfile = "multi_tenant"
	TenancyProfileSingleTenant TenancyProfile = "single_tenant"
)

// TenancyProfileValues returns all valid TenancyProfile values.
func TenancyProfileValues() []TenancyProfile {
	return []TenancyProfile{
		TenancyProfileMultiTenant,
		TenancyProfileSingleTenant,
	}
}
