// Code generated from traust-contracts 8f684d27d978e53c0ffc2b43af379fc1284c26eb. DO NOT EDIT.

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
