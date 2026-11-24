// Code generated from traust-contracts c2f6e6fb12270f8b3c8703da7bf39bd01776e4b1. DO NOT EDIT.

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
