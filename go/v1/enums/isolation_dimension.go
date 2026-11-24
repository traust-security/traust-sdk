// Code generated from traust-contracts 369a2ab2c9b65225c6878eca75bada6a65df548f. DO NOT EDIT.

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
