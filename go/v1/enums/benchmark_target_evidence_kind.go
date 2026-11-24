// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

package enums

// BenchmarkTargetEvidenceKind — Values of `kind` in benchmark-target.schema.json.
type BenchmarkTargetEvidenceKind string

const (
	BenchmarkTargetEvidenceKindVerificationReport BenchmarkTargetEvidenceKind = "verification_report"
	BenchmarkTargetEvidenceKindLedgerEvent        BenchmarkTargetEvidenceKind = "ledger_event"
	BenchmarkTargetEvidenceKindAdvisory           BenchmarkTargetEvidenceKind = "advisory"
	BenchmarkTargetEvidenceKindFuzzCrash          BenchmarkTargetEvidenceKind = "fuzz_crash"
	BenchmarkTargetEvidenceKindSeedRecipe         BenchmarkTargetEvidenceKind = "seed_recipe"
)

// BenchmarkTargetEvidenceKindValues returns all valid BenchmarkTargetEvidenceKind values.
func BenchmarkTargetEvidenceKindValues() []BenchmarkTargetEvidenceKind {
	return []BenchmarkTargetEvidenceKind{
		BenchmarkTargetEvidenceKindVerificationReport,
		BenchmarkTargetEvidenceKindLedgerEvent,
		BenchmarkTargetEvidenceKindAdvisory,
		BenchmarkTargetEvidenceKindFuzzCrash,
		BenchmarkTargetEvidenceKindSeedRecipe,
	}
}
