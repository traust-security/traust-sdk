// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

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
