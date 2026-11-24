// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

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
