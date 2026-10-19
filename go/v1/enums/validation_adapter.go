// Code generated from traust-contracts 3ed68158121322ce06525de170aa383e057b6a7b. DO NOT EDIT.

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
