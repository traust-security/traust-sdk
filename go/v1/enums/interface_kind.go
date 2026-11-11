// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

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
