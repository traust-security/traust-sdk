// Code generated from traust-contracts 369a2ab2c9b65225c6878eca75bada6a65df548f. DO NOT EDIT.

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
