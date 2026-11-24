// Code generated from traust-contracts 58803ea798ad6431e10da337eb3a28349c1c9992. DO NOT EDIT.

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
