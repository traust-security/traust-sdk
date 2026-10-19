// Code generated from traust-contracts 3ed68158121322ce06525de170aa383e057b6a7b. DO NOT EDIT.

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
