// Package ops runs traust work for a consumer: it holds an order's lease, calls the consumer's
// Executor, and records and completes the order through core. Core selects and authorizes work.
package ops

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/traust-security/traust-sdk/go/v1/traust"
)

const (
	defaultRequestTimeout = time.Minute
	maxLeaseSeconds       = 24 * 60 * 60
)

// Executor performs the work. On cancellation it stops the work, including child processes,
// waits for it to stop, and returns.
type Executor interface {
	Execute(context.Context, Execution) (ProducedReport, error)
}

// RecordingPreparer selects a recording template and its explicit bindings for a queued order.
// The SDK calls it under the held lease and prepares any domain input before execution.
type RecordingPreparer interface {
	PrepareRecording(context.Context, traust.RunOrder) (RecordingRequest, error)
}

// PreparationRefusedError means a preparer definitively cannot run this order.
// Err is a consumer diagnostic logged by Runner.Run unless it contains an HTTP problem; keep it secret-free.
type PreparationRefusedError struct {
	Reason traust.RunOrderReason
	Err    error
}

func (e *PreparationRefusedError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("ops: recording preparation refused (%s): %v", e.Reason, e.Err)
	}
	return fmt.Sprintf("ops: recording preparation refused (%s)", e.Reason)
}

func (e *PreparationRefusedError) Unwrap() error { return e.Err }

// Execution is what an Executor runs: server context, never credentials.
type Execution struct {
	Order       traust.RunOrder     `json:"order"`
	TriageInput *traust.TriageInput `json:"triage_input,omitempty"`
	Input       json.RawMessage     `json:"input,omitempty"`
	Argv        []string            `json:"argv,omitempty"`
}

// ProducedReport is the executor's outcome and its original output, byte for byte.
// Recording runs send Payload to their record door without rewriting it.
type ProducedReport struct {
	Payload []byte
	Result  traust.JobResult
}

// LeaseLostError means the order is no longer this process's: nothing may be recorded or completed.
type LeaseLostError struct{ Err error }

func (e *LeaseLostError) Error() string { return "ops: order lease lost: " + e.Err.Error() }
func (e *LeaseLostError) Unwrap() error { return e.Err }

// unresolvedParam means the order cannot supply a parameter its command names, or names a
// repository this runner cannot resolve through core: the order is refused, never run. The
// reason is the refusal's classifier — bad_request for an order that lacks a value or whose
// repository URL core would not parse, not_found for a repository core will not hand over.
type unresolvedParam struct {
	reason traust.RunOrderReason
	err    error
}

func (e *unresolvedParam) Error() string { return "ops: unresolved parameter: " + e.err.Error() }
func (e *unresolvedParam) Unwrap() error { return e.err }

func lost(err error) bool {
	var lease *LeaseLostError
	return errors.As(err, &lease)
}

func checkOutcome(result traust.JobResult) error {
	switch result.Outcome {
	case traust.OutcomeSucceeded, traust.OutcomeDegraded, traust.OutcomeRefused, traust.OutcomeFailed:
		return nil
	}
	return fmt.Errorf("ops: executor returned outcome %q", result.Outcome)
}

func leaseDuration(seconds int) time.Duration { return time.Duration(seconds) * time.Second }
