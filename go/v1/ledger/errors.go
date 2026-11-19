package ledger

import (
	"errors"
	"fmt"

	"github.com/traust-security/traust-sdk/go/v1/ledger/internal/transport"
)

// ErrIncompleteLayerDocument is returned when a raw layer document does
// not carry the structure every layer holds (a metadata object, events and
// needs_review arrays) — a truncated or foreign body, never a layer.
var ErrIncompleteLayerDocument = errors.New("ledger layer document is incomplete")

// ErrResponseTooLarge is returned when a response body exceeds the limit
// the client was constructed with (WithResponseLimit): the read stops at
// the limit, the body is never buffered whole.
var ErrResponseTooLarge = transport.ErrResponseTooLarge

// Phase identifies where in the submission pipeline an error occurred.
type Phase string

const (
	PhaseMarshal     Phase = "marshal"
	PhaseValidate    Phase = "validate"
	PhaseFingerprint Phase = "fingerprint"
	PhaseConvert     Phase = "convert"
	PhaseIngest      Phase = "ingest"
	PhaseDecode      Phase = "decode"
)

// IngestError wraps pipeline errors with the operation kind and phase for
// structured diagnostics.
type IngestError struct {
	Kind  IngestKind
	Phase Phase
	Err   error
}

func (e *IngestError) Error() string {
	return fmt.Sprintf("ingest[%s] %s: %v", e.Kind, e.Phase, e.Err)
}

func (e *IngestError) Unwrap() error { return e.Err }

func ingestErr(kind IngestKind, phase Phase, err error) *IngestError {
	return &IngestError{Kind: kind, Phase: phase, Err: err}
}
