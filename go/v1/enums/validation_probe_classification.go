// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

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
