// Code generated from traust-contracts ac517047c6566b019e63b911f6b7522000826b48. DO NOT EDIT.

package enums

// CvssUserInteraction — Values of `user_interaction` in validation.schema.json.
type CvssUserInteraction string

const (
	CvssUserInteractionNone     CvssUserInteraction = "none"
	CvssUserInteractionRequired CvssUserInteraction = "required"
)

// CvssUserInteractionValues returns all valid CvssUserInteraction values.
func CvssUserInteractionValues() []CvssUserInteraction {
	return []CvssUserInteraction{
		CvssUserInteractionNone,
		CvssUserInteractionRequired,
	}
}
