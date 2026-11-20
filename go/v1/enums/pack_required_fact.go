// Code generated from traust-contracts 15bded2e3e4b6c70bd1e20ef8df50c7b66b5ebc1. DO NOT EDIT.

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
