// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

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
