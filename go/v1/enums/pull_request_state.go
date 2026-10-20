// Code generated from traust-contracts ac517047c6566b019e63b911f6b7522000826b48. DO NOT EDIT.

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
