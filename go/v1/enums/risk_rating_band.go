// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

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
