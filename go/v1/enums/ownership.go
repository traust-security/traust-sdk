// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

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
