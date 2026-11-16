// Code generated from traust-contracts 31d0d884672149e606934c2844ffcbf3bdd4b925. DO NOT EDIT.

package enums

// ValidationApprovalMode — Values of `mode` in validation.schema.json.
type ValidationApprovalMode string

const (
	ValidationApprovalModeInteractive ValidationApprovalMode = "interactive"
	ValidationApprovalModeAuto        ValidationApprovalMode = "auto"
	ValidationApprovalModeDryRun      ValidationApprovalMode = "dry-run"
)

// ValidationApprovalModeValues returns all valid ValidationApprovalMode values.
func ValidationApprovalModeValues() []ValidationApprovalMode {
	return []ValidationApprovalMode{
		ValidationApprovalModeInteractive,
		ValidationApprovalModeAuto,
		ValidationApprovalModeDryRun,
	}
}
