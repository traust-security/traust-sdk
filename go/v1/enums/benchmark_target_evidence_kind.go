// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

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
