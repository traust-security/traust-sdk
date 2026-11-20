// Code generated from traust-contracts 15bded2e3e4b6c70bd1e20ef8df50c7b66b5ebc1. DO NOT EDIT.

package enums

// IsolationDimension — Values of `isolation_dimensions` in cloud-config-audit.schema.json.
type IsolationDimension string

const (
	IsolationDimensionPrivilege      IsolationDimension = "privilege"
	IsolationDimensionEncryption     IsolationDimension = "encryption"
	IsolationDimensionAuthentication IsolationDimension = "authentication"
	IsolationDimensionConnectivity   IsolationDimension = "connectivity"
	IsolationDimensionHygiene        IsolationDimension = "hygiene"
)

// IsolationDimensionValues returns all valid IsolationDimension values.
func IsolationDimensionValues() []IsolationDimension {
	return []IsolationDimension{
		IsolationDimensionPrivilege,
		IsolationDimensionEncryption,
		IsolationDimensionAuthentication,
		IsolationDimensionConnectivity,
		IsolationDimensionHygiene,
	}
}
