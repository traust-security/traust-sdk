// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

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
