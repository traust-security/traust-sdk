// Code generated from traust-contracts 369a2ab2c9b65225c6878eca75bada6a65df548f. DO NOT EDIT.

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
