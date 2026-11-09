package ops

import (
	"cmp"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"
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
	Remediation  *traust.RemediationRecordBody   `json:"remediation,omitempty"`
	AssessmentID uuid.UUID                       `json:"assessment_id,omitempty"`
	Input        json.RawMessage                 `json:"input,omitempty"`
}

// RecordingReceipt holds core's receipt for the one template that was recorded.
type RecordingReceipt struct {
	Assessment   *traust.AssessmentRecordResult `json:"assessment,omitempty"`
	Triage       *traust.TriageRecordResult     `json:"triage,omitempty"`
	Validation   *traust.ValidationRecordResult `json:"validation,omitempty"`
	Verification *traust.Verification           `json:"verification,omitempty"`
	Remediation  *traust.WorkItem               `json:"remediation,omitempty"`
}

type RecordingResult struct {
	Order   traust.RunOrder  `json:"order"`
	Receipt RecordingReceipt `json:"receipt"`
	Result  traust.JobResult `json:"result"`
}

// RecordingRun starts one explicitly dispatched order, holds its lease while the executor runs,
// records the executor's report — the remediation template's source_finding refs reconciled to
// the door's bare finding ids (values raw, the document re-marshaled), every other template's
// bytes original — and completes the order. No POST is retried.
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
	return run.runHeld(ctx, l, Execution{})
}

func (run RecordingRun) runHeld(ctx context.Context, l *lease, execution Execution) (*RecordingResult, error) {
	if err := run.Request.validate(); err != nil {
		return run.fail(ctx, l, traust.JobResult{}, err)
	}
	id, payload, provenance := run.Request.template()
	if err := l.stamp(&id, &provenance, &payload, ""); err != nil {
		return run.fail(ctx, l, traust.JobResult{}, err)
	}
	prepared, preparedFindings, err := run.prepare(ctx, l, execution)
	if err != nil {
		return run.preparationFailed(ctx, l, err)
	}
	produced, err := l.run(ctx, run.Executor, prepared)
	switch {
	case lost(err) || ctx.Err() != nil:
		return &RecordingResult{Order: *l.snapshot(), Result: produced.Result}, err
	case err != nil:
		produced.Result.Reason = ptr(traust.RunOrderReasonRunnerCrashed)
		return run.fail(ctx, l, produced.Result, err)
	}
	if err := checkOutcome(produced.Result); err != nil {
		produced.Result.Reason = ptr(traust.RunOrderReasonRunnerCrashed)
		return run.fail(ctx, l, produced.Result, err)
	}
	if produced.Result.Outcome == traust.OutcomeRefused || (produced.Result.Outcome == traust.OutcomeFailed && len(produced.Payload) == 0) {
		completed, err := l.complete(ctx, produced.Result)
		return &RecordingResult{Order: l.latest(completed), Result: produced.Result}, err
	}
	if !utf8.Valid(produced.Payload) {
		return run.fail(ctx, l, produced.Result, errors.New("ops: executor report is not valid UTF-8"))
	}
	if err := l.renew(ctx); err != nil {
		return &RecordingResult{Order: *l.snapshot(), Result: produced.Result}, err
	}
	receipt, err := superviseRequest(ctx, l, func(work context.Context) (RecordingReceipt, error) {
		return run.record(work, l, produced.Payload, preparedFindings)
	})
	if err != nil {
		if lost(err) || ctx.Err() != nil {
			return &RecordingResult{Order: *l.snapshot(), Result: produced.Result}, err
		}
		refusal := causeProblem(err)
		if definitive(refusal) {
			// Only a payload rejection classifies the output; order-specific errors close failed.
			if produced.Result.Reason == nil && payloadRejected(refusal.StatusCode) {
				produced.Result.Reason = ptr(traust.RunOrderReasonSessionOutputInvalid)
			}
			return run.fail(ctx, l, produced.Result, err)
		}
		// Host access/rate failures stop the runner; ambiguous writes remain for reconciliation.
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
	return run.Request.validate()
}

func (request RecordingRequest) validate() error {
	templates := 0
	for _, set := range []bool{
		request.Assessment != nil, request.Triage != nil, request.Validation != nil,
		request.Verification != nil, request.Remediation != nil,
	} {
		if set {
			templates++
		}
	}
	if templates != 1 {
		return errors.New("ops: recording run needs exactly one request template")
	}
	if request.Triage != nil && request.AssessmentID == uuid.Nil {
		return errors.New("ops: triage recording needs the source assessment id")
	}
	if _, payload, provenance := request.template(); payload != "" || provenance.Source == "" {
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

// prepare runs the run's domain preparation under the lease and returns the execution it
// shaped plus, for the remediation template, the prepared work item's finding ids — the ids
// core's record door validates source_finding refs against, kept for the record-time
// reconciliation of the qualified spelling contracts documents.
func (run RecordingRun) prepare(ctx context.Context, l *lease, execution Execution) (Execution, []uuid.UUID, error) {
	execution.Input = run.Request.Input
	if run.Request.Triage == nil && run.Request.Remediation == nil && run.Request.Verification == nil {
		return execution, nil, nil
	}
	if (run.Request.Remediation != nil || run.Request.Verification != nil) &&
		(l.origin.WorkItemId == nil || *l.origin.WorkItemId == uuid.Nil) {
		return execution, nil, &PreparationRefusedError{Reason: traust.RunOrderReasonBadRequest,
			Err: errors.New("ops: recording order needs a work_item_id")}
	}
	var preparedFindings []uuid.UUID
	shaped, err := superviseRequest(ctx, l, func(work context.Context) (Execution, error) {
		switch {
		case run.Request.Triage != nil:
			input, err := run.Client.PrepareTriage(work, run.Request.AssessmentID)
			if err != nil {
				return execution, fmt.Errorf("ops: prepare triage: %w", err)
			}
			execution.Input, execution.TriageInput = nil, input
		case run.Request.Remediation != nil:
			input, err := run.Client.PrepareRemediation(work, *l.origin.WorkItemId)
			if err != nil {
				return execution, fmt.Errorf("ops: prepare remediation: %w", err)
			}
			if input == nil || input.WorkItem.WorkItemId != *l.origin.WorkItemId ||
				(l.origin.RepositoryId != nil && input.Subject.RepositoryId != *l.origin.RepositoryId) {
				return execution, &PreparationRefusedError{Reason: traust.RunOrderReasonBadRequest,
					Err: errors.New("ops: remediation input does not match the bound work item and repository")}
			}
			if len(input.WorkItem.FindingIds) == 0 {
				return execution, &PreparationRefusedError{Reason: traust.RunOrderReasonBadRequest,
					Err: errors.New("ops: prepared work item names no findings to record against")}
			}
			preparedFindings = slices.Clone(input.WorkItem.FindingIds)
			execution.Input, err = json.Marshal(input)
			if err != nil {
				return execution, fmt.Errorf("ops: encode remediation input: %w", err)
			}
		case run.Request.Verification != nil:
			input, err := run.Client.PrepareVerification(work, *l.origin.WorkItemId)
			if err != nil {
				return execution, fmt.Errorf("ops: prepare verification: %w", err)
			}
			if err := run.verificationRef(work, l, input); err != nil {
				return execution, err
			}
			execution.Input, err = json.Marshal(input)
			if err != nil {
				return execution, fmt.Errorf("ops: encode verification input: %w", err)
			}
		}
		return execution, nil
	})
	return shaped, preparedFindings, err
}

func (run RecordingRun) verificationRef(ctx context.Context, l *lease, input *traust.VerificationInput) error {
	if input == nil || input.Subject.RepositoryId == uuid.Nil || input.WorkItem.WorkItemId != *l.origin.WorkItemId ||
		(l.origin.RepositoryId != nil && input.Subject.RepositoryId != *l.origin.RepositoryId) ||
		input.WorkItem.FixCommit == nil || *input.WorkItem.FixCommit == "" ||
		input.WorkItem.FixBranch == nil || *input.WorkItem.FixBranch == "" {
		return &PreparationRefusedError{Reason: traust.RunOrderReasonBadRequest,
			Err: errors.New("ops: verification input needs the bound work item's repository, fix commit and branch")}
	}
	workItem := &input.WorkItem
	repository, commit, branch := input.Subject.RepositoryId, *workItem.FixCommit, *workItem.FixBranch
	matches := func(ref *traust.RepositoryRef) bool {
		return ref != nil && ref.Id != nil && *ref.Id != uuid.Nil &&
			ref.RepositoryId == repository && ref.Sha == commit &&
			ref.Name != nil && *ref.Name == branch && ref.Kind != nil && *ref.Kind == string(traust.Branch)
	}
	refs, err := run.Client.ListRefs(ctx, repository, &traust.ListRefsParams{Sha: workItem.FixCommit})
	if err != nil {
		return fmt.Errorf("ops: look up verification fix ref: %w", err)
	}
	for i := range refs {
		if matches(&refs[i]) {
			return nil
		}
	}
	kind := traust.Branch
	registered, err := run.Client.RegisterRef(ctx, repository, traust.RegisterRefJSONRequestBody{
		Sha: commit, Name: workItem.FixBranch, Kind: &kind,
	})
	if err != nil {
		return fmt.Errorf("ops: register verification fix ref: %w", err)
	}
	if !matches(registered) {
		return &PreparationRefusedError{Reason: traust.RunOrderReasonBadRequest,
			Err: errors.New("ops: registered verification ref does not match the prepared fix context")}
	}
	return nil
}

func (run RecordingRun) preparationFailed(ctx context.Context, l *lease, cause error) (*RecordingResult, error) {
	if lost(cause) || ctx.Err() != nil {
		return &RecordingResult{Order: *l.snapshot()}, cause
	}
	if problem := causeProblem(cause); problem != nil && !definitive(problem) {
		return &RecordingResult{Order: *l.snapshot()}, cause
	}
	var refusal *PreparationRefusedError
	if !errors.As(cause, &refusal) {
		if definitive(causeProblem(cause)) {
			// A definitive preparation read (404 from prepare's own reads) completes
			// failed carrying the read's diagnosis, so the row stays attributable.
			return run.fail(ctx, l, traust.JobResult{Reason: reason(readFailedReason(cause))}, cause)
		}
		return &RecordingResult{Order: *l.snapshot()}, cause
	}
	if refusal.Reason == "" {
		return &RecordingResult{Order: *l.snapshot()}, errors.Join(cause,
			errors.New("ops: preparation refusal needs a reason"))
	}
	result := refused(refusal.Reason)
	completed, err := l.complete(ctx, result)
	return &RecordingResult{Order: l.latest(completed), Result: result}, errors.Join(cause, err)
}

// fail completes the order failed while it is still held, keeping the executor's result.
func (run RecordingRun) fail(ctx context.Context, l *lease, result traust.JobResult, cause error) (*RecordingResult, error) {
	result.Outcome = traust.OutcomeFailed
	completed, err := l.complete(ctx, result)
	return &RecordingResult{Order: l.latest(completed), Result: result}, errors.Join(cause, err)
}

func causeProblem(err error) *traust.ProblemError {
	var problem *traust.ProblemError
	errors.As(err, &problem)
	return problem
}

func definitive(problem *traust.ProblemError) bool {
	if problem == nil {
		return false
	}
	switch problem.StatusCode {
	case http.StatusUnauthorized, http.StatusForbidden, http.StatusTooManyRequests:
		return false
	}
	return problem.StatusCode >= 400 && problem.StatusCode < 500
}

func ptr[T any](value T) *T { return &value }

// payloadRejected reports whether a record-gate status means the payload itself is
// unacceptable — core maps every payload refusal to 422 (both ValidationError and FastAPI's
// RequestValidationError, problems.py STATUS), so no shape core produces answers 400 — rather
// than an access or rate refusal that says nothing about the output.
func payloadRejected(status int) bool {
	return status == http.StatusUnprocessableEntity
}
