// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

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
