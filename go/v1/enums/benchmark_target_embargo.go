// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

package enums

// BenchmarkTargetEmbargo — Values of `embargo` in benchmark-target.schema.json.
type BenchmarkTargetEmbargo string

const (
	BenchmarkTargetEmbargoInternal BenchmarkTargetEmbargo = "internal"
	BenchmarkTargetEmbargoPublic   BenchmarkTargetEmbargo = "public"
)

// BenchmarkTargetEmbargoValues returns all valid BenchmarkTargetEmbargo values.
func BenchmarkTargetEmbargoValues() []BenchmarkTargetEmbargo {
	return []BenchmarkTargetEmbargo{
		BenchmarkTargetEmbargoInternal,
		BenchmarkTargetEmbargoPublic,
	}
}
