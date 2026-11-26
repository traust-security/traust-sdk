// Code generated from traust-contracts 485578cbe584ad224d57b2c7c778174c0b66ba7c. DO NOT EDIT.

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
