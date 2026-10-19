// Code generated from traust-contracts 3ed68158121322ce06525de170aa383e057b6a7b. DO NOT EDIT.

package enums

// RiskRatingBand — Values of `band` in risk-rating-methodology.schema.json.
type RiskRatingBand string

const (
	RiskRatingBandCritical RiskRatingBand = "critical"
	RiskRatingBandHigh     RiskRatingBand = "high"
	RiskRatingBandMedium   RiskRatingBand = "medium"
	RiskRatingBandLow      RiskRatingBand = "low"
	RiskRatingBandNote     RiskRatingBand = "note"
)

// RiskRatingBandValues returns all valid RiskRatingBand values.
func RiskRatingBandValues() []RiskRatingBand {
	return []RiskRatingBand{
		RiskRatingBandCritical,
		RiskRatingBandHigh,
		RiskRatingBandMedium,
		RiskRatingBandLow,
		RiskRatingBandNote,
	}
}
