// Code generated from traust-contracts 369a2ab2c9b65225c6878eca75bada6a65df548f. DO NOT EDIT.

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
