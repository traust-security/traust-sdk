// Code generated from traust-contracts c2f6e6fb12270f8b3c8703da7bf39bd01776e4b1. DO NOT EDIT.

package enums

// SourceChannel — Decision.source_channel: the channel carrying a human judgement, not the producing assessment kind.
type SourceChannel string

const (
	SourceChannelMrComment   SourceChannel = "mr_comment"
	SourceChannelCommit      SourceChannel = "commit"
	SourceChannelJira        SourceChannel = "jira"
	SourceChannelInteractive SourceChannel = "interactive"
)

// SourceChannelValues returns all valid SourceChannel values.
func SourceChannelValues() []SourceChannel {
	return []SourceChannel{
		SourceChannelMrComment,
		SourceChannelCommit,
		SourceChannelJira,
		SourceChannelInteractive,
	}
}
