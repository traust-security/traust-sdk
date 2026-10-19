// Code generated from traust-contracts 3ed68158121322ce06525de170aa383e057b6a7b. DO NOT EDIT.

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
