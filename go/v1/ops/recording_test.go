package ops

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/traust-security/traust-sdk/go/v1/traust"
)

type executorFunc func(context.Context, Execution) (ProducedReport, error)

func (execute executorFunc) Execute(ctx context.Context, execution Execution) (ProducedReport, error) {
	return execute(ctx, execution)
}

type lifecycleServer struct {
	mu              sync.Mutex
	order           traust.RunOrder
	leaseLength     time.Duration
	heartbeatStatus int
	recordStatus    int
	disconnect      bool
	changeOrigin    bool
	events          []string
	completed       traust.JobResult
}

func (server *lifecycleServer) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	server.mu.Lock()
	defer server.mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	switch {
	case request.URL.Path == "/v1/orders/start":
		server.events = append(server.events, "start")
		expires := time.Now().Add(server.leaseLength)
		server.order.LeaseExpiresAt = &expires
		json.NewEncoder(w).Encode(server.order)
	case strings.HasSuffix(request.URL.Path, "/heartbeat"):
		server.events = append(server.events, "heartbeat")
		if server.heartbeatStatus != 0 {
			w.Header().Set("Content-Type", "application/problem+json")
			w.WriteHeader(server.heartbeatStatus)
			json.NewEncoder(w).Encode(traust.Problem{Title: "lease refused", Status: server.heartbeatStatus})
			return
		}
		expires := time.Now().Add(server.leaseLength)
		server.order.LeaseExpiresAt = &expires
		if server.changeOrigin {
			changed := server.order.StartedAt.Add(time.Second)
			server.order.StartedAt = &changed
		}
		json.NewEncoder(w).Encode(server.order)
	case request.URL.Path == "/v1/assessment-records":
		server.events = append(server.events, "record")
		if server.disconnect {
			connection, _, err := w.(http.Hijacker).Hijack()
			if err == nil {
				connection.Close()
			}
			return
		}
		if server.recordStatus != 0 {
			w.Header().Set("Content-Type", "application/problem+json")
			w.WriteHeader(server.recordStatus)
			json.NewEncoder(w).Encode(traust.Problem{Title: "record refused", Status: server.recordStatus})
			return
		}
		json.NewEncoder(w).Encode(traust.AssessmentRecordResult{
			AssessmentId: uuid.MustParse("cbb60c79-1c48-4938-b71e-e19d0b9b9ca7"),
			ArtifactId:   uuid.MustParse("06415613-84a6-4a90-8e39-a7a5d8d5fcef"),
		})
	case strings.HasSuffix(request.URL.Path, "/complete"):
		server.events = append(server.events, "complete")
		var body traust.CompleteRequest
		if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
			w.WriteHeader(400)
			return
		}
		server.completed = body.Result
		server.order.State = traust.RunOrderState(body.Result.Outcome)
		json.NewEncoder(w).Encode(server.order)
	default:
		server.events = append(server.events, "unexpected")
		w.WriteHeader(404)
	}
}

func newLifecycleRun(t *testing.T, executor Executor) (RecordingRun, *lifecycleServer) {
	t.Helper()
	id := uuid.MustParse("dfc7fdca-cba2-4a54-a9d3-4b57041bbb5f")
	started := time.Date(2026, 10, 28, 1, 2, 3, 123456789, time.UTC)
	server := &lifecycleServer{
		order:       traust.RunOrder{Id: &id, StartedAt: &started, State: traust.RunOrderStateLeased},
		leaseLength: time.Second,
	}
	httpServer := httptest.NewServer(server)
	t.Cleanup(httpServer.Close)
	tokenFile := filepath.Join(t.TempDir(), "token")
	if err := os.WriteFile(tokenFile, []byte("test-token"), 0600); err != nil {
		t.Fatal(err)
	}
	client, err := traust.NewClient(traust.Config{BaseURL: httpServer.URL, Tokens: traust.FileTokenSource{Path: tokenFile}})
	if err != nil {
		t.Fatal(err)
	}
	return RecordingRun{
		Client: client, Start: traust.StartRecordingRequest{LeaseSeconds: 30, Order: traust.OrderRequest{Asset: "explicit-asset"}},
		Request: RecordingRequest{Assessment: &traust.AssessmentRecordRequest{
			Provenance: traust.Provenance{Source: traust.SourceTypeVulnScanReport}, AssessorVersion: "test-assessor",
		}},
		Executor: executor, HeartbeatInterval: 10 * time.Millisecond, RequestTimeout: time.Second,
	}, server
}

func (server *lifecycleServer) counts() (heartbeats, records, completions int) {
	server.mu.Lock()
	defer server.mu.Unlock()
	for _, event := range server.events {
		switch event {
		case "heartbeat":
			heartbeats++
		case "record":
			records++
		case "complete":
			completions++
		}
	}
	return
}

func TestRecordingRunMaintainsLeaseAndCompletesNativeOutcome(t *testing.T) {
	run, server := newLifecycleRun(t, executorFunc(func(ctx context.Context, execution Execution) (ProducedReport, error) {
		if execution.Order.Id == nil || execution.Order.StartedAt == nil {
			return ProducedReport{}, errors.New("missing server provenance")
		}
		select {
		case <-ctx.Done():
			return ProducedReport{}, ctx.Err()
		case <-time.After(45 * time.Millisecond):
			return ProducedReport{Payload: []byte("original report"), Result: traust.JobResult{Outcome: traust.OutcomeDegraded}}, nil
		}
	}))
	result, err := run.Run(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	heartbeats, records, completions := server.counts()
	if heartbeats < 3 || records != 1 || completions != 1 || result.Receipt.Assessment == nil || result.Result.Outcome != traust.OutcomeDegraded {
		t.Fatal("run did not heartbeat during execution and record/complete exactly once")
	}
	if result.Order.Id == nil || *result.Order.Id != *server.order.Id || !result.Order.StartedAt.Equal(*server.order.StartedAt) {
		t.Fatal("native run identity/provenance changed")
	}
}

func TestHeartbeatRefusalStopsExecutorBeforeReturningAndNeverWrites(t *testing.T) {
	stopped := make(chan struct{})
	run, server := newLifecycleRun(t, executorFunc(func(ctx context.Context, _ Execution) (ProducedReport, error) {
		<-ctx.Done()
		close(stopped)
		return ProducedReport{}, ctx.Err()
	}))
	server.heartbeatStatus = 403
	result, err := run.Run(context.Background())
	var lost *LeaseLostError
	var problem *traust.ProblemError
	if !errors.As(err, &lost) || !errors.As(err, &problem) || problem.StatusCode != 403 || result == nil || result.Order.Id == nil {
		t.Fatal("lease loss did not preserve typed refusal and native run identity")
	}
	select {
	case <-stopped:
	default:
		t.Fatal("run returned before executor stopped")
	}
	_, records, completions := server.counts()
	if records != 0 || completions != 0 {
		t.Fatal("recorded or completed after known lease loss")
	}
}

func TestLeaseExpiryCancelsWorkWithoutFurtherMutation(t *testing.T) {
	run, server := newLifecycleRun(t, executorFunc(func(ctx context.Context, _ Execution) (ProducedReport, error) {
		<-ctx.Done()
		return ProducedReport{}, ctx.Err()
	}))
	server.leaseLength = 15 * time.Millisecond
	run.HeartbeatInterval = time.Second
	_, err := run.Run(context.Background())
	var lost *LeaseLostError
	if !errors.As(err, &lost) {
		t.Fatal("expired lease was not terminal")
	}
	heartbeats, records, completions := server.counts()
	if heartbeats != 0 || records != 0 || completions != 0 {
		t.Fatal("mutated an expired order")
	}
}

func TestChangedServerOriginIsLeaseLossNotNewProvenance(t *testing.T) {
	run, server := newLifecycleRun(t, executorFunc(func(ctx context.Context, _ Execution) (ProducedReport, error) {
		<-ctx.Done()
		return ProducedReport{}, ctx.Err()
	}))
	server.changeOrigin = true
	_, err := run.Run(context.Background())
	var lost *LeaseLostError
	if !errors.As(err, &lost) {
		t.Fatal("heartbeat was permitted to replace the immutable provenance origin")
	}
	_, records, completions := server.counts()
	if records != 0 || completions != 0 {
		t.Fatal("mutated after inconsistent heartbeat")
	}
}

func TestExecutorFailureCompletesFailedOnlyWhileStillOwned(t *testing.T) {
	run, server := newLifecycleRun(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
		return ProducedReport{}, errors.New("executor crashed")
	}))
	_, err := run.Run(context.Background())
	if err == nil {
		t.Fatal("executor failure was hidden")
	}
	_, records, completions := server.counts()
	if records != 0 || completions != 1 || server.completed.Outcome != traust.OutcomeFailed || server.completed.Reason != nil {
		t.Fatal("executor failure did not use native failed outcome without an invented reason classifier")
	}
}

func TestAmbiguousRecordingFailurePreservesRunWithoutRetryOrCompletion(t *testing.T) {
	for _, disconnect := range []bool{false, true} {
		run, server := newLifecycleRun(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
			return ProducedReport{Payload: []byte("original report"), Result: traust.JobResult{Outcome: traust.OutcomeSucceeded}}, nil
		}))
		server.disconnect = disconnect
		server.recordStatus = 500
		result, err := run.Run(context.Background())
		if err == nil || result == nil || result.Order.Id == nil {
			t.Fatal("ambiguous recording failure lost the native reconciliation identity")
		}
		_, records, completions := server.counts()
		if records != 1 || completions != 0 {
			t.Fatal("ambiguous recording POST was retried or marked failed/complete")
		}
	}
}

func TestDeterministicRefusalCanCompleteFailedAfterOwnershipCheck(t *testing.T) {
	run, server := newLifecycleRun(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
		return ProducedReport{Payload: []byte("original report"), Result: traust.JobResult{Outcome: traust.OutcomeSucceeded}}, nil
	}))
	server.recordStatus = 422
	_, err := run.Run(context.Background())
	var problem *traust.ProblemError
	if !errors.As(err, &problem) || problem.StatusCode != 422 {
		t.Fatal("native deterministic refusal was not preserved")
	}
	heartbeats, records, completions := server.counts()
	if heartbeats != 2 || records != 1 || completions != 1 || server.completed.Outcome != traust.OutcomeFailed {
		t.Fatal("refusal did not recheck ownership and complete using native outcome")
	}
	// The refused payload is attributable from the order row alone.
	if server.completed.Reason == nil || *server.completed.Reason != traust.RunOrderReasonBadRequest {
		t.Fatalf("record refusal did not stamp a reason: %+v", server.completed.Reason)
	}
}

func TestUnknownExecutorOutcomeCompletesFailedWithoutRecording(t *testing.T) {
	run, server := newLifecycleRun(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
		return ProducedReport{Payload: []byte("original report"), Result: traust.JobResult{Outcome: "exploded"}}, nil
	}))
	if _, err := run.Run(context.Background()); err == nil {
		t.Fatal("an unknown outcome was accepted")
	}
	_, records, completions := server.counts()
	if records != 0 || completions != 1 || server.completed.Outcome != traust.OutcomeFailed {
		t.Fatalf("records=%d completions=%d completed=%+v", records, completions, server.completed)
	}
}
