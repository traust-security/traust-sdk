// Code generated from traust-contracts ac517047c6566b019e63b911f6b7522000826b48. DO NOT EDIT.

package enums

// BenchmarkTargetProvenance — self_replay = campaign ledger resolved-with-fix-commit; cve_replay = public advisory (training-contamination caveat applies); seeded = injected pattern
type BenchmarkTargetProvenance string

const (
	BenchmarkTargetProvenanceSelfReplay BenchmarkTargetProvenance = "self_replay"
	BenchmarkTargetProvenanceCveReplay  BenchmarkTargetProvenance = "cve_replay"
	BenchmarkTargetProvenanceSeeded     BenchmarkTargetProvenance = "seeded"
)

// BenchmarkTargetProvenanceValues returns all valid BenchmarkTargetProvenance values.
func BenchmarkTargetProvenanceValues() []BenchmarkTargetProvenance {
	return []BenchmarkTargetProvenance{
		BenchmarkTargetProvenanceSelfReplay,
		BenchmarkTargetProvenanceCveReplay,
		BenchmarkTargetProvenanceSeeded,
	}
}
