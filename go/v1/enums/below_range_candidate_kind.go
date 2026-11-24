// Code generated from traust-contracts 369a2ab2c9b65225c6878eca75bada6a65df548f. DO NOT EDIT.

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
