// Code generated from traust-contracts ac517047c6566b019e63b911f6b7522000826b48. DO NOT EDIT.

package enums

// PullRequestTarget — Where the PR is opened.
type PullRequestTarget string

const (
	PullRequestTargetPrivateFork PullRequestTarget = "private-fork"
	PullRequestTargetDownstream  PullRequestTarget = "downstream"
	PullRequestTargetUpstream    PullRequestTarget = "upstream"
)

// PullRequestTargetValues returns all valid PullRequestTarget values.
func PullRequestTargetValues() []PullRequestTarget {
	return []PullRequestTarget{
		PullRequestTargetPrivateFork,
		PullRequestTargetDownstream,
		PullRequestTargetUpstream,
	}
}
