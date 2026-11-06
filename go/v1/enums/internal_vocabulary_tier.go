// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

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
