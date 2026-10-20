// Code generated from traust-contracts ac517047c6566b019e63b911f6b7522000826b48. DO NOT EDIT.

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
