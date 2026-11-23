// Code generated from traust-contracts 8f684d27d978e53c0ffc2b43af379fc1284c26eb. DO NOT EDIT.

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
