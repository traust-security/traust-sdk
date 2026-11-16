// Code generated from traust-contracts 31d0d884672149e606934c2844ffcbf3bdd4b925. DO NOT EDIT.

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
