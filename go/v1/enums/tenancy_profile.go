// Code generated from traust-contracts 15bded2e3e4b6c70bd1e20ef8df50c7b66b5ebc1. DO NOT EDIT.

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
