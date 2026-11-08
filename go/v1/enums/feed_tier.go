// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

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
