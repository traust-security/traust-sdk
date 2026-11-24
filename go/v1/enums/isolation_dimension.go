// Code generated from traust-contracts 58803ea798ad6431e10da337eb3a28349c1c9992. DO NOT EDIT.

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
