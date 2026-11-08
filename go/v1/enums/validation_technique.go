// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

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
