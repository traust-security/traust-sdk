// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

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
