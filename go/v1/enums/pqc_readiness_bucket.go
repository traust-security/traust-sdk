// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

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
