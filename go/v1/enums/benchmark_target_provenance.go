// Code generated from traust-contracts c2f6e6fb12270f8b3c8703da7bf39bd01776e4b1. DO NOT EDIT.

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
