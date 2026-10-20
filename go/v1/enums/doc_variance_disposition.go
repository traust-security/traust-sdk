// Code generated from traust-contracts ac517047c6566b019e63b911f6b7522000826b48. DO NOT EDIT.

package enums

// DocVarianceDisposition — Values of `disposition` in doc-variance.schema.json.
type DocVarianceDisposition string

const (
	DocVarianceDispositionOpen         DocVarianceDisposition = "open"
	DocVarianceDispositionDocCorrected DocVarianceDisposition = "doc_corrected"
	DocVarianceDispositionCodeFixed    DocVarianceDisposition = "code_fixed"
	DocVarianceDispositionAccepted     DocVarianceDisposition = "accepted"
	DocVarianceDispositionSuperseded   DocVarianceDisposition = "superseded"
)

// DocVarianceDispositionValues returns all valid DocVarianceDisposition values.
func DocVarianceDispositionValues() []DocVarianceDisposition {
	return []DocVarianceDisposition{
		DocVarianceDispositionOpen,
		DocVarianceDispositionDocCorrected,
		DocVarianceDispositionCodeFixed,
		DocVarianceDispositionAccepted,
		DocVarianceDispositionSuperseded,
	}
}
