// Code generated from traust-contracts c2f6e6fb12270f8b3c8703da7bf39bd01776e4b1. DO NOT EDIT.

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
