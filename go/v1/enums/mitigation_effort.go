// Code generated from traust-contracts 15bded2e3e4b6c70bd1e20ef8df50c7b66b5ebc1. DO NOT EDIT.

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
