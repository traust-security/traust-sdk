// Code generated from traust-contracts 2e521561f022fe839dd828ddf2185ba3a0a1045f. DO NOT EDIT.

package enums

// ForkVisibility — Values of `visibility` in remediation.schema.json.
type ForkVisibility string

const (
	ForkVisibilityPrivate  ForkVisibility = "private"
	ForkVisibilityInternal ForkVisibility = "internal"
	ForkVisibilityPublic   ForkVisibility = "public"
)

// ForkVisibilityValues returns all valid ForkVisibility values.
func ForkVisibilityValues() []ForkVisibility {
	return []ForkVisibility{
		ForkVisibilityPrivate,
		ForkVisibilityInternal,
		ForkVisibilityPublic,
	}
}
