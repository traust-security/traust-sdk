// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

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
