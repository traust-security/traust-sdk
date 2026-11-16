// Code generated from traust-contracts 31d0d884672149e606934c2844ffcbf3bdd4b925. DO NOT EDIT.

package enums

// InterfaceRole — Recorded code-interface role.
type InterfaceRole string

const (
	InterfaceRoleOwner    InterfaceRole = "owner"
	InterfaceRoleConsumer InterfaceRole = "consumer"
	InterfaceRoleRbac     InterfaceRole = "rbac"
	InterfaceRoleWebhook  InterfaceRole = "webhook"
)

// InterfaceRoleValues returns all valid InterfaceRole values.
func InterfaceRoleValues() []InterfaceRole {
	return []InterfaceRole{
		InterfaceRoleOwner,
		InterfaceRoleConsumer,
		InterfaceRoleRbac,
		InterfaceRoleWebhook,
	}
}
