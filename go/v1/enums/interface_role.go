// Code generated from traust-contracts 369a2ab2c9b65225c6878eca75bada6a65df548f. DO NOT EDIT.

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
