// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

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
