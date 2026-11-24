// Code generated from traust-contracts 58803ea798ad6431e10da337eb3a28349c1c9992. DO NOT EDIT.

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
