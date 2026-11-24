// Code generated from traust-contracts c2f6e6fb12270f8b3c8703da7bf39bd01776e4b1. DO NOT EDIT.

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
