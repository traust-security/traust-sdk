// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

package enums

// PackRequiredFact — Adopter facts a pack asset can require, by their traust-core artifact names; preflight checks each record exists.
type PackRequiredFact string

const (
	PackRequiredFactAdrRegistry           PackRequiredFact = "adr-registry"
	PackRequiredFactBenchmarkTarget       PackRequiredFact = "benchmark-target"
	PackRequiredFactOrgParameters         PackRequiredFact = "org-parameters"
	PackRequiredFactProductDefinitionsMap PackRequiredFact = "product-definitions-map"
	PackRequiredFactSlaPolicy             PackRequiredFact = "sla-policy"
)

// PackRequiredFactValues returns all valid PackRequiredFact values.
func PackRequiredFactValues() []PackRequiredFact {
	return []PackRequiredFact{
		PackRequiredFactAdrRegistry,
		PackRequiredFactBenchmarkTarget,
		PackRequiredFactOrgParameters,
		PackRequiredFactProductDefinitionsMap,
		PackRequiredFactSlaPolicy,
	}
}
