// Code generated from traust-contracts c2f6e6fb12270f8b3c8703da7bf39bd01776e4b1. DO NOT EDIT.

package enums

// Ownership — Who a corpus subject belongs to, which decides the denominator it counts toward. Previously declared only in the deployment's corpus-config and in the harness projection, never in the contract.
type Ownership string

const (
	OwnershipOwned      Ownership = "owned"
	OwnershipUpstream   Ownership = "upstream"
	OwnershipExternalBu Ownership = "external-bu"
	OwnershipHarnessQa  Ownership = "harness-qa"
)

// OwnershipValues returns all valid Ownership values.
func OwnershipValues() []Ownership {
	return []Ownership{
		OwnershipOwned,
		OwnershipUpstream,
		OwnershipExternalBu,
		OwnershipHarnessQa,
	}
}
