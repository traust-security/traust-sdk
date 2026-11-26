// Code generated from traust-contracts 485578cbe584ad224d57b2c7c778174c0b66ba7c. DO NOT EDIT.

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
