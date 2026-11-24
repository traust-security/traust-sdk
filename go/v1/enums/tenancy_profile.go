// Code generated from traust-contracts 369a2ab2c9b65225c6878eca75bada6a65df548f. DO NOT EDIT.

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
