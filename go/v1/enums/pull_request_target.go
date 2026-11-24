// Code generated from traust-contracts 369a2ab2c9b65225c6878eca75bada6a65df548f. DO NOT EDIT.

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
