// Code generated from traust-contracts 58803ea798ad6431e10da337eb3a28349c1c9992. DO NOT EDIT.

package enums

// PqcBlastRadius — Values of `blast_radius` in pqc-readiness.schema.json.
type PqcBlastRadius string

const (
	PqcBlastRadiusFleet     PqcBlastRadius = "fleet"
	PqcBlastRadiusProduct   PqcBlastRadius = "product"
	PqcBlastRadiusService   PqcBlastRadius = "service"
	PqcBlastRadiusComponent PqcBlastRadius = "component"
)

// PqcBlastRadiusValues returns all valid PqcBlastRadius values.
func PqcBlastRadiusValues() []PqcBlastRadius {
	return []PqcBlastRadius{
		PqcBlastRadiusFleet,
		PqcBlastRadiusProduct,
		PqcBlastRadiusService,
		PqcBlastRadiusComponent,
	}
}
