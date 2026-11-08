// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

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
