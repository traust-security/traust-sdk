// Code generated from traust-contracts 15bded2e3e4b6c70bd1e20ef8df50c7b66b5ebc1. DO NOT EDIT.

package enums

// ValidationEvidenceType — Values of `type` in validation.schema.json.
type ValidationEvidenceType string

const (
	ValidationEvidenceTypeLog        ValidationEvidenceType = "log"
	ValidationEvidenceTypeManifest   ValidationEvidenceType = "manifest"
	ValidationEvidenceTypeHttp       ValidationEvidenceType = "http"
	ValidationEvidenceTypeDump       ValidationEvidenceType = "dump"
	ValidationEvidenceTypeScreenshot ValidationEvidenceType = "screenshot"
	ValidationEvidenceTypeDiff       ValidationEvidenceType = "diff"
	ValidationEvidenceTypeStdout     ValidationEvidenceType = "stdout"
	ValidationEvidenceTypeOther      ValidationEvidenceType = "other"
)

// ValidationEvidenceTypeValues returns all valid ValidationEvidenceType values.
func ValidationEvidenceTypeValues() []ValidationEvidenceType {
	return []ValidationEvidenceType{
		ValidationEvidenceTypeLog,
		ValidationEvidenceTypeManifest,
		ValidationEvidenceTypeHttp,
		ValidationEvidenceTypeDump,
		ValidationEvidenceTypeScreenshot,
		ValidationEvidenceTypeDiff,
		ValidationEvidenceTypeStdout,
		ValidationEvidenceTypeOther,
	}
}
