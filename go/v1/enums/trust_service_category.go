// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

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
