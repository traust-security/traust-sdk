// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

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
