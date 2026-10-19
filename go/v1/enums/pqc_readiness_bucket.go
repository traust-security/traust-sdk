// Code generated from traust-contracts 3ed68158121322ce06525de170aa383e057b6a7b. DO NOT EDIT.

package enums

// PqcReadinessBucket — Values of `bucket` in pqc-decision-tree.schema.json.
type PqcReadinessBucket string

const (
	PqcReadinessBucketReady           PqcReadinessBucket = "ready"
	PqcReadinessBucketPartial         PqcReadinessBucket = "partial"
	PqcReadinessBucketNotReady        PqcReadinessBucket = "not-ready"
	PqcReadinessBucketBlockedExternal PqcReadinessBucket = "blocked-external"
	PqcReadinessBucketNotApplicable   PqcReadinessBucket = "not-applicable"
)

// PqcReadinessBucketValues returns all valid PqcReadinessBucket values.
func PqcReadinessBucketValues() []PqcReadinessBucket {
	return []PqcReadinessBucket{
		PqcReadinessBucketReady,
		PqcReadinessBucketPartial,
		PqcReadinessBucketNotReady,
		PqcReadinessBucketBlockedExternal,
		PqcReadinessBucketNotApplicable,
	}
}
