// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

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
