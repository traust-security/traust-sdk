// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

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
