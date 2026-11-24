// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

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
