package ops

import (
	"cmp"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"unicode/utf8"

	"github.com/google/uuid"
	"github.com/traust-security/traust-sdk/go/v1/traust"
)

// RecordingRequest selects exactly one generated request template. Payload must be empty: the
// executor supplies it. RunOrderId and provenance RunId/OccurredAt may be unset or match the
// server's order; source, artifact provenance, repository references and bindings are explicit.
type RecordingRequest struct {
	Assessment   *traust.AssessmentRecordRequest `json:"assessment,omitempty"`
	Triage       *traust.TriageRecordBody        `json:"triage,omitempty"`
	Validation   *traust.RecordValidation        `json:"validation,omitempty"`
	Verification *traust.VerificationRecordBody  `json:"verification,omitempty"`
	AssessmentID uuid.UUID                       `json:"assessment_id,omitempty"`
	Input        json.RawMessage                 `json:"input,omitempty"`
}

// RecordingReceipt holds core's receipt for the one template that was recorded.
type RecordingReceipt struct {
	Assessment   *traust.AssessmentRecordResult `json:"assessment,omitempty"`
	Triage       *traust.TriageRecordResult     `json:"triage,omitempty"`
	Validation   *traust.ValidationRecordResult `json:"validation,omitempty"`
	Verification *traust.Verification           `json:"verification,omitempty"`
}

type RecordingResult struct {
	Order   traust.RunOrder  `json:"order"`
	Receipt RecordingReceipt `json:"receipt"`
	Result  traust.JobResult `json:"result"`
}

// RecordingRun starts one explicitly dispatched order, holds its lease while the executor runs,
// records the executor's original bytes and completes the order. No POST is retried.
type RecordingRun struct {
	Client            *traust.Client
	Start             traust.StartRecordingRequest
	Request           RecordingRequest
	Executor          Executor
	HeartbeatInterval time.Duration
	RequestTimeout    time.Duration
}

func (run RecordingRun) Run(ctx context.Context) (*RecordingResult, error) {
	if err := run.validate(); err != nil {
		return nil, err
	}
	timeout := cmp.Or(run.RequestTimeout, defaultRequestTimeout)
	order, err := run.start(ctx, timeout)
	if err != nil {
		return nil, err
	}
	interval := cmp.Or(run.HeartbeatInterval, leaseDuration(run.Start.LeaseSeconds)/3)
	l, err := hold(run.Client, order, run.Start.LeaseSeconds, interval, timeout)
	if err != nil {
		return nil, err
	}
	execution := Execution{Input: run.Request.Input}
	if run.Request.Triage != nil {
		input, err := run.prepareTriage(ctx, l)
		if err != nil {
			return run.fail(ctx, l, traust.JobResult{}, err)
		}
		execution = Execution{TriageInput: input}
	}
	produced, err := l.run(ctx, run.Executor, execution)
	switch {
	case lost(err) || ctx.Err() != nil:
		return &RecordingResult{Order: *l.snapshot(), Result: produced.Result}, err
	case err != nil:
		return run.fail(ctx, l, produced.Result, err)
	case !utf8.Valid(produced.Payload):
		return run.fail(ctx, l, produced.Result, errors.New("ops: executor report is not valid UTF-8"))
	}
	if err := checkOutcome(produced.Result); err != nil {
		return run.fail(ctx, l, produced.Result, err)
	}
	if err := l.renew(ctx); err != nil {
		return &RecordingResult{Order: *l.snapshot(), Result: produced.Result}, err
	}
	receipt, err := run.record(ctx, l, produced.Payload)
	if err != nil {
		var refusal *traust.ProblemError
		if errors.As(err, &refusal) && refusal.StatusCode >= 400 && refusal.StatusCode < 500 {
			// The gate answered definitively: a terminally rejected artifact (never
			// retried, never mangled), completed failed while the lease is still held.
			// Only a payload refusal carries a reason — the order row alone says what
			// happened, the pod being gone by the time anyone looks; an access or rate
			// refusal (401/403/429) says nothing about the output, so no classifier is
			// invented for it.
			if produced.Result.Reason == nil && payloadRejected(refusal.StatusCode) {
				produced.Result.Reason = ptr(sessionOutputInvalid)
			}
			return run.fail(ctx, l, produced.Result, err)
		}
		// A transport or server error may follow a committed write: keep the order for
		// reconciliation and never complete it.
		return &RecordingResult{Order: *l.snapshot(), Receipt: receipt, Result: produced.Result}, err
	}
	completed, err := l.complete(ctx, produced.Result)
	return &RecordingResult{Order: l.latest(completed), Receipt: receipt, Result: produced.Result}, err
}

func (run RecordingRun) validate() error {
	if run.Client == nil || run.Executor == nil {
		return errors.New("ops: recording run needs a Client and an Executor")
	}
	lease := leaseDuration(run.Start.LeaseSeconds)
	if run.Start.LeaseSeconds <= 0 || run.RequestTimeout < 0 || run.HeartbeatInterval < 0 || run.HeartbeatInterval >= lease {
		return errors.New("ops: recording run needs a positive lease and a heartbeat shorter than it")
	}
	templates := 0
	for _, set := range []bool{run.Request.Assessment != nil, run.Request.Triage != nil, run.Request.Validation != nil, run.Request.Verification != nil} {
		if set {
			templates++
		}
	}
	if templates != 1 {
		return errors.New("ops: recording run needs exactly one request template")
	}
	if run.Request.Triage != nil && run.Request.AssessmentID == uuid.Nil {
		return errors.New("ops: triage recording needs the source assessment id")
	}
	if payload, provenance := run.Request.template(); payload != "" || provenance.Source == "" {
		return errors.New("ops: request template needs an empty payload and an explicit provenance source")
	}
	return nil
}

func (run RecordingRun) start(ctx context.Context, timeout time.Duration) (*traust.RunOrder, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	order, err := run.Client.Start(ctx, run.Start)
	if err != nil {
		return nil, fmt.Errorf("ops: start order: %w", err)
	}
	return order, nil
}

func (run RecordingRun) prepareTriage(ctx context.Context, l *lease) (*traust.TriageInput, error) {
	ctx, cancel, err := l.call(ctx)
	if err != nil {
		return nil, err
	}
	defer cancel()
	input, err := run.Client.PrepareTriage(ctx, run.Request.AssessmentID)
	if err != nil {
		return nil, fmt.Errorf("ops: prepare triage: %w", err)
	}
	return input, nil
}

// fail completes the order failed while it is still held, keeping the executor's result.
func (run RecordingRun) fail(ctx context.Context, l *lease, result traust.JobResult, cause error) (*RecordingResult, error) {
	result.Outcome = traust.OutcomeFailed
	completed, err := l.complete(ctx, result)
	return &RecordingResult{Order: l.latest(completed), Result: result}, errors.Join(cause, err)
}

func ptr[T any](value T) *T { return &value }

// sessionOutputInvalid is contracts' run-order-reason "session-output-invalid": the session's
// output failed the record gate's schema validation or lint. Core's reason vocabulary loads at
// runtime from contracts, so core takes the write today; this generated client predates the
// constant (client.gen.go carries 11 of the 12 contract values). contracts#109 merged the
// write-rejected definition (the gate-refusal spelling, distinct from this runner-caught
// validation spelling), but the switch needs core's own contracts pin to roll past 38f01a1
// first — then CORE_REF, the regenerated client, and the stamp change together.
const sessionOutputInvalid = traust.RunOrderReason("session-output-invalid")

// payloadRejected reports whether a record-gate status means the payload itself is
// unacceptable — core maps every payload refusal to 422 (both ValidationError and FastAPI's
// RequestValidationError, problems.py STATUS), so no shape core produces answers 400 — rather
// than an access or rate refusal that says nothing about the output.
func payloadRejected(status int) bool {
	return status == http.StatusUnprocessableEntity
}
