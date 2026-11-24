// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

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
