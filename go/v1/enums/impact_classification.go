// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

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
