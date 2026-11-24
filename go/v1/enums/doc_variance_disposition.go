// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

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
