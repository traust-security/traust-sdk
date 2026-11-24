// Code generated from traust-contracts c2f6e6fb12270f8b3c8703da7bf39bd01776e4b1. DO NOT EDIT.

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
