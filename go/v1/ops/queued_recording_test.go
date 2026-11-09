package ops

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"slices"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/traust-security/traust-sdk/go/v1/traust"
)

type recordingPreparerFunc func(context.Context, traust.RunOrder) (RecordingRequest, error)

func (prepare recordingPreparerFunc) PrepareRecording(ctx context.Context, order traust.RunOrder) (RecordingRequest, error) {
	return prepare(ctx, order)
}

type queuedRecordingServer struct {
	queue                  *queuedServer
	assessmentID           uuid.UUID
	prepareStatus          int
	problemDetail          string
	recordStatus           int
	malformedPrepare       bool
	malformedRecord        bool
	disconnectPrepare      bool
	disconnectRecord       bool
	prepareGate            <-chan struct{}
	recordGate             <-chan struct{}
	prepareStarted         chan struct{}
	recordStarted          chan struct{}
	requestStopped         chan struct{}
	heartbeat              chan struct{}
	recorded               traust.RemediationRecordBody
	recordPath             string
	preparePath            string
	repositoryID           uuid.UUID
	fixRefID               uuid.UUID
	verificationInput      *traust.VerificationInput
	remediationInput       *traust.RemediationInput
	registrationReply      *traust.RepositoryRef
	refs                   []traust.RepositoryRef
	lookupStatus           int
	registerStatus         int
	disconnectRegister     bool
	mismatchedRegistration bool
	registered             traust.RefRegistration
	registerGate           <-chan struct{}
	registerStarted        chan struct{}
}

func (server *queuedRecordingServer) event(event string) {
	server.queue.mu.Lock()
	defer server.queue.mu.Unlock()
	server.queue.events = append(server.queue.events, event)
}

func signalRecordingStage(stage chan struct{}) {
	select {
	case stage <- struct{}{}:
	default:
	}
}

func (server *queuedRecordingServer) wait(request *http.Request, gate <-chan struct{}) bool {
	if gate == nil {
		return true
	}
	select {
	case <-gate:
		return true
	case <-request.Context().Done():
		signalRecordingStage(server.requestStopped)
		return false
	}
}

func (server *queuedRecordingServer) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	path := request.URL.Path
	if path == "/v1/orders/start" {
		server.queue.mu.Lock()
		defer server.queue.mu.Unlock()
		server.queue.events = append(server.queue.events, "start")
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(server.queue.order)
		return
	}
	if strings.HasSuffix(path, "/heartbeat") {
		server.queue.ServeHTTP(w, request)
		signalRecordingStage(server.heartbeat)
		return
	}
	if strings.HasSuffix(path, "-input") {
		server.event("prepare")
		server.queue.mu.Lock()
		server.preparePath = path
		server.queue.mu.Unlock()
		signalRecordingStage(server.prepareStarted)
		if !server.wait(request, server.prepareGate) {
			return
		}
		if server.disconnectPrepare {
			connection, _, err := w.(http.Hijacker).Hijack()
			if err == nil {
				_ = connection.Close()
			}
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if server.prepareStatus != 0 {
			w.Header().Set("Content-Type", "application/problem+json")
			w.WriteHeader(server.prepareStatus)
			_ = json.NewEncoder(w).Encode(traust.Problem{Title: "preparation unavailable", Status: server.prepareStatus, Detail: server.problemDetail})
			return
		}
		if server.malformedPrepare {
			_, _ = w.Write([]byte(`{"items":`))
			return
		}
		server.queue.mu.Lock()
		workItem := *server.queue.order.WorkItemId
		server.queue.mu.Unlock()
		switch {
		case path == "/v1/assessments/"+server.assessmentID.String()+"/triage-input":
			_ = json.NewEncoder(w).Encode(traust.TriageInput{Assessor: "core-assessor"})
		case path == "/v1/work-items/"+workItem.String()+"/remediation-input":
			_ = json.NewEncoder(w).Encode(server.remediationInput)
		case path == "/v1/work-items/"+workItem.String()+"/verification-input":
			_ = json.NewEncoder(w).Encode(server.verificationInput)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
		return
	}
	if path == "/v1/repositories/"+server.repositoryID.String()+"/refs" {
		w.Header().Set("Content-Type", "application/json")
		if request.Method == http.MethodGet {
			server.event("find-ref")
			if server.lookupStatus != 0 {
				w.Header().Set("Content-Type", "application/problem+json")
				w.WriteHeader(server.lookupStatus)
				_ = json.NewEncoder(w).Encode(traust.Problem{Title: "ref lookup unavailable", Status: server.lookupStatus, Detail: server.problemDetail})
				return
			}
			if request.URL.Query().Get("sha") != *server.verificationInput.WorkItem.FixCommit {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			_ = json.NewEncoder(w).Encode(server.refs)
			return
		}
		if request.Method == http.MethodPost {
			server.event("register-ref")
			var body traust.RefRegistration
			if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			server.queue.mu.Lock()
			server.registered = body
			server.queue.mu.Unlock()
			signalRecordingStage(server.registerStarted)
			if !server.wait(request, server.registerGate) {
				return
			}
			if server.disconnectRegister {
				connection, _, err := w.(http.Hijacker).Hijack()
				if err == nil {
					_ = connection.Close()
				}
				return
			}
			if server.registerStatus != 0 {
				w.Header().Set("Content-Type", "application/problem+json")
				w.WriteHeader(server.registerStatus)
				_ = json.NewEncoder(w).Encode(traust.Problem{Title: "ref registration unavailable", Status: server.registerStatus, Detail: server.problemDetail})
				return
			}
			ref := server.preparedFixRef()
			if server.registrationReply != nil {
				ref = *server.registrationReply
			}
			if server.mismatchedRegistration {
				ref.Sha = strings.Repeat("d", 40)
			}
			_ = json.NewEncoder(w).Encode(ref)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if strings.HasSuffix(path, "-records") {
		server.event("record")
		var body traust.RemediationRecordBody
		if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		server.queue.mu.Lock()
		server.recorded, server.recordPath = body, path
		server.queue.mu.Unlock()
		signalRecordingStage(server.recordStarted)
		if !server.wait(request, server.recordGate) {
			return
		}
		if server.disconnectRecord {
			connection, _, err := w.(http.Hijacker).Hijack()
			if err == nil {
				_ = connection.Close()
			}
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if server.recordStatus != 0 {
			w.Header().Set("Content-Type", "application/problem+json")
			w.WriteHeader(server.recordStatus)
			_ = json.NewEncoder(w).Encode(traust.Problem{Title: "record unavailable", Status: server.recordStatus, Detail: server.problemDetail})
			return
		}
		if server.malformedRecord {
			_, _ = w.Write([]byte(`{"artifact_id":`))
			return
		}
		_, _ = w.Write([]byte(`{}`))
		return
	}
	server.queue.ServeHTTP(w, request)
}

func newQueuedRecordingRunner(t *testing.T, executor Executor, preparer RecordingPreparer) (*Runner, *queuedRecordingServer) {
	t.Helper()
	pack, err := LoadPack([]byte(ledgerPack))
	if err != nil {
		t.Fatal(err)
	}
	id, workItemID := uuid.New(), uuid.New()
	started, expires := time.Now(), time.Now().Add(time.Minute)
	server := &queuedRecordingServer{
		queue: &queuedServer{pack: pack, asset: "ledger_audit", order: traust.RunOrder{
			Id: &id, StartedAt: &started, LeaseExpiresAt: &expires, State: traust.RunOrderStateLeased,
			Asset: "ledger_audit", WorkItemId: &workItemID,
		}},
		assessmentID: uuid.New(), prepareStarted: make(chan struct{}, 8), recordStarted: make(chan struct{}, 8),
		requestStopped: make(chan struct{}, 8), heartbeat: make(chan struct{}, 64),
		repositoryID: uuid.New(), fixRefID: uuid.New(),
		refs:            []traust.RepositoryRef{},
		registerStarted: make(chan struct{}, 8),
	}
	commit, branch, baselineID := strings.Repeat("a", 40), "fix/authoritative", uuid.New()
	server.verificationInput = &traust.VerificationInput{
		Subject:  traust.SubjectRef{RepositoryId: server.repositoryID, RepositoryRefId: baselineID, Sha: strings.Repeat("b", 40)},
		WorkItem: traust.WorkItem{WorkItemId: workItemID, RepositoryRefId: baselineID, FixCommit: &commit, FixBranch: &branch},
	}
	server.remediationInput = &traust.RemediationInput{
		Subject: server.verificationInput.Subject, FixBranch: "core-fix-branch",
		WorkItem: traust.WorkItem{WorkItemId: workItemID, FindingIds: []uuid.UUID{uuid.New()}},
	}
	httpServer := httptest.NewServer(server)
	t.Cleanup(httpServer.Close)
	tokenFile := filepath.Join(t.TempDir(), "token")
	if err := os.WriteFile(tokenFile, []byte("test-token"), 0o600); err != nil {
		t.Fatal(err)
	}
	client, err := traust.NewClient(traust.Config{
		BaseURL: httpServer.URL, Tokens: traust.FileTokenSource{Path: tokenFile}, Backoff: time.Millisecond,
	})
	if err != nil {
		t.Fatal(err)
	}
	cfg := RunnerConfig{
		Client: client, Pack: pack, Executor: executor, Assets: []string{"ledger_audit"},
		Logger:       slog.New(slog.NewTextHandler(os.Stderr, nil)),
		LeaseSeconds: 30, HeartbeatInterval: 5 * time.Millisecond, RequestTimeout: time.Second, PollInterval: time.Millisecond,
	}
	if preparer != nil {
		cfg.Recordings = map[string]RecordingPreparer{"ledger_audit": preparer}
	}
	runner, err := NewRunner(cfg)
	if err != nil {
		t.Fatal(err)
	}
	return runner, server
}

func assessmentRecordingTemplate() RecordingRequest {
	return RecordingRequest{Assessment: &traust.AssessmentRecordRequest{
		AssessorVersion: "test-assessor", Provenance: traust.Provenance{Source: traust.SourceTypeVulnScanReport},
	}}
}

func fixedRecordingTemplate(request RecordingRequest) RecordingPreparer {
	return recordingPreparerFunc(func(context.Context, traust.RunOrder) (RecordingRequest, error) { return request, nil })
}

func recordingEvents(server *queuedRecordingServer) ([]string, traust.JobResult) {
	events, result := server.queue.snapshot()
	return slices.DeleteFunc(events, func(event string) bool { return event == "heartbeat" }), result
}

func TestQueuedRecordingPreparesExecutesRecordsAndCompletesOriginalReport(t *testing.T) {
	for _, kind := range []string{"assessment", "triage", "remediation", "verification", "validation"} {
		t.Run(kind, func(t *testing.T) {
			original := []byte(" \r\n{\"finding_ref\":\"worker-label\",\"text\":\"é <tag> & \\u0061\",\"n\":1.00}\t\n")
			var server *queuedRecordingServer
			degraded := map[string]int{"partial": 2}
			result := traust.JobResult{Outcome: traust.OutcomeDegraded, Degraded: &degraded}
			runner, core := newQueuedRecordingRunner(t, executorFunc(func(_ context.Context, execution Execution) (ProducedReport, error) {
				server.event("execute")
				if execution.Order.Id == nil || execution.Order.StartedAt == nil || len(execution.Argv) == 0 {
					return ProducedReport{}, errors.New("missing leased order or pinned command")
				}
				switch kind {
				case "triage":
					if execution.TriageInput == nil || execution.TriageInput.Assessor != "core-assessor" || execution.Input != nil {
						return ProducedReport{}, errors.New("triage did not receive the selected assessment's prepared input")
					}
				case "remediation":
					var input traust.RemediationInput
					if err := json.Unmarshal(execution.Input, &input); err != nil || input.FixBranch != "core-fix-branch" || input.WorkItem.WorkItemId != *execution.Order.WorkItemId {
						return ProducedReport{}, errors.New("remediation did not receive core's work-item input")
					}
				case "verification":
					var input traust.VerificationInput
					if err := json.Unmarshal(execution.Input, &input); err != nil || input.WorkItem.WorkItemId != *execution.Order.WorkItemId {
						return ProducedReport{}, errors.New("verification did not receive core's work-item input")
					}
				default:
					if string(execution.Input) != `{"consumer":"input"}` {
						return ProducedReport{}, errors.New("explicit consumer input was not transported")
					}
				}
				// Consumer mutations must not replace the server identity or work-item selection.
				*execution.Order.Id = uuid.New()
				*execution.Order.StartedAt = time.Time{}
				return ProducedReport{Payload: original, Result: result}, nil
			}), recordingPreparerFunc(func(_ context.Context, order traust.RunOrder) (RecordingRequest, error) {
				server.event("template")
				if order.Id == nil || order.StartedAt == nil || order.WorkItemId == nil {
					return RecordingRequest{}, errors.New("preparer did not receive server context")
				}
				*order.WorkItemId = uuid.New()
				request := RecordingRequest{Input: json.RawMessage(`{"consumer":"input"}`)}
				provenance := traust.Provenance{Source: traust.SourceTypeTriageReport}
				switch kind {
				case "assessment":
					request.Assessment = &traust.AssessmentRecordRequest{AssessorVersion: "test-assessor", Provenance: provenance}
				case "triage":
					request.AssessmentID = server.assessmentID
					request.Triage = &traust.TriageRecordBody{Provenance: provenance}
				case "remediation":
					request.Remediation = &traust.RemediationRecordBody{Provenance: provenance}
				case "verification":
					request.Verification = &traust.VerificationRecordBody{Provenance: provenance}
				case "validation":
					request.Validation = &traust.RecordValidation{Provenance: provenance}
				}
				return request, nil
			}))
			server = core
			order, err := runner.RunOnce(context.Background())
			if err != nil || order == nil || order.State != traust.RunOrderStateDegraded {
				t.Fatalf("order=%+v err=%v", order, err)
			}
			events, completed := recordingEvents(server)
			want := []string{"claim", "manifest", "template"}
			if kind == "triage" || kind == "remediation" || kind == "verification" {
				want = append(want, "prepare")
			}
			if kind == "verification" {
				want = append(want, "find-ref", "register-ref")
			}
			want = append(want, "execute", "record", "complete")
			if !slices.Equal(events, want) || !reflect.DeepEqual(completed, result) || completed.Usage != nil {
				t.Fatalf("events=%v result=%+v want=%v/%+v", events, completed, want, result)
			}
			body := server.recorded
			if !bytes.Equal([]byte(body.Payload), original) || body.RunOrderId != *order.Id || body.Provenance.RunId != *order.Id || !body.Provenance.OccurredAt.Equal(*order.StartedAt) || body.Provenance.Source != traust.SourceTypeTriageReport {
				t.Fatalf("record did not preserve original bytes and server provenance: %+v", body)
			}
			wantPath := "/v1/" + kind + "-records"
			if kind == "triage" {
				wantPath = "/v1/assessments/" + server.assessmentID.String() + "/triage-records"
			}
			if server.recordPath != wantPath {
				t.Fatalf("record path=%s want=%s", server.recordPath, wantPath)
			}
			wantPreparation := ""
			switch kind {
			case "triage":
				wantPreparation = "/v1/assessments/" + server.assessmentID.String() + "/triage-input"
			case "remediation", "verification":
				wantPreparation = "/v1/work-items/" + order.WorkItemId.String() + "/" + kind + "-input"
			}
			if server.preparePath != wantPreparation {
				t.Fatalf("preparation did not use the selected assessment/order work item: got=%s want=%s", server.preparePath, wantPreparation)
			}
		})
	}
}

func TestQueuedAssetsWithoutRecordingDeclarationRetainPlainExecution(t *testing.T) {
	runner, server := newQueuedRecordingRunner(t, executorFunc(func(_ context.Context, execution Execution) (ProducedReport, error) {
		if execution.Input != nil || execution.TriageInput != nil || len(execution.Argv) == 0 {
			return ProducedReport{}, errors.New("plain execution changed")
		}
		return ProducedReport{Payload: []byte{0xff}, Result: traust.JobResult{Outcome: traust.OutcomeSucceeded}}, nil
	}), nil)
	order, err := runner.RunOnce(context.Background())
	events, completed := recordingEvents(server)
	if err != nil || order == nil || completed.Outcome != traust.OutcomeSucceeded || !slices.Equal(events, []string{"claim", "manifest", "complete"}) {
		t.Fatalf("order=%+v err=%v events=%v result=%+v", order, err, events, completed)
	}
}

func TestRunnerRejectsInvalidRecordingDeclarationsAndOwnsItsMapping(t *testing.T) {
	runner, server := newQueuedRecordingRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
		return ProducedReport{Payload: []byte("report"), Result: traust.JobResult{Outcome: traust.OutcomeSucceeded}}, nil
	}), nil)
	for _, tc := range []struct {
		name       string
		recordings map[string]RecordingPreparer
	}{
		{"nil preparer", map[string]RecordingPreparer{"ledger_audit": nil}},
		{"typed nil preparer", map[string]RecordingPreparer{"ledger_audit": recordingPreparerFunc(nil)}},
		{"unclaimed asset", map[string]RecordingPreparer{"other": fixedRecordingTemplate(assessmentRecordingTemplate())}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			cfg := runner.config
			cfg.Recordings = tc.recordings
			if _, err := NewRunner(cfg); err == nil {
				t.Fatal("invalid recording declaration accepted")
			}
		})
	}
	cfg := runner.config
	cfg.Recordings = map[string]RecordingPreparer{"ledger_audit": fixedRecordingTemplate(assessmentRecordingTemplate())}
	configured, err := NewRunner(cfg)
	if err != nil {
		t.Fatal(err)
	}
	delete(cfg.Recordings, "ledger_audit")
	order, err := configured.RunOnce(context.Background())
	events, _ := recordingEvents(server)
	if err != nil || order == nil || order.State != traust.RunOrderStateSucceeded || !slices.Contains(events, "record") {
		t.Fatalf("consumer mutation removed the recording declaration: order=%+v err=%v events=%v", order, err, events)
	}
}

func TestQueuedRecordingRejectsInvalidTemplateBeforeExecution(t *testing.T) {
	for _, tc := range []struct {
		name   string
		mutate func(*RecordingRequest)
	}{
		{"missing template", func(request *RecordingRequest) { *request = RecordingRequest{} }},
		{"multiple templates", func(request *RecordingRequest) { request.Triage = &traust.TriageRecordBody{} }},
		{"prepopulated payload", func(request *RecordingRequest) { request.Assessment.Payload = "not executor output" }},
		{"missing source", func(request *RecordingRequest) { request.Assessment.Provenance.Source = "" }},
		{"wrong order", func(request *RecordingRequest) { request.Assessment.RunOrderId = uuid.New() }},
		{"wrong provenance run", func(request *RecordingRequest) { request.Assessment.Provenance.RunId = uuid.New() }},
		{"wrong provenance time", func(request *RecordingRequest) { request.Assessment.Provenance.OccurredAt = time.Unix(1, 0) }},
		{"missing assessment", func(request *RecordingRequest) {
			request.Triage = &traust.TriageRecordBody{Provenance: request.Assessment.Provenance}
			request.Assessment = nil
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			request := assessmentRecordingTemplate()
			tc.mutate(&request)
			executed := false
			runner, server := newQueuedRecordingRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
				executed = true
				return ProducedReport{}, nil
			}), fixedRecordingTemplate(request))
			order, err := runner.RunOnce(context.Background())
			events, completed := recordingEvents(server)
			if err == nil || order == nil || executed || completed.Outcome != traust.OutcomeFailed || slices.Contains(events, "prepare") || slices.Contains(events, "record") || !slices.Contains(events, "complete") {
				t.Fatalf("invalid template executed: order=%+v err=%v executed=%v events=%v result=%+v", order, err, executed, events, completed)
			}
		})
	}
}

func TestQueuedRecordingPreparationRefusalsPreserveErrorsAndUseCoreReasonValidation(t *testing.T) {
	absent := errors.New("source assessment absent")
	for _, tc := range []struct {
		name   string
		reason traust.RunOrderReason
		status int
	}{
		{"no audit", traust.RunOrderReasonNotFound, 0},
		{"source unavailable", traust.RunOrderReasonSourceUnavailable, 0},
		{"reason newly supported by Core", "new-core-reason", 0},
		{"empty reason", "", 0},
		{"reason rejected by Core", "consumer-invented", http.StatusUnprocessableEntity},
	} {
		t.Run(tc.name, func(t *testing.T) {
			executed := false
			runner, server := newQueuedRecordingRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
				executed = true
				return ProducedReport{}, nil
			}), recordingPreparerFunc(func(context.Context, traust.RunOrder) (RecordingRequest, error) {
				return RecordingRequest{}, &PreparationRefusedError{Reason: tc.reason, Err: absent}
			}))
			server.queue.completeStatus = tc.status
			order, err := runner.RunOnce(context.Background())
			var refusal *PreparationRefusedError
			events, completed := recordingEvents(server)
			attempted := tc.reason != ""
			if order == nil || !errors.Is(err, absent) || !errors.As(err, &refusal) || executed ||
				slices.Contains(events, "record") || slices.Contains(events, "complete") != attempted {
				t.Fatalf("order=%+v err=%v executed=%v events=%v", order, err, executed, events)
			}
			if attempted && (completed.Outcome != traust.OutcomeRefused || completed.Reason == nil || *completed.Reason != tc.reason) {
				t.Fatalf("refusal was not forwarded unchanged to Core: %+v", completed)
			}
			if attempted && tc.status == 0 && order.State != traust.RunOrderStateRefused {
				t.Fatalf("accepted refusal did not close the order: %+v", order)
			}
			if tc.status != 0 {
				var problem *traust.ProblemError
				if !errors.As(err, &problem) || problem.StatusCode != tc.status || order.State != traust.RunOrderStateLeased {
					t.Fatalf("Core's reason rejection was hidden or closed: order=%+v err=%v", order, err)
				}
			}
		})
	}
}

func TestQueuedWorkItemRecordingRequiresOrderBindingBeforePreparation(t *testing.T) {
	for _, kind := range []string{"remediation", "verification"} {
		for _, zero := range []bool{false, true} {
			t.Run(kind+"/"+map[bool]string{false: "missing", true: "zero"}[zero], func(t *testing.T) {
				request := RecordingRequest{}
				provenance := traust.Provenance{Source: traust.SourceTypeTriageReport}
				if kind == "remediation" {
					request.Remediation = &traust.RemediationRecordBody{Provenance: provenance}
				} else {
					request.Verification = &traust.VerificationRecordBody{Provenance: provenance}
				}
				executed := false
				runner, server := newQueuedRecordingRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
					executed = true
					return ProducedReport{}, nil
				}), fixedRecordingTemplate(request))
				server.queue.order.WorkItemId = nil
				if zero {
					id := uuid.Nil
					server.queue.order.WorkItemId = &id
				}
				rule := "do-not-derive-a-work-item-from-this-rule"
				server.queue.order.Rule = &rule
				_, err := runner.RunOnce(context.Background())
				events, completed := recordingEvents(server)
				var refusal *PreparationRefusedError
				if !errors.As(err, &refusal) || executed || slices.Contains(events, "prepare") || slices.Contains(events, "record") || completed.Outcome != traust.OutcomeRefused || completed.Reason == nil || *completed.Reason != traust.RunOrderReasonBadRequest {
					t.Fatalf("err=%v executed=%v events=%v result=%+v", err, executed, events, completed)
				}
			})
		}
	}
}

func TestQueuedPreparationProblemsNeverMasqueradeAsDomainRefusals(t *testing.T) {
	for _, status := range []int{400, 401, 403, 404, 409, 422, 429, 503} {
		t.Run(http.StatusText(status), func(t *testing.T) {
			var server *queuedRecordingServer
			executed := false
			runner, core := newQueuedRecordingRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
				executed = true
				return ProducedReport{}, nil
			}), recordingPreparerFunc(func(context.Context, traust.RunOrder) (RecordingRequest, error) {
				return RecordingRequest{Triage: &traust.TriageRecordBody{Provenance: traust.Provenance{Source: traust.SourceTypeTriageReport}}, AssessmentID: server.assessmentID}, nil
			}))
			server = core
			server.prepareStatus = status
			order, err := runner.RunOnce(context.Background())
			var problem *traust.ProblemError
			var refusal *PreparationRefusedError
			events, completed := recordingEvents(server)
			wantClosed := status == 400 || status == 404 || status == 409 || status == 422
			if order == nil || !errors.As(err, &problem) || problem.StatusCode != status || errors.As(err, &refusal) ||
				executed || slices.Contains(events, "record") || slices.Contains(events, "complete") != wantClosed {
				t.Fatalf("order=%+v err=%v executed=%v events=%v", order, err, executed, events)
			}
			if wantClosed && (completed.Outcome != traust.OutcomeFailed || completed.Reason != nil || order.State != traust.RunOrderStateFailed) {
				t.Fatalf("HTTP preparation error became a domain refusal: order=%+v result=%+v", order, completed)
			}
		})
	}
}

func TestQueuedRecordingExecutionFailuresDoNotRecord(t *testing.T) {
	crashed := errors.New("executor crashed")
	for _, tc := range []struct {
		name   string
		report ProducedReport
		err    error
	}{
		{"execution error", ProducedReport{Result: traust.JobResult{Outcome: traust.OutcomeDegraded}}, crashed},
		{"unknown outcome", ProducedReport{Payload: []byte("report"), Result: traust.JobResult{Outcome: "invented"}}, nil},
		{"invalid UTF-8", ProducedReport{Payload: []byte{0xff, 0xfe}, Result: traust.JobResult{Outcome: traust.OutcomeSucceeded}}, nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			runner, server := newQueuedRecordingRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
				return tc.report, tc.err
			}), fixedRecordingTemplate(assessmentRecordingTemplate()))
			_, err := runner.RunOnce(context.Background())
			events, completed := recordingEvents(server)
			if err == nil || (tc.err != nil && !errors.Is(err, tc.err)) || slices.Contains(events, "record") || completed.Outcome != traust.OutcomeFailed || completed.Usage != nil {
				t.Fatalf("err=%v events=%v result=%+v", err, events, completed)
			}
			if tc.name != "invalid UTF-8" && (completed.Reason == nil || *completed.Reason != traust.RunOrderReasonRunnerCrashed) {
				t.Fatalf("executor crash classifier lost: %+v", completed)
			}
		})
	}
}

func TestQueuedRecordingPostsOnceAndSeparatesDefinitiveFromAmbiguousFailures(t *testing.T) {
	for _, tc := range []struct {
		name       string
		status     int
		disconnect bool
		malformed  bool
		complete   bool
	}{
		{"payload refusal", 422, false, false, true},
		{"access refusal", 403, false, false, false},
		{"rate refusal", 429, false, false, false},
		{"server error", 500, false, false, false},
		{"transport lost after post", 0, true, false, false},
		{"malformed success receipt", 0, false, true, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			original := []byte("malformed report\r\n finding_ref: worker-chosen\n")
			runner, server := newQueuedRecordingRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
				return ProducedReport{Payload: original, Result: traust.JobResult{Outcome: traust.OutcomeSucceeded}}, nil
			}), fixedRecordingTemplate(assessmentRecordingTemplate()))
			server.recordStatus, server.disconnectRecord, server.malformedRecord = tc.status, tc.disconnect, tc.malformed
			order, err := runner.RunOnce(context.Background())
			events, completed := recordingEvents(server)
			if err == nil || order == nil || order.Id == nil || order.StartedAt == nil || slices.Contains(events, "complete") != tc.complete {
				t.Fatalf("order=%+v err=%v events=%v", order, err, events)
			}
			posts := 0
			for _, event := range events {
				if event == "record" {
					posts++
				}
			}
			if posts != 1 || !bytes.Equal([]byte(server.recorded.Payload), original) {
				t.Fatalf("record retried or normalized: posts=%d payload=%q", posts, server.recorded.Payload)
			}
			if tc.status != 0 {
				var problem *traust.ProblemError
				if !errors.As(err, &problem) || problem.StatusCode != tc.status {
					t.Fatalf("typed record refusal lost: %v", err)
				}
			}
			if tc.complete && completed.Outcome != traust.OutcomeFailed {
				t.Fatalf("definitive refusal did not complete failed: %+v", completed)
			}
			if tc.status == 422 && (completed.Reason == nil || *completed.Reason != traust.RunOrderReasonSessionOutputInvalid) {
				t.Fatalf("payload refusal reason lost: %+v", completed)
			}
			if tc.status != 422 && completed.Reason != nil {
				t.Fatalf("invented payload refusal classifier: %+v", completed)
			}
		})
	}
}

func awaitRecordingStage(t *testing.T, stage <-chan struct{}) {
	t.Helper()
	select {
	case <-stage:
	case <-time.After(2 * time.Second):
		t.Fatal("recording stage did not make progress")
	}
}

func awaitRecordingHeartbeat(t *testing.T, server *queuedRecordingServer) {
	t.Helper()
	server.queue.mu.Lock()
	baseline := server.queue.heartbeatCalls
	server.queue.mu.Unlock()
	deadline := time.NewTimer(2 * time.Second)
	defer deadline.Stop()
	for {
		select {
		case <-server.heartbeat:
			server.queue.mu.Lock()
			count := server.queue.heartbeatCalls
			server.queue.mu.Unlock()
			if count > baseline {
				return
			}
		case <-deadline.C:
			t.Fatal("lease was not renewed while the request was blocked")
		}
	}
}

func TestQueuedRecordingRenewsShortLeasesBeyondOriginalRequestExpiry(t *testing.T) {
	for _, stage := range []string{"template", "input", "fix-ref", "record"} {
		t.Run(stage, func(t *testing.T) {
			gate, started := make(chan struct{}), make(chan struct{}, 1)
			var server *queuedRecordingServer
			runner, core := newQueuedRecordingRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
				return ProducedReport{Payload: []byte("report"), Result: traust.JobResult{Outcome: traust.OutcomeSucceeded}}, nil
			}), recordingPreparerFunc(func(ctx context.Context, _ traust.RunOrder) (RecordingRequest, error) {
				if stage == "template" {
					signalRecordingStage(started)
					select {
					case <-gate:
					case <-ctx.Done():
						return RecordingRequest{}, ctx.Err()
					}
				}
				if stage == "fix-ref" {
					return verificationRecordingTemplate(), nil
				}
				return RecordingRequest{Triage: &traust.TriageRecordBody{Provenance: traust.Provenance{Source: traust.SourceTypeTriageReport}}, AssessmentID: server.assessmentID}, nil
			}))
			server = core
			server.queue.leaseLength = 200 * time.Millisecond
			if stage == "input" {
				server.prepareGate, server.prepareStarted = gate, started
			} else if stage == "record" {
				server.recordGate, server.recordStarted = gate, started
			} else if stage == "fix-ref" {
				server.registerGate, server.registerStarted = gate, started
			}
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			done := make(chan error, 1)
			go func() { _, err := runner.RunOnce(ctx); done <- err }()
			awaitRecordingStage(t, started)
			server.queue.mu.Lock()
			originalExpiry := *server.queue.order.LeaseExpiresAt
			server.queue.mu.Unlock()
			pastOriginalExpiry := time.NewTimer(time.Until(originalExpiry) + 30*time.Millisecond)
			defer pastOriginalExpiry.Stop()
			select {
			case err := <-done:
				t.Fatalf("request stopped at the old expiry despite renewal: %v", err)
			case <-pastOriginalExpiry.C:
			}
			awaitRecordingHeartbeat(t, server)
			close(gate)
			select {
			case err := <-done:
				if err != nil {
					t.Fatal(err)
				}
			case <-time.After(2 * time.Second):
				t.Fatal("recording did not finish after release")
			}
			events, completed := recordingEvents(server)
			if !slices.Contains(events, "record") || !slices.Contains(events, "complete") || completed.Outcome != traust.OutcomeSucceeded {
				t.Fatalf("events=%v result=%+v", events, completed)
			}
		})
	}
}

func TestQueuedRecordingCancellationOrLeaseLossStopsAndWaitsWithoutCompletion(t *testing.T) {
	for _, stage := range []string{"template", "input", "fix-ref", "execute", "record"} {
		for _, loseLease := range []bool{false, true} {
			name := stage + "/cancel"
			if loseLease {
				name = stage + "/lease-loss"
			}
			t.Run(name, func(t *testing.T) {
				started, stopped := make(chan struct{}, 1), make(chan struct{}, 1)
				gate := make(chan struct{})
				var server *queuedRecordingServer
				runner, core := newQueuedRecordingRunner(t, executorFunc(func(ctx context.Context, _ Execution) (ProducedReport, error) {
					server.event("execute")
					if stage == "execute" {
						signalRecordingStage(started)
						<-ctx.Done()
						signalRecordingStage(stopped)
						return ProducedReport{}, ctx.Err()
					}
					return ProducedReport{Payload: []byte("report"), Result: traust.JobResult{Outcome: traust.OutcomeSucceeded}}, nil
				}), recordingPreparerFunc(func(ctx context.Context, _ traust.RunOrder) (RecordingRequest, error) {
					if stage == "template" {
						signalRecordingStage(started)
						<-ctx.Done()
						signalRecordingStage(stopped)
						return RecordingRequest{}, ctx.Err()
					}
					if stage == "fix-ref" {
						return verificationRecordingTemplate(), nil
					}
					return RecordingRequest{Triage: &traust.TriageRecordBody{Provenance: traust.Provenance{Source: traust.SourceTypeTriageReport}}, AssessmentID: server.assessmentID}, nil
				}))
				server = core
				if stage == "input" {
					server.prepareGate, server.prepareStarted = gate, started
					server.requestStopped = stopped
				} else if stage == "record" {
					server.recordGate, server.recordStarted = gate, started
					server.requestStopped = stopped
				} else if stage == "fix-ref" {
					server.registerGate, server.registerStarted = gate, started
					server.requestStopped = stopped
				}
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()
				type outcome struct {
					order *traust.RunOrder
					err   error
				}
				done := make(chan outcome, 1)
				go func() { order, err := runner.RunOnce(ctx); done <- outcome{order, err} }()
				awaitRecordingStage(t, started)
				if loseLease {
					server.queue.mu.Lock()
					server.queue.heartbeatStatus = http.StatusConflict
					server.queue.mu.Unlock()
				} else {
					cancel()
				}
				select {
				case result := <-done:
					var lost *LeaseLostError
					if result.order == nil || result.order.Id == nil || !errors.As(result.err, &lost) {
						t.Fatalf("lost reconciliation identity or typed lease error: %+v", result)
					}
					if loseLease {
						var problem *traust.ProblemError
						if !errors.As(result.err, &problem) || problem.StatusCode != http.StatusConflict {
							t.Fatalf("heartbeat refusal was not preserved: %v", result.err)
						}
					}
				case <-time.After(2 * time.Second):
					t.Fatal("lost lease did not stop and wait for work")
				}
				awaitRecordingStage(t, stopped)
				events, _ := recordingEvents(server)
				if slices.Contains(events, "complete") || (stage != "record" && slices.Contains(events, "record")) {
					t.Fatalf("wrote after lease loss: %v", events)
				}
				if stage != "execute" && stage != "record" && slices.Contains(events, "execute") {
					t.Fatalf("executed after preparation lost its lease: %v", events)
				}
			})
		}
	}
}

func TestQueuedPreparationKeepsRequestDeadlineAndMalformedInputDoesNotExecute(t *testing.T) {
	for _, malformed := range []bool{false, true} {
		name := "request deadline"
		if malformed {
			name = "malformed input"
		}
		t.Run(name, func(t *testing.T) {
			var server *queuedRecordingServer
			executed := false
			runner, core := newQueuedRecordingRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
				executed = true
				return ProducedReport{}, nil
			}), recordingPreparerFunc(func(ctx context.Context, _ traust.RunOrder) (RecordingRequest, error) {
				if _, ok := ctx.Deadline(); !ok {
					return RecordingRequest{}, errors.New("preparer request was unbounded")
				}
				return RecordingRequest{Triage: &traust.TriageRecordBody{Provenance: traust.Provenance{Source: traust.SourceTypeTriageReport}}, AssessmentID: server.assessmentID}, nil
			}))
			server = core
			server.malformedPrepare = malformed
			if !malformed {
				server.prepareGate = make(chan struct{})
				runner.config.RequestTimeout = 100 * time.Millisecond
			}
			order, err := runner.RunOnce(context.Background())
			events, _ := recordingEvents(server)
			if err == nil || order == nil || executed || slices.Contains(events, "record") || slices.Contains(events, "complete") {
				t.Fatalf("order=%+v err=%v executed=%v events=%v", order, err, executed, events)
			}
			if !malformed && !errors.Is(err, context.DeadlineExceeded) {
				t.Fatalf("request deadline was not preserved: %v", err)
			}
		})
	}
}

func TestQueuedRecordingTemplateAndRecordRequestsKeepTheirDeadlines(t *testing.T) {
	for _, stage := range []string{"template", "record"} {
		t.Run(stage, func(t *testing.T) {
			executed := false
			runner, server := newQueuedRecordingRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
				executed = true
				return ProducedReport{Payload: []byte("report"), Result: traust.JobResult{Outcome: traust.OutcomeSucceeded}}, nil
			}), recordingPreparerFunc(func(ctx context.Context, _ traust.RunOrder) (RecordingRequest, error) {
				if stage == "template" {
					<-ctx.Done()
					return RecordingRequest{}, ctx.Err()
				}
				return assessmentRecordingTemplate(), nil
			}))
			runner.config.RequestTimeout = 100 * time.Millisecond
			if stage == "record" {
				server.recordGate = make(chan struct{})
			}
			order, err := runner.RunOnce(context.Background())
			events, _ := recordingEvents(server)
			if order == nil || order.Id == nil || !errors.Is(err, context.DeadlineExceeded) || slices.Contains(events, "complete") {
				t.Fatalf("request deadline lost or ambiguous request completed: order=%+v err=%v events=%v", order, err, events)
			}
			records := 0
			for _, event := range events {
				if event == "record" {
					records++
				}
			}
			if (stage == "template" && (executed || records != 0)) || (stage == "record" && (!executed || records != 1)) {
				t.Fatalf("deadline retried or advanced work: executed=%v records=%d events=%v", executed, records, events)
			}
		})
	}
}

func (server *queuedRecordingServer) preparedFixRef() traust.RepositoryRef {
	kind := string(traust.Branch)
	return traust.RepositoryRef{
		Id: &server.fixRefID, RepositoryId: server.repositoryID,
		Sha:  *server.verificationInput.WorkItem.FixCommit,
		Name: server.verificationInput.WorkItem.FixBranch, Kind: &kind,
	}
}

func verificationRecordingTemplate() RecordingRequest {
	return RecordingRequest{Verification: &traust.VerificationRecordBody{
		Provenance: traust.Provenance{Source: traust.SourceTypeTriageReport},
	}}
}

func TestQueuedVerificationReusesOnlyAnExactRegisteredPreparedFixRef(t *testing.T) {
	for _, tc := range []struct {
		name     string
		mutate   func(*traust.RepositoryRef)
		register bool
	}{
		{"already registered", func(*traust.RepositoryRef) {}, false},
		{"different commit", func(ref *traust.RepositoryRef) { ref.Sha = strings.Repeat("c", 40) }, true},
		{"different branch", func(ref *traust.RepositoryRef) { name := "other-branch"; ref.Name = &name }, true},
		{"different kind", func(ref *traust.RepositoryRef) { kind := string(traust.Tag); ref.Kind = &kind }, true},
		{"different repository", func(ref *traust.RepositoryRef) { ref.RepositoryId = uuid.New() }, true},
		{"missing identity", func(ref *traust.RepositoryRef) { ref.Id = nil }, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			runner, server := newQueuedRecordingRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
				return ProducedReport{Payload: []byte("verification report"), Result: traust.JobResult{Outcome: traust.OutcomeSucceeded}}, nil
			}), fixedRecordingTemplate(verificationRecordingTemplate()))
			ref := server.preparedFixRef()
			tc.mutate(&ref)
			server.refs = []traust.RepositoryRef{ref}
			order, err := runner.RunOnce(context.Background())
			events, completed := recordingEvents(server)
			if err != nil || order == nil || completed.Outcome != traust.OutcomeSucceeded || !slices.Contains(events, "find-ref") || slices.Contains(events, "register-ref") != tc.register || !slices.Contains(events, "record") {
				t.Fatalf("order=%+v err=%v events=%v result=%+v", order, err, events, completed)
			}
			if tc.register && (server.registered.Sha != *server.verificationInput.WorkItem.FixCommit ||
				server.registered.Name == nil || *server.registered.Name != *server.verificationInput.WorkItem.FixBranch ||
				server.registered.Kind == nil || *server.registered.Kind != traust.Branch) {
				t.Fatalf("registration used something other than prepared fix context: %+v", server.registered)
			}
		})
	}
}

func TestQueuedVerificationRefusesMissingAuthoritativeFixContext(t *testing.T) {
	for _, tc := range []struct {
		name   string
		mutate func(*traust.VerificationInput)
	}{
		{"missing fix commit", func(input *traust.VerificationInput) { input.WorkItem.FixCommit = nil }},
		{"empty fix commit", func(input *traust.VerificationInput) { empty := ""; input.WorkItem.FixCommit = &empty }},
		{"missing fix branch", func(input *traust.VerificationInput) { input.WorkItem.FixBranch = nil }},
		{"empty fix branch", func(input *traust.VerificationInput) { empty := ""; input.WorkItem.FixBranch = &empty }},
		{"missing repository", func(input *traust.VerificationInput) { input.Subject.RepositoryId = uuid.Nil }},
		{"different work item", func(input *traust.VerificationInput) { input.WorkItem.WorkItemId = uuid.New() }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			executed := false
			runner, server := newQueuedRecordingRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
				executed = true
				return ProducedReport{}, nil
			}), fixedRecordingTemplate(verificationRecordingTemplate()))
			tc.mutate(server.verificationInput)
			order, err := runner.RunOnce(context.Background())
			events, completed := recordingEvents(server)
			var refusal *PreparationRefusedError
			if order == nil || !errors.As(err, &refusal) || executed || slices.Contains(events, "find-ref") || slices.Contains(events, "register-ref") || slices.Contains(events, "record") ||
				completed.Outcome != traust.OutcomeRefused || completed.Reason == nil || *completed.Reason != traust.RunOrderReasonBadRequest {
				t.Fatalf("missing context was guessed from baseline/report: order=%+v err=%v executed=%v events=%v result=%+v", order, err, executed, events, completed)
			}
		})
	}
}

func TestQueuedRecordingDoesNotRecordARefusalOrEmptyFailure(t *testing.T) {
	for _, outcome := range []traust.Outcome{traust.OutcomeRefused, traust.OutcomeFailed} {
		t.Run(string(outcome), func(t *testing.T) {
			result := traust.JobResult{Outcome: outcome}
			var core *queuedRecordingServer
			runner, server := newQueuedRecordingRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
				core.event("execute")
				return ProducedReport{Result: result}, nil
			}), fixedRecordingTemplate(assessmentRecordingTemplate()))
			core = server
			order, err := runner.RunOnce(context.Background())
			events, completed := recordingEvents(server)
			if err != nil || order == nil || slices.Contains(events, "record") || !slices.Contains(events, "execute") || !slices.Contains(events, "complete") || completed.Outcome != outcome {
				t.Fatalf("terminal non-report outcome was recorded: order=%+v err=%v events=%v result=%+v", order, err, events, completed)
			}
		})
	}
}

func TestQueuedVerificationRefFailuresNeverExecuteRecordOrRetryRegistration(t *testing.T) {
	for _, tc := range []struct {
		name       string
		lookup     int
		register   int
		disconnect bool
		mismatch   bool
	}{
		{"lookup access refusal", 403, 0, false, false},
		{"lookup server error", 503, 0, false, false},
		{"registration access refusal", 0, 403, false, false},
		{"registration validation refusal", 0, 422, false, false},
		{"registration ambiguous server error", 0, 500, false, false},
		{"registration ambiguous transport error", 0, 0, true, false},
		{"registration mismatched receipt", 0, 0, false, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			executed := false
			runner, server := newQueuedRecordingRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
				executed = true
				return ProducedReport{}, nil
			}), fixedRecordingTemplate(verificationRecordingTemplate()))
			server.lookupStatus, server.registerStatus = tc.lookup, tc.register
			server.disconnectRegister, server.mismatchedRegistration = tc.disconnect, tc.mismatch
			order, err := runner.RunOnce(context.Background())
			events, completed := recordingEvents(server)
			wantClosed := tc.register == 422 || tc.mismatch
			if err == nil || order == nil || order.Id == nil || executed || slices.Contains(events, "record") || slices.Contains(events, "complete") != wantClosed {
				t.Fatalf("ref failure advanced lifecycle incorrectly: order=%+v err=%v executed=%v events=%v", order, err, executed, events)
			}
			if tc.mismatch {
				var refusal *PreparationRefusedError
				if !errors.As(err, &refusal) || order.State != traust.RunOrderStateRefused ||
					completed.Outcome != traust.OutcomeRefused || completed.Reason == nil || *completed.Reason != traust.RunOrderReasonBadRequest {
					t.Fatalf("successful mismatched ref receipt was left ambiguous: order=%+v err=%v result=%+v", order, err, completed)
				}
			} else if wantClosed && (order.State != traust.RunOrderStateFailed || completed.Outcome != traust.OutcomeFailed || completed.Reason != nil) {
				t.Fatalf("definitive ref error became a domain refusal or remained leased: order=%+v result=%+v", order, completed)
			}
			posts := 0
			for _, event := range events {
				if event == "register-ref" {
					posts++
				}
			}
			if (tc.lookup != 0 && posts != 0) || (tc.lookup == 0 && posts != 1) {
				t.Fatalf("registration retried or followed failed lookup: posts=%d events=%v", posts, events)
			}
			status := tc.lookup
			if tc.register != 0 {
				status = tc.register
			}
			if status != 0 {
				var problem *traust.ProblemError
				if !errors.As(err, &problem) || problem.StatusCode != status {
					t.Fatalf("typed ref failure was lost: %v", err)
				}
			}
		})
	}
}

func TestQueuedVerificationNeverRegistersAWorkerSelectedRefOrRewritesItsRejectedReport(t *testing.T) {
	original := []byte("{\"meta\":{\"patched_commit\":\"cccccccccccccccccccccccccccccccccccccccc\",\"patched_ref\":\"worker-selected-branch\"}}\\n")
	runner, server := newQueuedRecordingRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
		return ProducedReport{Payload: original, Result: traust.JobResult{Outcome: traust.OutcomeSucceeded}}, nil
	}), fixedRecordingTemplate(verificationRecordingTemplate()))
	server.recordStatus = http.StatusUnprocessableEntity
	order, err := runner.RunOnce(context.Background())
	var problem *traust.ProblemError
	events, completed := recordingEvents(server)
	if order == nil || !errors.As(err, &problem) || problem.StatusCode != http.StatusUnprocessableEntity ||
		server.registered.Sha != *server.verificationInput.WorkItem.FixCommit ||
		server.registered.Name == nil || *server.registered.Name != *server.verificationInput.WorkItem.FixBranch ||
		!bytes.Equal([]byte(server.recorded.Payload), original) || completed.Outcome != traust.OutcomeFailed {
		t.Fatalf("worker ref was trusted or refused bytes changed: order=%+v err=%v registered=%+v payload=%q result=%+v", order, err, server.registered, server.recorded.Payload, completed)
	}
	registrations := 0
	for _, event := range events {
		if event == "register-ref" {
			registrations++
		}
	}
	if registrations != 1 {
		t.Fatalf("worker ref triggered another registration: %v", events)
	}
}

func TestQueuedRecordingOrderBindingsCannotBeRewrittenByConsumerPointers(t *testing.T) {
	for _, mismatch := range []bool{false, true} {
		name := "original binding survives mutation"
		if mismatch {
			name = "mutated repository cannot bypass verification binding"
		}
		t.Run(name, func(t *testing.T) {
			executed := false
			repository, replacement := uuid.New(), uuid.New()
			runner, server := newQueuedRecordingRunner(t, executorFunc(func(_ context.Context, execution Execution) (ProducedReport, error) {
				executed = true
				if execution.Order.RepositoryId == nil || *execution.Order.RepositoryId != repository {
					return ProducedReport{}, errors.New("preparer changed executor repository binding")
				}
				*execution.Order.RepositoryId = replacement
				*execution.Order.WorkItemId = uuid.New()
				return ProducedReport{Payload: []byte("original report"), Result: traust.JobResult{Outcome: traust.OutcomeSucceeded}}, nil
			}), recordingPreparerFunc(func(_ context.Context, order traust.RunOrder) (RecordingRequest, error) {
				*order.RepositoryId = replacement
				*order.WorkItemId = uuid.New()
				return verificationRecordingTemplate(), nil
			}))
			server.repositoryID = repository
			server.verificationInput.Subject.RepositoryId = repository
			server.queue.order.RepositoryId = &repository
			server.queue.repository = &traust.Repository{Id: &repository, Url: "https://forge.test/example/repository"}
			runner.config.ForgeHosts = []string{"forge.test"}
			if mismatch {
				server.verificationInput.Subject.RepositoryId = replacement
			}
			order, err := runner.RunOnce(context.Background())
			events, completed := recordingEvents(server)
			if order == nil || order.RepositoryId == nil || *order.RepositoryId != repository ||
				order.WorkItemId == nil || *order.WorkItemId != server.verificationInput.WorkItem.WorkItemId {
				t.Fatalf("consumer mutated the immutable binding: order=%+v err=%v", order, err)
			}
			if mismatch {
				var refusal *PreparationRefusedError
				if !errors.As(err, &refusal) || executed || completed.Outcome != traust.OutcomeRefused ||
					slices.Contains(events, "find-ref") || slices.Contains(events, "register-ref") || slices.Contains(events, "record") {
					t.Fatalf("mutated repository bypassed verification binding: err=%v events=%v result=%+v", err, events, completed)
				}
			} else if err != nil || !executed || completed.Outcome != traust.OutcomeSucceeded ||
				server.registered.Sha != *server.verificationInput.WorkItem.FixCommit {
				t.Fatalf("original binding was not retained end to end: err=%v events=%v result=%+v", err, events, completed)
			}
		})
	}
}

func TestLeaseIsolatesEveryMutableOrderPointerAndReturnedSnapshot(t *testing.T) {
	id, started, expires := uuid.New(), time.Now().UTC(), time.Now().Add(time.Minute).UTC()
	order := traust.RunOrder{Id: &id, StartedAt: &started, LeaseExpiresAt: &expires, State: traust.RunOrderStateLeased}
	fields := reflect.ValueOf(&order).Elem()
	for i := range fields.NumField() {
		field := fields.Field(i)
		if field.Kind() != reflect.Pointer || !field.IsNil() {
			continue
		}
		field.Set(reflect.New(field.Type().Elem()))
		switch field.Type().Elem() {
		case reflect.TypeFor[uuid.UUID]():
			field.Elem().Set(reflect.ValueOf(uuid.New()))
		case reflect.TypeFor[time.Time]():
			field.Elem().Set(reflect.ValueOf(started))
		default:
			if field.Elem().Kind() != reflect.String {
				t.Fatalf("add a nonzero fixture for pointer field %s", fields.Type().Field(i).Name)
			}
			field.Elem().SetString("core-context")
		}
	}
	encoded, err := json.Marshal(order)
	if err != nil {
		t.Fatal(err)
	}
	var baseline traust.RunOrder
	if err := json.Unmarshal(encoded, &baseline); err != nil {
		t.Fatal(err)
	}
	l, err := hold(nil, &order, 30, time.Second, time.Second)
	if err != nil {
		t.Fatal(err)
	}
	completed := copyOrder(order)
	for _, exposed := range []*traust.RunOrder{&order, ptr(l.order()), l.snapshot(), ptr(l.latest(nil)), ptr(l.latest(&completed))} {
		fields := reflect.ValueOf(exposed).Elem()
		for i := range fields.NumField() {
			field := fields.Field(i)
			if field.Kind() == reflect.Pointer && !field.IsNil() {
				field.Elem().Set(reflect.Zero(field.Type().Elem()))
			}
		}
		if !reflect.DeepEqual(l.origin, baseline) || !reflect.DeepEqual(l.current, baseline) {
			t.Fatal("exposed order pointers rewrote held lease state")
		}
	}
	if !reflect.DeepEqual(completed, baseline) {
		t.Fatal("returned completed order shared mutable pointers with its source")
	}
}

func TestDispatchedPreparationFailuresCloseOnlyDefinitiveErrors(t *testing.T) {
	for _, tc := range []struct {
		name       string
		status     int
		lookup     int
		register   int
		disconnect bool
		malformed  bool
	}{
		{"preparation access error", 403, 0, 0, false, false},
		{"preparation validation error", 422, 0, 0, false, false},
		{"preparation server error", 503, 0, 0, false, false},
		{"preparation transport error", 0, 0, 0, true, false},
		{"malformed preparation", 0, 0, 0, false, true},
		{"fix ref lookup access error", 0, 403, 0, false, false},
		{"fix ref registration validation error", 0, 0, 422, false, false},
		{"fix ref registration server error", 0, 0, 503, false, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			executed := false
			runner, server := newQueuedRecordingRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
				executed = true
				return ProducedReport{}, nil
			}), nil)
			server.prepareStatus, server.lookupStatus, server.registerStatus = tc.status, tc.lookup, tc.register
			server.disconnectPrepare, server.malformedPrepare = tc.disconnect, tc.malformed
			run := RecordingRun{
				Client: runner.config.Client, Request: verificationRecordingTemplate(), Executor: runner.config.Executor,
				Start:             traust.StartRecordingRequest{LeaseSeconds: 30, Order: traust.OrderRequest{Asset: "ledger_audit"}},
				HeartbeatInterval: 5 * time.Millisecond, RequestTimeout: time.Second,
			}
			result, err := run.Run(context.Background())
			events, completed := recordingEvents(server)
			status := tc.status
			if tc.lookup != 0 {
				status = tc.lookup
			}
			if tc.register != 0 {
				status = tc.register
			}
			wantClosed := status == http.StatusUnprocessableEntity
			var refusal *PreparationRefusedError
			if result == nil || err == nil || executed || errors.As(err, &refusal) || slices.Contains(events, "record") ||
				slices.Contains(events, "complete") != wantClosed {
				t.Fatalf("dispatched error advanced or stranded lifecycle: result=%+v err=%v events=%v", result, err, events)
			}
			if status != 0 {
				var problem *traust.ProblemError
				if !errors.As(err, &problem) || problem.StatusCode != status {
					t.Fatalf("typed preparation error lost: %v", err)
				}
			}
			if wantClosed && (result.Order.State != traust.RunOrderStateFailed || completed.Outcome != traust.OutcomeFailed || completed.Reason != nil) {
				t.Fatalf("generic HTTP error did not close failed without a domain reason: %+v", result)
			}
			if !wantClosed && result.Order.State != traust.RunOrderStateLeased {
				t.Fatalf("ambiguous preparation did not retain reconciliation identity: %+v", result)
			}
			registrations := 0
			for _, event := range events {
				if event == "register-ref" {
					registrations++
				}
			}
			if tc.register != 0 && registrations != 1 {
				t.Fatalf("registration POST was retried: %v", events)
			}
		})
	}
}

func TestRunnerContinuesToSecondOrderAfterClosedRecordingFailuresAndRefusals(t *testing.T) {
	for _, failure := range []string{
		"invalid template", "consumer refusal", "HTTP-backed consumer refusal", "preparation 4xx",
		"fix ref 4xx", "mismatched ref receipt", "mismatched remediation input", "record 4xx",
		"executor crash", "invalid outcome", "reportless failure", "execution refusal",
	} {
		t.Run(failure, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			var attempt atomic.Int32
			var server *queuedRecordingServer
			runner, core := newQueuedRecordingRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
				if attempt.Load() == 1 {
					switch failure {
					case "executor crash":
						return ProducedReport{}, errors.New("executor crashed")
					case "invalid outcome":
						return ProducedReport{Result: traust.JobResult{Outcome: "invented"}}, nil
					case "reportless failure":
						return ProducedReport{Result: traust.JobResult{Outcome: traust.OutcomeFailed}}, nil
					case "execution refusal":
						return ProducedReport{Result: refused(traust.RunOrderReasonBudgetExhausted)}, nil
					}
				}
				return ProducedReport{Payload: []byte("second report"), Result: traust.JobResult{Outcome: traust.OutcomeSucceeded}}, nil
			}), recordingPreparerFunc(func(context.Context, traust.RunOrder) (RecordingRequest, error) {
				if attempt.Load() == 1 {
					switch failure {
					case "invalid template":
						return RecordingRequest{}, nil
					case "consumer refusal":
						return RecordingRequest{}, &PreparationRefusedError{
							Reason: traust.RunOrderReasonSourceUnavailable, Err: errors.New("secret-consumer-message"),
						}
					case "HTTP-backed consumer refusal":
						return RecordingRequest{}, &PreparationRefusedError{
							Reason: traust.RunOrderReasonSourceUnavailable,
							Err: &traust.ProblemError{StatusCode: http.StatusUnprocessableEntity, Body: "secret-response-body",
								Problem: &traust.Problem{Title: "secret-response-body", Detail: "secret-response-body"}},
						}
					case "preparation 4xx":
						return RecordingRequest{AssessmentID: server.assessmentID,
							Triage: &traust.TriageRecordBody{Provenance: traust.Provenance{Source: traust.SourceTypeTriageReport}}}, nil
					case "fix ref 4xx", "mismatched ref receipt":
						return verificationRecordingTemplate(), nil
					case "mismatched remediation input":
						return RecordingRequest{Remediation: &traust.RemediationRecordBody{
							Provenance: traust.Provenance{Source: traust.SourceTypeTriageReport},
						}}, nil
					}
				}
				return assessmentRecordingTemplate(), nil
			}))
			server = core
			server.problemDetail = "secret-response-body"
			switch failure {
			case "preparation 4xx":
				server.prepareStatus = http.StatusUnprocessableEntity
			case "fix ref 4xx":
				server.registerStatus = http.StatusUnprocessableEntity
			case "mismatched ref receipt":
				ref := server.preparedFixRef()
				ref.Kind = nil
				server.registrationReply = &ref
			case "mismatched remediation input":
				server.remediationInput.WorkItem.WorkItemId = uuid.New()
			case "record 4xx":
				server.recordStatus = http.StatusUnprocessableEntity
			}
			firstID, secondID := *server.queue.order.Id, uuid.New()
			var completed []traust.JobResult
			var completedIDs []uuid.UUID
			server.queue.onClaim = func() {
				switch attempt.Add(1) {
				case 2:
					completed = append(completed, server.queue.result)
					completedIDs = append(completedIDs, *server.queue.order.Id)
					expires := time.Now().Add(time.Minute)
					server.queue.order.Id, server.queue.order.LeaseExpiresAt = &secondID, &expires
					server.queue.order.State, server.queue.order.Reason = traust.RunOrderStateLeased, nil
					server.prepareStatus, server.registerStatus, server.recordStatus = 0, 0, 0
				case 3:
					completed = append(completed, server.queue.result)
					completedIDs = append(completedIDs, *server.queue.order.Id)
					server.queue.empty = true
					cancel()
				}
			}
			var logs bytes.Buffer
			runner.config.Logger = slog.New(slog.NewJSONHandler(&logs, nil))
			err := runner.Run(ctx)
			if !errors.Is(err, context.Canceled) || attempt.Load() != 3 || len(completed) != 2 ||
				!slices.Equal(completedIDs, []uuid.UUID{firstID, secondID}) || completed[1].Outcome != traust.OutcomeSucceeded {
				t.Fatalf("runner stopped before the second order completed: err=%v attempts=%d ids=%v results=%+v", err, attempt.Load(), completedIDs, completed)
			}
			var entry struct {
				State                traust.RunOrderState `json:"state"`
				HTTPStatus           int                  `json:"http_status"`
				PreparationCauseType string               `json:"preparation_cause_type"`
			}
			if err := json.Unmarshal(bytes.TrimSpace(logs.Bytes()), &entry); err != nil {
				t.Fatalf("expected one completed-order diagnostic: %q: %v", logs.String(), err)
			}
			if (entry.State != traust.RunOrderStateFailed && entry.State != traust.RunOrderStateRefused) ||
				strings.Contains(logs.String(), "secret-response-body") || strings.Contains(logs.String(), "secret-consumer-message") {
				t.Fatalf("missing safe terminal diagnostic or leaked HTTP content: %s", logs.String())
			}
			if failure == "consumer refusal" && entry.PreparationCauseType != "*errors.errorString" {
				t.Fatalf("consumer refusal cause type was not observable: %s", logs.String())
			}
			if (failure == "preparation 4xx" || failure == "HTTP-backed consumer refusal") && entry.HTTPStatus != http.StatusUnprocessableEntity {
				t.Fatalf("safe HTTP status was not observable: %s", logs.String())
			}
			if failure == "record 4xx" && (completed[0].Reason == nil || *completed[0].Reason != traust.RunOrderReasonSessionOutputInvalid) {
				t.Fatalf("rejected report did not retain its classifier: %+v", completed[0])
			}
		})
	}
}

func TestRunnerDoesNotClaimAnotherOrderAfterAmbiguousRecordingOrCompletion(t *testing.T) {
	for _, failure := range []string{"record server error", "record transport error", "malformed record receipt", "completion refusal"} {
		t.Run(failure, func(t *testing.T) {
			runner, server := newQueuedRecordingRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
				return ProducedReport{Payload: []byte("original report"), Result: traust.JobResult{Outcome: traust.OutcomeSucceeded}}, nil
			}), fixedRecordingTemplate(assessmentRecordingTemplate()))
			switch failure {
			case "record server error":
				server.recordStatus = http.StatusInternalServerError
			case "record transport error":
				server.disconnectRecord = true
			case "malformed record receipt":
				server.malformedRecord = true
			case "completion refusal":
				server.queue.completeStatus = http.StatusForbidden
			}
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()
			if err := runner.Run(ctx); err == nil {
				t.Fatal("ambiguous order failure was hidden")
			}
			events, _ := recordingEvents(server)
			claims, records := 0, 0
			for _, event := range events {
				if event == "claim" {
					claims++
				}
				if event == "record" {
					records++
				}
			}
			if claims != 1 || records != 1 {
				t.Fatalf("runner advanced or retried an ambiguous POST: %v", events)
			}
		})
	}
}

func TestConsumerPreparationHTTPFailuresCloseOnlyDefinitiveErrors(t *testing.T) {
	for _, status := range []int{http.StatusForbidden, http.StatusUnprocessableEntity, http.StatusServiceUnavailable} {
		t.Run(http.StatusText(status), func(t *testing.T) {
			cause := &traust.ProblemError{StatusCode: status, Body: "private diagnostic"}
			executed := false
			runner, server := newQueuedRecordingRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
				executed = true
				return ProducedReport{}, nil
			}), recordingPreparerFunc(func(context.Context, traust.RunOrder) (RecordingRequest, error) {
				return RecordingRequest{}, errors.Join(errors.New("consumer input read"), cause)
			}))
			order, err := runner.RunOnce(context.Background())
			events, result := recordingEvents(server)
			var refusal *PreparationRefusedError
			if order == nil || !errors.Is(err, cause) || errors.As(err, &refusal) || executed ||
				slices.Contains(events, "record") || slices.Contains(events, "complete") != (status == http.StatusUnprocessableEntity) {
				t.Fatalf("consumer HTTP error was lost or reclassified: order=%+v err=%v events=%v", order, err, events)
			}
			if status == http.StatusUnprocessableEntity && (order.State != traust.RunOrderStateFailed || result.Outcome != traust.OutcomeFailed || result.Reason != nil) {
				t.Fatalf("definitive consumer HTTP error did not close failed: order=%+v result=%+v", order, result)
			}
			if status != http.StatusUnprocessableEntity && order.State != traust.RunOrderStateLeased {
				t.Fatalf("ambiguous consumer read did not preserve held identity: %+v", order)
			}
		})
	}
}

func TestRunnerHostAuthenticationAndRateFailuresDoNotClaimSecondOrder(t *testing.T) {
	for _, status := range []int{http.StatusUnauthorized, http.StatusForbidden, http.StatusTooManyRequests} {
		for _, stage := range []string{"manifest", "repository", "consumer", "consumer refusal", "input", "lookup", "registration", "record"} {
			t.Run(http.StatusText(status)+"/"+stage, func(t *testing.T) {
				ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
				defer cancel()
				var attempt atomic.Int32
				var server *queuedRecordingServer
				privateCause := &traust.ProblemError{StatusCode: status, Body: "secret-host-diagnostic"}
				runner, core := newQueuedRecordingRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
					return ProducedReport{Payload: []byte("original report"), Result: traust.JobResult{Outcome: traust.OutcomeSucceeded}}, nil
				}), recordingPreparerFunc(func(context.Context, traust.RunOrder) (RecordingRequest, error) {
					if attempt.Load() != 1 {
						return assessmentRecordingTemplate(), nil
					}
					switch stage {
					case "consumer":
						return RecordingRequest{}, errors.Join(errors.New("consumer read failed"), privateCause)
					case "consumer refusal":
						return RecordingRequest{}, &PreparationRefusedError{
							Reason: traust.RunOrderReasonSourceUnavailable, Err: privateCause,
						}
					case "input":
						return RecordingRequest{AssessmentID: server.assessmentID,
							Triage: &traust.TriageRecordBody{Provenance: traust.Provenance{Source: traust.SourceTypeTriageReport}}}, nil
					case "lookup", "registration":
						return verificationRecordingTemplate(), nil
					}
					return assessmentRecordingTemplate(), nil
				}))
				server = core
				server.problemDetail = "secret-host-diagnostic"
				switch stage {
				case "manifest":
					server.queue.manifestStatus = status
				case "repository":
					server.queue.order.RepositoryId = &server.repositoryID
					server.queue.repositoryStatus = status
					runner.config.ForgeHosts = []string{"forge.test"}
				case "input":
					server.prepareStatus = status
				case "lookup":
					server.lookupStatus = status
				case "registration":
					server.registerStatus = status
				case "record":
					server.recordStatus = status
				}
				firstID, secondID := *server.queue.order.Id, uuid.New()
				server.queue.onClaim = func() {
					switch attempt.Add(1) {
					case 2:
						expires := time.Now().Add(time.Minute)
						server.queue.order.Id, server.queue.order.LeaseExpiresAt = &secondID, &expires
						server.queue.order.State, server.queue.order.Reason = traust.RunOrderStateLeased, nil
						server.queue.order.RepositoryId = nil
						server.queue.manifestStatus, server.queue.repositoryStatus = 0, 0
						server.prepareStatus, server.lookupStatus, server.registerStatus, server.recordStatus = 0, 0, 0, 0
					case 3:
						server.queue.empty = true
						cancel()
					}
				}
				var logs bytes.Buffer
				runner.config.Logger = slog.New(slog.NewJSONHandler(&logs, nil))
				err := runner.Run(ctx)
				var problem *traust.ProblemError
				events, _ := recordingEvents(server)
				if !errors.As(err, &problem) || problem.StatusCode != status || attempt.Load() != 1 ||
					slices.Contains(events, "complete") {
					t.Fatalf("host error closed/drained orders or lost typed cause: err=%v attempts=%d events=%v", err, attempt.Load(), events)
				}
				server.queue.mu.Lock()
				retained := copyOrder(server.queue.order)
				server.queue.mu.Unlock()
				if retained.Id == nil || *retained.Id != firstID || retained.State != traust.RunOrderStateLeased {
					t.Fatalf("host error lost first order's reconciliation identity: %+v", retained)
				}
				if strings.Contains(logs.String(), "secret-host-diagnostic") {
					t.Fatalf("host diagnostic leaked into ordinary logs: %s", logs.String())
				}
				records, registrations := 0, 0
				for _, event := range events {
					if event == "record" {
						records++
					}
					if event == "register-ref" {
						registrations++
					}
				}
				if records > 1 || registrations > 1 || (stage == "record" && records != 1) ||
					(stage == "registration" && registrations != 1) {
					t.Fatalf("host error retried or advanced a POST: %v", events)
				}
			})
		}
	}
}

func TestRemediationPreparationRefusesInconsistentWorkItemAndRepository(t *testing.T) {
	for _, mismatch := range []string{"work item", "repository"} {
		for _, dispatched := range []bool{false, true} {
			name := mismatch + "/queued"
			if dispatched {
				name = mismatch + "/dispatched"
			}
			t.Run(name, func(t *testing.T) {
				executed := false
				request := RecordingRequest{Remediation: &traust.RemediationRecordBody{
					Provenance: traust.Provenance{Source: traust.SourceTypeTriageReport},
				}}
				runner, server := newQueuedRecordingRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
					executed = true
					return ProducedReport{}, nil
				}), fixedRecordingTemplate(request))
				server.queue.order.RepositoryId = &server.repositoryID
				server.queue.repository = &traust.Repository{Id: &server.repositoryID, Url: "https://forge.test/example/repository"}
				runner.config.ForgeHosts = []string{"forge.test"}
				switch mismatch {
				case "work item":
					server.remediationInput.WorkItem.WorkItemId = uuid.New()
				case "repository":
					server.remediationInput.Subject.RepositoryId = uuid.New()
				}
				var order *traust.RunOrder
				var err error
				if dispatched {
					run := RecordingRun{
						Client: runner.config.Client, Request: request, Executor: runner.config.Executor,
						Start:             traust.StartRecordingRequest{LeaseSeconds: 30, Order: traust.OrderRequest{Asset: "ledger_audit"}},
						HeartbeatInterval: 5 * time.Millisecond, RequestTimeout: time.Second,
					}
					result, runErr := run.Run(context.Background())
					err = runErr
					if result != nil {
						order = &result.Order
					}
				} else {
					order, err = runner.RunOnce(context.Background())
				}
				var refusal *PreparationRefusedError
				events, result := recordingEvents(server)
				if !errors.As(err, &refusal) || refusal.Reason != traust.RunOrderReasonBadRequest ||
					order == nil || order.State != traust.RunOrderStateRefused || executed ||
					result.Outcome != traust.OutcomeRefused || slices.Contains(events, "record") ||
					!slices.Contains(events, "complete") {
					t.Fatalf("mismatched remediation input reached execution or remained leased: order=%+v err=%v events=%v result=%+v", order, err, events, result)
				}
			})
		}
	}
}

func TestSuccessfulMismatchedVerificationRefReceiptRefusesBeforeExecution(t *testing.T) {
	for _, mismatch := range []string{"commit", "kind"} {
		t.Run(mismatch, func(t *testing.T) {
			executed := false
			runner, server := newQueuedRecordingRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
				executed = true
				return ProducedReport{}, nil
			}), fixedRecordingTemplate(verificationRecordingTemplate()))
			ref := server.preparedFixRef()
			switch mismatch {
			case "commit":
				ref.Sha = strings.Repeat("d", 40)
			case "kind":
				ref.Kind = nil
			}
			server.registrationReply = &ref
			order, err := runner.RunOnce(context.Background())
			var refusal *PreparationRefusedError
			events, result := recordingEvents(server)
			if !errors.As(err, &refusal) || refusal.Reason != traust.RunOrderReasonBadRequest ||
				order == nil || order.State != traust.RunOrderStateRefused || executed ||
				result.Outcome != traust.OutcomeRefused || slices.Contains(events, "record") ||
				!slices.Contains(events, "complete") {
				t.Fatalf("successful mismatched ref receipt executed or remained leased: order=%+v err=%v events=%v result=%+v", order, err, events, result)
			}
			registrations := 0
			for _, event := range events {
				if event == "register-ref" {
					registrations++
				}
			}
			if registrations != 1 {
				t.Fatalf("successful mismatched registration was retried: %v", events)
			}
		})
	}
}

func TestReconcileFindingRefs(t *testing.T) {
	preparedA := uuid.MustParse("1e4a4a1c-636b-4f73-9ca1-3b62e639f30b")
	preparedB := uuid.MustParse("2f5b5b2d-747c-4084-adb2-4c73f740a41c")
	foreign := uuid.MustParse("3a6c6c3e-858d-5195-bec3-5d84a851b52d")
	qualified := func(report, id string) string {
		return `{"source_findings":[{"finding_ref":"` + report + `#` + id + `"}],"summary":{"status":"patched"}}`
	}
	for _, tc := range []struct {
		name     string
		report   string
		prepared []uuid.UUID
		want     string
	}{
		{
			name:     "the qualified spelling rewrites to the prepared bare id",
			report:   qualified("service-security-audit.json", preparedA.String()),
			prepared: []uuid.UUID{preparedA, preparedB},
			want:     `{"source_findings":[{"finding_ref":"` + preparedA.String() + `"}],"summary":{"status":"patched"}}`,
		},
		{
			name:     "a bare prepared id rides through untouched",
			report:   `{"source_findings":[{"finding_ref":"` + preparedA.String() + `"}]}`,
			prepared: []uuid.UUID{preparedA},
			want:     `{"source_findings":[{"finding_ref":"` + preparedA.String() + `"}]}`,
		},
		{
			name:     "a tail naming no prepared finding rides through for the door to refuse",
			report:   qualified("service-security-audit.json", foreign.String()),
			prepared: []uuid.UUID{preparedA},
			want:     qualified("service-security-audit.json", foreign.String()),
		},
		{
			name:     "an unparseable tail rides through for the door to refuse",
			report:   qualified("service-security-audit.json", "FIND-005"),
			prepared: []uuid.UUID{preparedA},
			want:     qualified("service-security-audit.json", "FIND-005"),
		},
		{
			name:     "a report without source_findings is untouched",
			report:   `{"finding_ref":"worker-label","text":"é"}`,
			prepared: []uuid.UUID{preparedA},
			want:     `{"finding_ref":"worker-label","text":"é"}`,
		},
		{
			name:     "a malformed report is untouched",
			report:   "malformed report\r\n finding_ref: worker-chosen\n",
			prepared: []uuid.UUID{preparedA},
			want:     "malformed report\r\n finding_ref: worker-chosen\n",
		},
		{
			name:     "an empty source_findings array is untouched",
			report:   `{"source_findings":[],"summary":{}}`,
			prepared: []uuid.UUID{preparedA},
			want:     `{"source_findings":[],"summary":{}}`,
		},
		{
			name:     "a non-string finding_ref is untouched",
			report:   `{"source_findings":[{"finding_ref":17}]}`,
			prepared: []uuid.UUID{preparedA},
			want:     `{"source_findings":[{"finding_ref":17}]}`,
		},
		{
			name:     "a source_finding without finding_ref is untouched",
			report:   `{"source_findings":[{"finding_ref":"` + preparedA.String() + `"},{"title":"no ref"}]}`,
			prepared: []uuid.UUID{preparedA},
			want:     `{"source_findings":[{"finding_ref":"` + preparedA.String() + `"},{"title":"no ref"}]}`,
		},
		{
			name:     "no prepared findings leaves the report untouched",
			report:   qualified("service-security-audit.json", preparedA.String()),
			prepared: nil,
			want:     qualified("service-security-audit.json", preparedA.String()),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := string(reconcileFindingRefs([]byte(tc.report), tc.prepared)); got != tc.want {
				t.Fatalf("reconcileFindingRefs(%q) = %q, want %q", tc.report, got, tc.want)
			}
		})
	}
}

func TestQueuedRemediationRecordingReconcilesQualifiedFindingRefs(t *testing.T) {
	preparedA := uuid.MustParse("1e4a4a1c-636b-4f73-9ca1-3b62e639f30b")
	preparedB := uuid.MustParse("2f5b5b2d-747c-4084-adb2-4c73f740a41c")
	foreign := uuid.MustParse("3a6c6c3e-858d-5195-bec3-5d84a851b52d")
	var server *queuedRecordingServer
	runner, core := newQueuedRecordingRunner(t, executorFunc(func(_ context.Context, execution Execution) (ProducedReport, error) {
		server.event("execute")
		report := `{"metadata":{"repository":"https://github.com/octo/service"},"source_findings":[` +
			`{"finding_ref":"service-security-audit.json#` + preparedA.String() + `","title":"runs as root","severity":"medium"},` +
			`{"finding_ref":"` + preparedB.String() + `","title":"hardcoded token","severity":"high"},` +
			`{"finding_ref":"other-report.json#` + foreign.String() + `","title":"foreign finding","severity":"low"}` +
			`],"fork":{"url":"https://github.com/octo-mirrors/sci-patch-service","fix_branch":"fix/authoritative"},"summary":{"status":"patched"}}`
		return ProducedReport{Payload: []byte(report), Result: traust.JobResult{Outcome: traust.OutcomeSucceeded}}, nil
	}), fixedRecordingTemplate(RecordingRequest{Remediation: &traust.RemediationRecordBody{
		Provenance: traust.Provenance{Source: traust.SourceTypeTriageReport},
	}}))
	server = core
	server.remediationInput.WorkItem.FindingIds = []uuid.UUID{preparedA, preparedB}
	order, err := runner.RunOnce(context.Background())
	if err != nil || order == nil || order.State != traust.RunOrderStateSucceeded {
		t.Fatalf("order=%+v err=%v", order, err)
	}
	events, _ := recordingEvents(server)
	if !slices.Contains(events, "prepare") || !slices.Contains(events, "record") || !slices.Contains(events, "complete") {
		t.Fatalf("the remediation recording did not run its lifecycle: %v", events)
	}
	var doc struct {
		SourceFindings []struct {
			FindingRef string `json:"finding_ref"`
		} `json:"source_findings"`
		Fork struct {
			URL string `json:"url"`
		} `json:"fork"`
	}
	if err := json.Unmarshal([]byte(server.recorded.Payload), &doc); err != nil {
		t.Fatalf("the recorded payload is not the reconciled report: %v %q", err, server.recorded.Payload)
	}
	if len(doc.SourceFindings) != 3 {
		t.Fatalf("the recorded payload lost findings: %q", server.recorded.Payload)
	}
	if doc.SourceFindings[0].FindingRef != preparedA.String() {
		t.Fatalf("the qualified prepared ref was not reconciled to the bare id: %q", doc.SourceFindings[0].FindingRef)
	}
	if doc.SourceFindings[1].FindingRef != preparedB.String() {
		t.Fatalf("the bare prepared ref was rewritten: %q", doc.SourceFindings[1].FindingRef)
	}
	if doc.SourceFindings[2].FindingRef != "other-report.json#"+foreign.String() {
		t.Fatalf("the foreign qualified ref was rewritten instead of left for the door: %q", doc.SourceFindings[2].FindingRef)
	}
	if doc.Fork.URL != "https://github.com/octo-mirrors/sci-patch-service" {
		t.Fatalf("the report's other bytes did not ride through: %q", server.recorded.Payload)
	}
}
