// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

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
