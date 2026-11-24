// Code generated from traust-contracts c2f6e6fb12270f8b3c8703da7bf39bd01776e4b1. DO NOT EDIT.

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
