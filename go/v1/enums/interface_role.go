// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

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
