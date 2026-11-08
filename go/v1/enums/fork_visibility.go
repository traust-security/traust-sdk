// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

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
