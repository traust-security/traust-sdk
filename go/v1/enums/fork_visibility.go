// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

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
