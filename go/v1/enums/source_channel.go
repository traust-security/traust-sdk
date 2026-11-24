// Code generated from traust-contracts 58803ea798ad6431e10da337eb3a28349c1c9992. DO NOT EDIT.

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
