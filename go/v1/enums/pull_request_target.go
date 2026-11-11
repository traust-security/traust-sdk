// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

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
