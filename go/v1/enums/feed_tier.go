// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

package enums

// FeedTier — Values of `tier` in feeds.schema.json.
type FeedTier string

const (
	FeedTierCached FeedTier = "cached"
	FeedTierLive   FeedTier = "live"
)

// FeedTierValues returns all valid FeedTier values.
func FeedTierValues() []FeedTier {
	return []FeedTier{
		FeedTierCached,
		FeedTierLive,
	}
}
