// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

package enums

// InternalVocabularyTier — Values of `tier` in internal-vocabulary.schema.json.
type InternalVocabularyTier string

const (
	InternalVocabularyTierBlock  InternalVocabularyTier = "block"
	InternalVocabularyTierReview InternalVocabularyTier = "review"
	InternalVocabularyTierNote   InternalVocabularyTier = "note"
)

// InternalVocabularyTierValues returns all valid InternalVocabularyTier values.
func InternalVocabularyTierValues() []InternalVocabularyTier {
	return []InternalVocabularyTier{
		InternalVocabularyTierBlock,
		InternalVocabularyTierReview,
		InternalVocabularyTierNote,
	}
}
