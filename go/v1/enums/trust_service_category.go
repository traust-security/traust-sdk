// Code generated from traust-contracts ac517047c6566b019e63b911f6b7522000826b48. DO NOT EDIT.

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
