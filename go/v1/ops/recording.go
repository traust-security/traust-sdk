// Package ops owns order leases and native recording lifecycles.
// Core selects and authorizes work; consumers supply an Executor.
package ops

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"
	"unicode/utf8"

	"github.com/google/uuid"
	"github.com/traust-security/traust-sdk/go/v1/api"
)

// Executor performs the dispatched work. On context cancellation it MUST stop
// actual work (including subprocess descendants), wait for it to stop, and return.
// Payload is the original produced report, never a report reconstructed from DTOs.
type Executor interface {
	Execute(context.Context, Execution) (ProducedReport, error)
}

// Execution contains only native server context, never authentication material.
type Execution struct {
	Order       api.RunOrder     `json:"order"`
	TriageInput *api.TriageInput `json:"triage_input,omitempty"`
	Input       json.RawMessage  `json:"input,omitempty"`
	Argv        []string         `json:"argv,omitempty"`
}

type ProducedReport struct {
	Payload []byte
	Result  api.JobResult
}

// RecordingRequest selects exactly one generated request template. Payload must
// be empty; the executor supplies it. RunOrderId and provenance RunId/OccurredAt
// may be unset or match the server's immutable order ID/start time. Source,
// artifact provenance, repository references and bindings are explicit inputs.
type RecordingRequest struct {
	Assessment   *api.AssessmentRecordRequest `json:"assessment,omitempty"`
	Triage       *api.TriageRecordBody        `json:"triage,omitempty"`
	Validation   *api.RecordValidation        `json:"validation,omitempty"`
	Verification *api.VerificationRecordBody  `json:"verification,omitempty"`
	AssessmentID uuid.UUID                    `json:"assessment_id,omitempty"`
	Input        json.RawMessage              `json:"input,omitempty"`
}

// RecordingReceipt wraps native receipts without translating their IDs or data.
type RecordingReceipt struct {
	Assessment   *api.AssessmentRecordResult `json:"assessment,omitempty"`
	Triage       *api.TriageRecordResult     `json:"triage,omitempty"`
	Validation   *api.ValidationRecordResult `json:"validation,omitempty"`
	Verification *api.Verification           `json:"verification,omitempty"`
}

type RecordingResult struct {
	Order   api.RunOrder     `json:"order"`
	Receipt RecordingReceipt `json:"receipt"`
	Result  api.JobResult    `json:"result"`
}

// RecordingRun starts one explicitly dispatched order, maintains its lease while
// executing, records original bytes, and completes the owned order. There are no
// POST retries. The caller supplies a BoundClient (normally with FileTokenSource).
type RecordingRun struct {
	Client            *api.BoundClient
	Start             api.StartRecordingRequest
	Request           RecordingRequest
	Executor          Executor
	HeartbeatInterval time.Duration
	RequestTimeout    time.Duration
}

// LeaseLostError means no recording/completion may follow. Its cause retains
// typed SDK problem details without requiring consumers to infer ownership policy.
type LeaseLostError struct{ Err error }

func (e *LeaseLostError) Error() string { return "recording run lost its lease" }
func (e *LeaseLostError) Unwrap() error { return e.Err }

func (run RecordingRun) Run(ctx context.Context) (*RecordingResult, error) {
	if err := run.validate(); err != nil {
		return nil, err
	}
	if run.RequestTimeout == 0 {
		run.RequestTimeout = 60 * time.Second
	}
	if run.HeartbeatInterval == 0 {
		run.HeartbeatInterval = time.Duration(run.Start.LeaseSeconds) * time.Second / 3
	}
	requestCtx, cancel := context.WithTimeout(ctx, run.RequestTimeout)
	response, err := run.Client.StartWithResponse(requestCtx, run.Start)
	cancel()
	if err != nil {
		return nil, err
	}
	order, err := nativeResponse(response.StatusCode(), response.Body, response.JSON200)
	if err != nil {
		return nil, err
	}
	if err := usableOrder(order); err != nil {
		return nil, &LeaseLostError{Err: err}
	}
	// Retain this original server snapshot. Heartbeats must never replace the
	// provenance origin with their own timestamps or any consumer clock value.
	original := *order
	var input *api.TriageInput
	if run.Request.Triage != nil {
		requestCtx, cancel = run.ownedContext(ctx, order)
		prepared, prepareErr := run.Client.PrepareTriageWithResponse(requestCtx, run.Request.AssessmentID)
		cancel()
		if prepareErr == nil {
			input, prepareErr = nativeResponse(prepared.StatusCode(), prepared.Body, prepared.JSON200)
		}
		if prepareErr != nil {
			return run.finishFailure(ctx, original, order, api.JobResult{}, prepareErr)
		}
	}
	execution := Execution{Order: original, TriageInput: input}
	if run.Request.Triage == nil {
		execution.Input = run.Request.Input
	}
	produced, executeErr := run.execute(ctx, original, order, execution)
	if executeErr != nil {
		var lost *LeaseLostError
		if errors.As(executeErr, &lost) || ctx.Err() != nil {
			return &RecordingResult{Order: *order, Result: produced.Result}, executeErr
		}
		return run.finishFailure(ctx, original, order, produced.Result, executeErr)
	}
	if !utf8.Valid(produced.Payload) {
		return run.finishFailure(ctx, original, order, produced.Result, errors.New("executor report is not valid UTF-8"))
	}
	if produced.Result.Outcome == "" {
		return run.finishFailure(ctx, original, order, produced.Result, errors.New("executor did not supply a native outcome"))
	}
	// Serialize lease checks and all final mutations. A heartbeat failure is
	// terminal even if it might have been transient; no ambiguous POST is retried.
	order, err = run.heartbeat(ctx, original, order)
	if err != nil {
		return &RecordingResult{Order: original, Result: produced.Result}, err
	}
	receipt, err := run.record(ctx, original, order, produced.Payload)
	if err != nil {
		var refusal *api.ProblemError
		if errors.As(err, &refusal) && refusal.StatusCode >= 400 && refusal.StatusCode < 500 {
			return run.finishFailure(ctx, original, order, produced.Result, err)
		}
		// A transport, decoding or server error may follow a committed write.
		// Preserve the run identity for reconciliation; never complete it failed.
		return &RecordingResult{Order: *order, Receipt: receipt, Result: produced.Result}, err
	}
	order, err = run.heartbeat(ctx, original, order)
	if err != nil {
		return &RecordingResult{Order: original, Receipt: receipt, Result: produced.Result}, err
	}
	completed, err := run.complete(ctx, order, produced.Result)
	return &RecordingResult{Order: valueOrOriginal(completed, original), Receipt: receipt, Result: produced.Result}, err
}

func (run RecordingRun) validate() error {
	if run.Client == nil || run.Executor == nil {
		return errors.New("recording run requires a bound client and executor")
	}
	if run.Start.LeaseSeconds <= 0 || run.RequestTimeout < 0 || run.HeartbeatInterval < 0 || (run.HeartbeatInterval != 0 && run.HeartbeatInterval >= time.Duration(run.Start.LeaseSeconds)*time.Second) {
		return errors.New("recording run requires positive bounded lease/request/heartbeat durations")
	}
	count := 0
	for _, configured := range []bool{run.Request.Assessment != nil, run.Request.Triage != nil, run.Request.Validation != nil, run.Request.Verification != nil} {
		if configured {
			count++
		}
	}
	if count != 1 {
		return errors.New("recording run requires exactly one native request template")
	}
	if run.Request.Triage != nil && run.Request.AssessmentID == uuid.Nil {
		return errors.New("triage recording requires the native source assessment UUID")
	}
	payload, provenance := run.Request.templateInput()
	if payload != "" || provenance.Source == "" {
		return errors.New("recording run requires an empty payload template and explicit native provenance source")
	}
	return nil
}

func usableOrder(order *api.RunOrder) error {
	if order == nil || order.State != api.RunOrderStateLeased || order.Id == nil || *order.Id == uuid.Nil || order.StartedAt == nil || order.StartedAt.IsZero() || order.LeaseExpiresAt == nil || !order.LeaseExpiresAt.After(time.Now()) {
		return errors.New("Core did not return a concrete live order lease and provenance origin")
	}
	return nil
}

func (run RecordingRun) ownedContext(ctx context.Context, order *api.RunOrder) (context.Context, context.CancelFunc) {
	deadline := time.Now().Add(run.RequestTimeout)
	if order.LeaseExpiresAt.Before(deadline) {
		deadline = *order.LeaseExpiresAt
	}
	return context.WithDeadline(ctx, deadline)
}

func (run RecordingRun) heartbeat(ctx context.Context, original api.RunOrder, current *api.RunOrder) (*api.RunOrder, error) {
	if err := usableOrder(current); err != nil {
		return nil, &LeaseLostError{Err: err}
	}
	if err := ctx.Err(); err != nil {
		return nil, &LeaseLostError{Err: err}
	}
	requestCtx, cancel := run.ownedContext(ctx, current)
	defer cancel()
	response, err := run.Client.HeartbeatWithResponse(requestCtx, *original.Id, api.LeaseRequest{LeaseSeconds: run.Start.LeaseSeconds})
	if err != nil {
		return nil, &LeaseLostError{Err: err}
	}
	order, err := nativeResponse(response.StatusCode(), response.Body, response.JSON200)
	if err == nil {
		err = usableOrder(order)
	}
	if err == nil && (*order.Id != *original.Id || !order.StartedAt.Equal(*original.StartedAt)) {
		err = errors.New("heartbeat changed immutable server run provenance")
	}
	if err != nil {
		return nil, &LeaseLostError{Err: err}
	}
	return order, nil
}

func (run RecordingRun) execute(ctx context.Context, original api.RunOrder, current *api.RunOrder, execution Execution) (ProducedReport, error) {
	if err := usableOrder(current); err != nil {
		return ProducedReport{}, &LeaseLostError{Err: err}
	}
	if err := ctx.Err(); err != nil {
		return ProducedReport{}, &LeaseLostError{Err: err}
	}
	executorCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	type executionResult struct {
		produced ProducedReport
		err      error
	}
	done := make(chan executionResult, 1)
	executionOrder := original
	// Native generated models use pointers. Do not share the SDK's authority
	// and lease/provenance origin storage with an executor implementation.
	id, started, expiresAt := *original.Id, *original.StartedAt, *original.LeaseExpiresAt
	executionOrder.Id, executionOrder.StartedAt, executionOrder.LeaseExpiresAt = &id, &started, &expiresAt
	execution.Order = executionOrder
	go func() {
		produced, err := run.Executor.Execute(executorCtx, execution)
		done <- executionResult{produced: produced, err: err}
	}()
	ticker := time.NewTicker(run.HeartbeatInterval)
	defer ticker.Stop()
	expires := time.NewTimer(time.Until(*current.LeaseExpiresAt))
	defer expires.Stop()
	stop := func(cause error) (ProducedReport, error) {
		cancel()
		<-done // Executor contract: actual work has stopped before it returns.
		return ProducedReport{}, cause
	}
	for {
		select {
		case <-ctx.Done():
			return stop(&LeaseLostError{Err: ctx.Err()})
		case <-expires.C:
			return stop(&LeaseLostError{Err: errors.New("order lease expired")})
		case <-ticker.C:
			updated, err := run.heartbeat(ctx, original, current)
			if err != nil {
				return stop(err)
			}
			*current = *updated
			expires.Reset(time.Until(*current.LeaseExpiresAt))
		case result := <-done:
			if ctx.Err() != nil || !current.LeaseExpiresAt.After(time.Now()) {
				return ProducedReport{}, &LeaseLostError{Err: errors.New("execution ended without a live lease")}
			}
			return result.produced, result.err
		}
	}
}

func (run RecordingRun) finishFailure(ctx context.Context, original api.RunOrder, current *api.RunOrder, result api.JobResult, cause error) (*RecordingResult, error) {
	owned, err := run.heartbeat(ctx, original, current)
	if err != nil {
		return &RecordingResult{Order: valueOrOriginal(current, original), Result: result}, errors.Join(cause, err)
	}
	result.Outcome = api.OutcomeFailed
	completed, completeErr := run.complete(ctx, owned, result)
	return &RecordingResult{Order: valueOrOriginal(completed, original), Result: result}, errors.Join(cause, completeErr)
}

func (run RecordingRun) complete(ctx context.Context, order *api.RunOrder, result api.JobResult) (*api.RunOrder, error) {
	if err := usableOrder(order); err != nil {
		return nil, &LeaseLostError{Err: err}
	}
	if err := ctx.Err(); err != nil {
		return nil, &LeaseLostError{Err: err}
	}
	requestCtx, cancel := run.ownedContext(ctx, order)
	defer cancel()
	response, err := run.Client.CompleteWithResponse(requestCtx, *order.Id, api.CompleteRequest{Result: result})
	if err != nil {
		return nil, err
	}
	return nativeResponse(response.StatusCode(), response.Body, response.JSON200)
}

func valueOrOriginal(order *api.RunOrder, original api.RunOrder) api.RunOrder {
	if order != nil {
		return *order
	}
	return original
}

func nativeResponse[T any](status int, raw []byte, receipt *T) (*T, error) {
	if status < 200 || status >= 300 {
		failure := &api.ProblemError{StatusCode: status, Body: string(raw)}
		var problem api.Problem
		if json.Unmarshal(raw, &problem) == nil && (problem.Title != "" || problem.Type != "" || problem.Status != 0) {
			failure.Problem = &problem
		}
		return nil, failure
	}
	if status != 200 || receipt == nil || bytes.Equal(bytes.TrimSpace(raw), []byte("null")) {
		return nil, fmt.Errorf("Core returned HTTP %d without a concrete native JSON200 response", status)
	}
	return receipt, nil
}

func bindProvenance(order api.RunOrder, requestOrder uuid.UUID, provenance api.Provenance) (api.Provenance, error) {
	if requestOrder != uuid.Nil && requestOrder != *order.Id {
		return api.Provenance{}, errors.New("request template run_order_id conflicts with the server order")
	}
	if provenance.RunId != uuid.Nil && provenance.RunId != *order.Id {
		return api.Provenance{}, errors.New("request provenance run_id conflicts with the server order")
	}
	if !provenance.OccurredAt.IsZero() && !provenance.OccurredAt.Equal(*order.StartedAt) {
		return api.Provenance{}, errors.New("request provenance occurred_at conflicts with the server start time")
	}
	if provenance.Source == "" {
		return api.Provenance{}, errors.New("request template requires an explicit native provenance source")
	}
	provenance.RunId = *order.Id
	provenance.OccurredAt = *order.StartedAt
	return provenance, nil
}
