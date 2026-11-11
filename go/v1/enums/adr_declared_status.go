// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

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
