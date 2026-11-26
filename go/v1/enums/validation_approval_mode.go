// Code generated from traust-contracts 485578cbe584ad224d57b2c7c778174c0b66ba7c. DO NOT EDIT.

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
