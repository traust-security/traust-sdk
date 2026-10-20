// Code generated from traust-contracts ac517047c6566b019e63b911f6b7522000826b48. DO NOT EDIT.

package enums

// InterfaceKind — Values of `kind` in isolation-review.schema.json.
type InterfaceKind string

const (
	InterfaceKindApi       InterfaceKind = "api"
	InterfaceKindDataStore InterfaceKind = "data-store"
	InterfaceKindQueue     InterfaceKind = "queue"
	InterfaceKindIngress   InterfaceKind = "ingress"
	InterfaceKindWebhook   InterfaceKind = "webhook"
	InterfaceKindCli       InterfaceKind = "cli"
	InterfaceKindOther     InterfaceKind = "other"
)

// InterfaceKindValues returns all valid InterfaceKind values.
func InterfaceKindValues() []InterfaceKind {
	return []InterfaceKind{
		InterfaceKindApi,
		InterfaceKindDataStore,
		InterfaceKindQueue,
		InterfaceKindIngress,
		InterfaceKindWebhook,
		InterfaceKindCli,
		InterfaceKindOther,
	}
}
