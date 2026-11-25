// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

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
