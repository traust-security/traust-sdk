// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

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
