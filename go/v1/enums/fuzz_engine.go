// Code generated from traust-contracts 15bded2e3e4b6c70bd1e20ef8df50c7b66b5ebc1. DO NOT EDIT.

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
