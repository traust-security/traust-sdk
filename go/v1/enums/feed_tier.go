// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

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
