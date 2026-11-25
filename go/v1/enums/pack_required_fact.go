// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

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
