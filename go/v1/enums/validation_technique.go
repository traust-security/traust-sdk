// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

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
