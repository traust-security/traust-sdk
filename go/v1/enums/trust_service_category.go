// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

package enums

// TrustServiceCategory — REQUIRED for SOC 2 runs: declared engagement categories (security mandatory)
type TrustServiceCategory string

const (
	TrustServiceCategorySecurity            TrustServiceCategory = "security"
	TrustServiceCategoryAvailability        TrustServiceCategory = "availability"
	TrustServiceCategoryConfidentiality     TrustServiceCategory = "confidentiality"
	TrustServiceCategoryProcessingIntegrity TrustServiceCategory = "processing_integrity"
	TrustServiceCategoryPrivacy             TrustServiceCategory = "privacy"
)

// TrustServiceCategoryValues returns all valid TrustServiceCategory values.
func TrustServiceCategoryValues() []TrustServiceCategory {
	return []TrustServiceCategory{
		TrustServiceCategorySecurity,
		TrustServiceCategoryAvailability,
		TrustServiceCategoryConfidentiality,
		TrustServiceCategoryProcessingIntegrity,
		TrustServiceCategoryPrivacy,
	}
}
