// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

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
