// Code generated from traust-contracts c2f6e6fb12270f8b3c8703da7bf39bd01776e4b1. DO NOT EDIT.

package enums

// PullRequestState — Values of `state` in remediation.schema.json.
type PullRequestState string

const (
	PullRequestStateDraft  PullRequestState = "draft"
	PullRequestStateOpen   PullRequestState = "open"
	PullRequestStateMerged PullRequestState = "merged"
	PullRequestStateClosed PullRequestState = "closed"
)

// PullRequestStateValues returns all valid PullRequestState values.
func PullRequestStateValues() []PullRequestState {
	return []PullRequestState{
		PullRequestStateDraft,
		PullRequestStateOpen,
		PullRequestStateMerged,
		PullRequestStateClosed,
	}
}
