// Code generated from traust-contracts 3ed68158121322ce06525de170aa383e057b6a7b. DO NOT EDIT.

package enums

// AdrDeclaredStatus — register-level status declaration for registers whose format carries no status field — applies to every decision, overrides parsing, and MUST cite its declarer/date in `note` (only non-terminal statuses declarable; superseded/deprecated always come from the records themselves)
type AdrDeclaredStatus string

const (
	AdrDeclaredStatusAccepted AdrDeclaredStatus = "accepted"
	AdrDeclaredStatusProposed AdrDeclaredStatus = "proposed"
)

// AdrDeclaredStatusValues returns all valid AdrDeclaredStatus values.
func AdrDeclaredStatusValues() []AdrDeclaredStatus {
	return []AdrDeclaredStatus{
		AdrDeclaredStatusAccepted,
		AdrDeclaredStatusProposed,
	}
}
