// Code generated from traust-contracts 485578cbe584ad224d57b2c7c778174c0b66ba7c. DO NOT EDIT.

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
