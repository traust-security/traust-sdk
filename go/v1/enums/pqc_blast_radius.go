// Code generated from traust-contracts 485578cbe584ad224d57b2c7c778174c0b66ba7c. DO NOT EDIT.

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
