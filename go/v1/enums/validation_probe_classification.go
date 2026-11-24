// Code generated from traust-contracts 58803ea798ad6431e10da337eb3a28349c1c9992. DO NOT EDIT.

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
