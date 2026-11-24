// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

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
