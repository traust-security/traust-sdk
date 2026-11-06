// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

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
