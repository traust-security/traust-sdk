// Code generated from traust-contracts 31d0d884672149e606934c2844ffcbf3bdd4b925. DO NOT EDIT.

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
