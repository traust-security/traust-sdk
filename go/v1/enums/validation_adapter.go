// Code generated from traust-contracts 15bded2e3e4b6c70bd1e20ef8df50c7b66b5ebc1. DO NOT EDIT.

package enums

// ValidationAdapter — Values of `adapter` in validation.schema.json.
type ValidationAdapter string

const (
	ValidationAdapterK8s       ValidationAdapter = "k8s"
	ValidationAdapterContainer ValidationAdapter = "container"
	ValidationAdapterWasm      ValidationAdapter = "wasm"
)

// ValidationAdapterValues returns all valid ValidationAdapter values.
func ValidationAdapterValues() []ValidationAdapter {
	return []ValidationAdapter{
		ValidationAdapterK8s,
		ValidationAdapterContainer,
		ValidationAdapterWasm,
	}
}
