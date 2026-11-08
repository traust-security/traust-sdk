// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

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
