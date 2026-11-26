// Code generated from traust-contracts 485578cbe584ad224d57b2c7c778174c0b66ba7c. DO NOT EDIT.

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
