// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

package enums

// BelowRangeCandidateKind — Where an installed version sits outside an advisory's stated spans: below every explicit start, or in a gap after a span left without a fix.
type BelowRangeCandidateKind string

const (
	BelowRangeCandidateKindBelowRange BelowRangeCandidateKind = "below_range"
	BelowRangeCandidateKindGap        BelowRangeCandidateKind = "gap"
)

// BelowRangeCandidateKindValues returns all valid BelowRangeCandidateKind values.
func BelowRangeCandidateKindValues() []BelowRangeCandidateKind {
	return []BelowRangeCandidateKind{
		BelowRangeCandidateKindBelowRange,
		BelowRangeCandidateKindGap,
	}
}
