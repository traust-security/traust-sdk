// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

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
