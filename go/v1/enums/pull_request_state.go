// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

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
