// Code generated from traust-contracts 2e521561f022fe839dd828ddf2185ba3a0a1045f. DO NOT EDIT.

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
