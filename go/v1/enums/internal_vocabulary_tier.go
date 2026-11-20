// Code generated from traust-contracts 15bded2e3e4b6c70bd1e20ef8df50c7b66b5ebc1. DO NOT EDIT.

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
