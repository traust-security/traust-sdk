// Code generated from traust-contracts 8f684d27d978e53c0ffc2b43af379fc1284c26eb. DO NOT EDIT.

package enums

// ValidationTechnique — Values of `technique` in validation.schema.json.
type ValidationTechnique string

const (
	ValidationTechniqueReplay  ValidationTechnique = "replay"
	ValidationTechniqueAdapted ValidationTechnique = "adapted"
	ValidationTechniqueChained ValidationTechnique = "chained"
	ValidationTechniqueNovel   ValidationTechnique = "novel"
	ValidationTechniqueRecon   ValidationTechnique = "recon"
	ValidationTechniqueSkip    ValidationTechnique = "skip"
)

// ValidationTechniqueValues returns all valid ValidationTechnique values.
func ValidationTechniqueValues() []ValidationTechnique {
	return []ValidationTechnique{
		ValidationTechniqueReplay,
		ValidationTechniqueAdapted,
		ValidationTechniqueChained,
		ValidationTechniqueNovel,
		ValidationTechniqueRecon,
		ValidationTechniqueSkip,
	}
}
