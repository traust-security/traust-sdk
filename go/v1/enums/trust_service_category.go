// Code generated from traust-contracts c2f6e6fb12270f8b3c8703da7bf39bd01776e4b1. DO NOT EDIT.

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
