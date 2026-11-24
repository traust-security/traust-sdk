// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

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
