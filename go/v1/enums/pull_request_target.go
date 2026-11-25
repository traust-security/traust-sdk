// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

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
