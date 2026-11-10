// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

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
