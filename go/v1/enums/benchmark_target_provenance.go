// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

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
