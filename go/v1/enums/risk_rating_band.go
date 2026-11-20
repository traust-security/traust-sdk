// Code generated from traust-contracts 15bded2e3e4b6c70bd1e20ef8df50c7b66b5ebc1. DO NOT EDIT.

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
