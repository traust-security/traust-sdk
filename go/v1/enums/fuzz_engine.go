// Code generated from traust-contracts c2f6e6fb12270f8b3c8703da7bf39bd01776e4b1. DO NOT EDIT.

package enums

// FuzzEngine — The fuzzing engine a harness ran under.
type FuzzEngine string

const (
	FuzzEngineGoNative    FuzzEngine = "go_native"
	FuzzEngineLibfuzzer   FuzzEngine = "libfuzzer"
	FuzzEngineAflplusplus FuzzEngine = "aflplusplus"
	FuzzEngineHonggfuzz   FuzzEngine = "honggfuzz"
	FuzzEngineAtheris     FuzzEngine = "atheris"
	FuzzEngineJazzer      FuzzEngine = "jazzer"
	FuzzEngineJazzerJs    FuzzEngine = "jazzer_js"
	FuzzEngineCargoFuzz   FuzzEngine = "cargo_fuzz"
)

// FuzzEngineValues returns all valid FuzzEngine values.
func FuzzEngineValues() []FuzzEngine {
	return []FuzzEngine{
		FuzzEngineGoNative,
		FuzzEngineLibfuzzer,
		FuzzEngineAflplusplus,
		FuzzEngineHonggfuzz,
		FuzzEngineAtheris,
		FuzzEngineJazzer,
		FuzzEngineJazzerJs,
		FuzzEngineCargoFuzz,
	}
}
