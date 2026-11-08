// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

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
