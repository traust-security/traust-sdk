// Code generated from traust-contracts 3ed68158121322ce06525de170aa383e057b6a7b. DO NOT EDIT.

package enums

// ValidationProbeClassification — Values of `classification` in validation.schema.json.
type ValidationProbeClassification string

const (
	ValidationProbeClassificationSafe        ValidationProbeClassification = "safe"
	ValidationProbeClassificationMutating    ValidationProbeClassification = "mutating"
	ValidationProbeClassificationDestructive ValidationProbeClassification = "destructive"
)

// ValidationProbeClassificationValues returns all valid ValidationProbeClassification values.
func ValidationProbeClassificationValues() []ValidationProbeClassification {
	return []ValidationProbeClassification{
		ValidationProbeClassificationSafe,
		ValidationProbeClassificationMutating,
		ValidationProbeClassificationDestructive,
	}
}
