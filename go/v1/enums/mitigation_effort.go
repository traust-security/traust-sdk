// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

package enums

// MitigationEffort — Values of `effort` in threat-model.schema.json.
type MitigationEffort string

const (
	MitigationEffortXS      MitigationEffort = "XS"
	MitigationEffortS       MitigationEffort = "S"
	MitigationEffortM       MitigationEffort = "M"
	MitigationEffortL       MitigationEffort = "L"
	MitigationEffortXsLower MitigationEffort = "xs"
	MitigationEffortSLower  MitigationEffort = "s"
	MitigationEffortMLower  MitigationEffort = "m"
	MitigationEffortLLower  MitigationEffort = "l"
	MitigationEffortXlLower MitigationEffort = "xl"
)

// MitigationEffortValues returns all valid MitigationEffort values.
func MitigationEffortValues() []MitigationEffort {
	return []MitigationEffort{
		MitigationEffortXS,
		MitigationEffortS,
		MitigationEffortM,
		MitigationEffortL,
		MitigationEffortXsLower,
		MitigationEffortSLower,
		MitigationEffortMLower,
		MitigationEffortLLower,
		MitigationEffortXlLower,
	}
}
