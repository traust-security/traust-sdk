// Code generated from traust-contracts 2e521561f022fe839dd828ddf2185ba3a0a1045f. DO NOT EDIT.

package enums

// ImpactClassification — Impact analysis repo classification.
type ImpactClassification string

const (
	ImpactClassificationAffected          ImpactClassification = "affected"
	ImpactClassificationLikelyAffected    ImpactClassification = "likely_affected"
	ImpactClassificationNotObserved       ImpactClassification = "not_observed"
	ImpactClassificationVersionNotInRange ImpactClassification = "version_not_in_range"
	ImpactClassificationNotImported       ImpactClassification = "not_imported"
	ImpactClassificationInconclusive      ImpactClassification = "inconclusive"
)

// ImpactClassificationValues returns all valid ImpactClassification values.
func ImpactClassificationValues() []ImpactClassification {
	return []ImpactClassification{
		ImpactClassificationAffected,
		ImpactClassificationLikelyAffected,
		ImpactClassificationNotObserved,
		ImpactClassificationVersionNotInRange,
		ImpactClassificationNotImported,
		ImpactClassificationInconclusive,
	}
}
