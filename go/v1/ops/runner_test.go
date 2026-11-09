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
	"slices"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/traust-security/traust-sdk/go/v1/traust"
)

const ledgerPack = `pack: security
assets:
  ledger_audit:
    run: [python3, -m, traust.cli, ledger, audit, --run-id, "{run_id}"]
    trigger: {every: 7d}
    runner: {labels: [deterministic]}
`

func TestLoadPackUsesPinnedSchema(t *testing.T) {
	pack, err := LoadPack([]byte(ledgerPack))
	if err != nil || pack.name != "security" || len(pack.digest) != 64 {
		t.Fatalf("pack=%+v err=%v", pack, err)
	}
	for _, tc := range []struct {
		name string
		raw  string
	}{
		{"unknown field", ledgerPack + "unexpected: true\n"},
		{"missing pack name", strings.Replace(ledgerPack, "pack: security\n", "", 1)},
		{"broken YAML", "pack: [\n"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := LoadPack([]byte(tc.raw)); err == nil {
				t.Fatal("invalid pack accepted")
			}
		})
	}
}

type queuedServer struct {
	mu               sync.Mutex
	order            traust.RunOrder
	pack             *Pack
	asset            string
	empty            bool
	wrongDigest      bool
	manifest         *traust.OrderManifest
	malformed        bool
	heartbeatStatus  int
	claimStatus      int
	manifestStatus   int
	completeStatus   int
	leaseLength      time.Duration
	repository       *traust.Repository
	repositoryStatus int
	manifestDelay    time.Duration
	manifestCanceled chan struct{}
	heartbeatCalls   int
	heartbeatFailAt  int
	events           []string
	result           traust.JobResult
	onClaim          func()
}

func (server *queuedServer) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	server.mu.Lock()
	defer server.mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	switch {
	case request.URL.Path == "/v1/orders/claim":
		server.events = append(server.events, "claim")
		if server.onClaim != nil {
			server.onClaim()
		}
		if server.claimStatus != 0 {
			w.WriteHeader(server.claimStatus)
			_ = json.NewEncoder(w).Encode(traust.Problem{Title: "claim unavailable", Status: server.claimStatus})
			return
		}
		var body traust.ClaimRequest
		if err := json.NewDecoder(request.Body).Decode(&body); err != nil || body.Assets == nil || len(*body.Assets) != 1 || (*body.Assets)[0] != server.asset {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if server.empty {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		_ = json.NewEncoder(w).Encode(server.order)
	case strings.HasSuffix(request.URL.Path, "/manifest"):
		server.events = append(server.events, "manifest")
		if server.manifestDelay > 0 {
			delay, canceled := server.manifestDelay, server.manifestCanceled
			server.mu.Unlock()
			timer := time.NewTimer(delay)
			select {
			case <-timer.C:
			case <-request.Context().Done():
				timer.Stop()
				if canceled != nil {
					select {
					case canceled <- struct{}{}:
					default:
					}
				}
				server.mu.Lock()
				return
			}
			server.mu.Lock()
		}
		if server.manifestStatus != 0 {
			w.WriteHeader(server.manifestStatus)
			_ = json.NewEncoder(w).Encode(traust.Problem{Title: "manifest unavailable", Status: server.manifestStatus})
			return
		}
		if server.malformed {
			_, _ = w.Write([]byte(`{"config":`))
			return
		}
		if server.manifest != nil {
			_ = json.NewEncoder(w).Encode(server.manifest)
			return
		}
		digest := server.pack.digest
		if server.wrongDigest {
			digest = strings.Repeat("0", 64)
		}
		_ = json.NewEncoder(w).Encode(traust.OrderManifest{"config": []traust.ConfigDigest{{Name: "pack:security", Digest: digest}}})
	case strings.HasPrefix(request.URL.Path, "/v1/repositories/"):
		server.events = append(server.events, "repository")
		if server.repositoryStatus != 0 {
			// Core answers every refusal with problem+json (RFC 9457), never FastAPI's
			// validation-error shape the generated client also knows for 422.
			w.Header().Set("Content-Type", "application/problem+json")
			w.WriteHeader(server.repositoryStatus)
			_ = json.NewEncoder(w).Encode(traust.Problem{Title: "repository unavailable", Status: server.repositoryStatus})
			return
		}
		if server.repository == nil {
			w.WriteHeader(http.StatusNotFound)
			_ = json.NewEncoder(w).Encode(traust.Problem{Title: "no such repository", Status: http.StatusNotFound})
			return
		}
		_ = json.NewEncoder(w).Encode(*server.repository)
	case strings.HasSuffix(request.URL.Path, "/heartbeat"):
		server.events = append(server.events, "heartbeat")
		server.heartbeatCalls++
		var body traust.LeaseRequest
		if err := json.NewDecoder(request.Body).Decode(&body); err != nil || body.LeaseSeconds <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if server.heartbeatStatus != 0 && (server.heartbeatFailAt == 0 || server.heartbeatCalls >= server.heartbeatFailAt) {
			w.Header().Set("Content-Type", "application/problem+json")
			w.WriteHeader(server.heartbeatStatus)
			_ = json.NewEncoder(w).Encode(traust.Problem{Title: "lease lost", Status: server.heartbeatStatus})
			return
		}
		expires := time.Now().Add(time.Duration(body.LeaseSeconds) * time.Second)
		if server.leaseLength > 0 {
			expires = time.Now().Add(server.leaseLength)
		}
		server.order.LeaseExpiresAt = &expires
		_ = json.NewEncoder(w).Encode(server.order)
	case strings.HasSuffix(request.URL.Path, "/complete"):
		server.events = append(server.events, "complete")
		var body traust.CompleteRequest
		if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		server.result = body.Result
		if server.completeStatus != 0 {
			w.Header().Set("Content-Type", "application/problem+json")
			w.WriteHeader(server.completeStatus)
			_ = json.NewEncoder(w).Encode(traust.Problem{Title: "completion rejected", Status: server.completeStatus})
			return
		}
		server.order.State = traust.RunOrderState(body.Result.Outcome)
		if body.Result.Reason != nil {
			value := string(*body.Result.Reason)
			server.order.Reason = &value
		}
		_ = json.NewEncoder(w).Encode(server.order)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func queuedRunner(t *testing.T, executor Executor) (*Runner, *queuedServer) {
	t.Helper()
	pack, err := LoadPack([]byte(ledgerPack))
	if err != nil {
		t.Fatal(err)
	}
	id, artifactID := uuid.New(), uuid.New()
	started, expires := time.Now(), time.Now().Add(time.Minute)
	server := &queuedServer{pack: pack, asset: "ledger_audit", order: traust.RunOrder{
		Id: &id, ArtifactId: &artifactID, Asset: "ledger_audit", State: traust.RunOrderStateLeased,
		StartedAt: &started, LeaseExpiresAt: &expires, RequestedBy: "human:operator",
	}}
	return serveRunner(t, executor, server, nil)
}

const vocabularyPack = `pack: security
assets:
  agentic_scan:
    run: [session, --run, "{run_id}", --worker, "{worker}", --role, "{role}", --rule, "{rule}", --repo, "{repo_url}"]
    model: {role: threat-model}
    trigger: {every: 7d}
`

// vocabularyRunner claims a repository-bound, routed order of an asset that names every
// parameter, held by worker runner-1, against a repository on github.com.
func vocabularyRunner(t *testing.T, executor Executor) (*Runner, *queuedServer) {
	t.Helper()
	pack, err := LoadPack([]byte(vocabularyPack))
	if err != nil {
		t.Fatal(err)
	}
	id, artifactID, repositoryID := uuid.New(), uuid.New(), uuid.New()
	started, expires := time.Now(), time.Now().Add(time.Minute)
	worker, rule := "runner-1", "rule-3+threat-model-pr"
	server := &queuedServer{pack: pack, asset: "agentic_scan", order: traust.RunOrder{
		Id: &id, ArtifactId: &artifactID, Asset: "agentic_scan", State: traust.RunOrderStateLeased,
		StartedAt: &started, LeaseExpiresAt: &expires, RequestedBy: "human:operator",
		LeaseHolder: &worker, RepositoryId: &repositoryID, Rule: &rule,
	}, repository: &traust.Repository{Id: &repositoryID, Url: "https://github.com/example/app"}}
	return serveRunner(t, executor, server, []string{"github.com"})
}

// serveRunner runs a client against server's fake core and a runner for its pack.
func serveRunner(t *testing.T, executor Executor, server *queuedServer, forgeHosts []string) (*Runner, *queuedServer) {
	t.Helper()
	httpServer := httptest.NewServer(server)
	t.Cleanup(httpServer.Close)
	tokenFile := filepath.Join(t.TempDir(), "token")
	if err := os.WriteFile(tokenFile, []byte("test-token"), 0o600); err != nil {
		t.Fatal(err)
	}
	client, err := traust.NewClient(traust.Config{BaseURL: httpServer.URL, Tokens: traust.FileTokenSource{Path: tokenFile}})
	if err != nil {
		t.Fatal(err)
	}
	runner, err := NewRunner(RunnerConfig{
		Client: client, Pack: server.pack, Executor: executor, Assets: []string{server.asset},
		Logger:       slog.New(slog.NewTextHandler(os.Stderr, nil)),
		ForgeHosts:   forgeHosts,
		LeaseSeconds: 30, HeartbeatInterval: 5 * time.Millisecond, RequestTimeout: time.Second, PollInterval: time.Millisecond,
	})
	if err != nil {
		t.Fatal(err)
	}
	return runner, server
}

func (server *queuedServer) snapshot() ([]string, traust.JobResult) {
	server.mu.Lock()
	defer server.mu.Unlock()
	return append([]string(nil), server.events...), server.result
}

func assertQueuedLifecycle(t *testing.T, events []string, minAfterManifest int) {
	t.Helper()
	manifest := slices.Index(events, "manifest")
	if len(events) < 4+minAfterManifest || events[0] != "claim" || manifest < 2 || len(events)-manifest-2 < minAfterManifest || events[len(events)-1] != "complete" {
		t.Fatalf("unexpected queued lifecycle: %v", events)
	}
	for index, event := range events[1 : len(events)-1] {
		if index+1 != manifest && event != "heartbeat" && event != "repository" {
			t.Fatalf("unexpected queued lifecycle: %v", events)
		}
	}
}

func TestQueuedRunnerClaimsExecutesAndCompletesLedgerAudit(t *testing.T) {
	var argv []string
	runner, server := queuedRunner(t, executorFunc(func(_ context.Context, execution Execution) (ProducedReport, error) {
		argv = append([]string(nil), execution.Argv...)
		return ProducedReport{
			Payload: []byte("command output need not be JSON"),
			Result:  traust.JobResult{Outcome: traust.OutcomeDegraded, Degraded: &map[string]int{"unverified": 1}},
		}, nil
	}))
	completed, err := runner.RunOnce(context.Background())
	if err != nil || completed == nil || completed.State != traust.RunOrderStateDegraded {
		t.Fatalf("completed=%+v err=%v", completed, err)
	}
	events, result := server.snapshot()
	assertQueuedLifecycle(t, events, 2)
	if result.Outcome != traust.OutcomeDegraded {
		t.Fatalf("events=%v result=%+v", events, result)
	}
	if len(argv) != 7 || argv[0] != "python3" || argv[3] != "ledger" || argv[6] != completed.Id.String() {
		t.Fatalf("resolved argv: %v", argv)
	}
}

func TestQueuedRunnerEmptyQueue(t *testing.T) {
	runner, server := queuedRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
		t.Fatal("executor ran for empty queue")
		return ProducedReport{}, nil
	}))
	server.empty = true
	order, err := runner.RunOnce(context.Background())
	if err != nil || order != nil {
		t.Fatalf("order=%+v err=%v", order, err)
	}
}

func TestQueuedRunnerPollsEmptyQueueUntilCanceled(t *testing.T) {
	runner, server := queuedRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
		t.Fatal("executor ran for empty queue")
		return ProducedReport{}, nil
	}))
	server.empty = true
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	server.onClaim = func() {
		if len(server.events) == 2 {
			cancel()
		}
	}
	if err := runner.Run(ctx); !errors.Is(err, context.Canceled) {
		t.Fatalf("runner stopped with %v", err)
	}
	events, _ := server.snapshot()
	if strings.Join(events, ",") != "claim,claim" {
		t.Fatalf("unexpected poll sequence: %v", events)
	}
}

func TestQueuedRunnerRefusesPackMismatchBeforeExecution(t *testing.T) {
	runner, server := queuedRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
		t.Fatal("executor ran with the wrong pack")
		return ProducedReport{}, nil
	}))
	server.wrongDigest = true
	order, err := runner.RunOnce(context.Background())
	if err != nil || order == nil || order.State != traust.RunOrderStateRefused {
		t.Fatalf("order=%+v err=%v", order, err)
	}
	events, _ := server.snapshot()
	assertQueuedLifecycle(t, events, 1)
}

func TestQueuedRunnerRefusesUnknownClaimedAsset(t *testing.T) {
	runner, server := queuedRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
		t.Fatal("executor ran an unconfigured asset")
		return ProducedReport{}, nil
	}))
	server.order.Asset = "unknown"
	order, err := runner.RunOnce(context.Background())
	if err != nil || order == nil || order.State != traust.RunOrderStateRefused {
		t.Fatalf("order=%+v err=%v", order, err)
	}
	events, result := server.snapshot()
	if strings.Join(events, ",") != "claim,heartbeat,complete" || result.Reason == nil || *result.Reason != "asset-not-in-pack" {
		t.Fatalf("events=%v result=%+v", events, result)
	}
}

func TestQueuedRunnerRenewsLeaseDuringSlowManifestRead(t *testing.T) {
	runner, server := queuedRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
		return ProducedReport{Result: traust.JobResult{Outcome: traust.OutcomeSucceeded}}, nil
	}))
	runner.config.LeaseSeconds = 1
	runner.config.HeartbeatInterval = 50 * time.Millisecond
	runner.config.RequestTimeout = 5 * time.Second
	expires := time.Now().Add(time.Second)
	server.order.LeaseExpiresAt = &expires
	server.manifestDelay = 1100 * time.Millisecond
	order, err := runner.RunOnce(context.Background())
	if err != nil || order == nil || order.State != traust.RunOrderStateSucceeded {
		t.Fatalf("order=%+v err=%v", order, err)
	}
	events, _ := server.snapshot()
	assertQueuedLifecycle(t, events, 3)
}

func TestQueuedRunnerCancelsManifestReadOnLeaseLoss(t *testing.T) {
	runner, server := queuedRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
		t.Fatal("executor started after lease loss")
		return ProducedReport{}, nil
	}))
	server.manifestDelay = 5 * time.Second
	server.manifestCanceled = make(chan struct{}, 1)
	server.heartbeatStatus = http.StatusConflict
	server.heartbeatFailAt = 2
	order, err := runner.RunOnce(context.Background())
	var lost *LeaseLostError
	if order == nil || !errors.As(err, &lost) {
		t.Fatalf("order=%+v err=%v", order, err)
	}
	select {
	case <-server.manifestCanceled:
	case <-time.After(time.Second):
		t.Fatal("manifest read was not canceled after lease loss")
	}
	events, _ := server.snapshot()
	if strings.Contains(strings.Join(events, ","), "complete") {
		t.Fatalf("completed after lease loss: %v", events)
	}
}

func TestQueuedRunnerRefusesARepositoryBoundOrderWithoutForgeHosts(t *testing.T) {
	runner, server := queuedRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
		t.Fatal("executor ran a repository-bound order without an allowlist to check it against")
		return ProducedReport{}, nil
	}))
	repositoryID := uuid.New()
	server.order.RepositoryId = &repositoryID
	server.repository = &traust.Repository{Id: &repositoryID, Url: "https://github.com/example/app"}
	order, err := runner.RunOnce(context.Background())
	if err != nil || order == nil || order.State != traust.RunOrderStateRefused {
		t.Fatalf("order=%+v err=%v", order, err)
	}
	if _, result := server.snapshot(); result.Reason == nil || *result.Reason != traust.RunOrderReasonBadRequest {
		t.Fatalf("result=%+v", result)
	}
}

func TestQueuedRunnerRunsARepositoryBoundOrderWhoseRepositoryIsAllowlisted(t *testing.T) {
	var argv []string
	runner, server := queuedRunner(t, executorFunc(func(_ context.Context, execution Execution) (ProducedReport, error) {
		argv = append([]string(nil), execution.Argv...)
		return ProducedReport{Result: traust.JobResult{Outcome: traust.OutcomeSucceeded}}, nil
	}))
	runner.config.ForgeHosts = []string{"github.com"}
	repositoryID := uuid.New()
	server.order.RepositoryId = &repositoryID
	server.repository = &traust.Repository{Id: &repositoryID, Url: "https://github.com/example/app.git"}
	order, err := runner.RunOnce(context.Background())
	if err != nil || order == nil || order.State != traust.RunOrderStateSucceeded {
		t.Fatalf("order=%+v err=%v", order, err)
	}
	if len(argv) != 7 || argv[6] != order.Id.String() {
		t.Fatalf("resolved argv: %v", argv)
	}
}

func TestVocabularyRunnerResolvesEveryParameterFromItsSource(t *testing.T) {
	var argv []string
	runner, server := vocabularyRunner(t, executorFunc(func(_ context.Context, execution Execution) (ProducedReport, error) {
		argv = append([]string(nil), execution.Argv...)
		return ProducedReport{Result: traust.JobResult{Outcome: traust.OutcomeSucceeded}}, nil
	}))
	order, err := runner.RunOnce(context.Background())
	if err != nil || order == nil || order.State != traust.RunOrderStateSucceeded {
		t.Fatalf("order=%+v err=%v", order, err)
	}
	want := []string{
		"session", "--run", order.Id.String(), "--worker", "runner-1", "--role", "threat-model",
		"--rule", "rule-3+threat-model-pr", "--repo", "https://github.com/example/app",
	}
	if !slices.Equal(argv, want) {
		t.Fatalf("resolved argv %v, want %v", argv, want)
	}
	events, _ := server.snapshot()
	assertQueuedLifecycle(t, events, 2)
}

func TestVocabularyRunnerRefusesAnOrderMissingAParameter(t *testing.T) {
	for _, tc := range []struct {
		name   string
		change func(*queuedServer)
	}{
		{"no repository", func(server *queuedServer) { server.order.RepositoryId = nil }},
		{"no rule", func(server *queuedServer) { server.order.Rule = nil }},
		{"no lease holder", func(server *queuedServer) { server.order.LeaseHolder = nil }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			runner, server := vocabularyRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
				t.Fatal("executor ran an order with an unresolved parameter")
				return ProducedReport{}, nil
			}))
			tc.change(server)
			order, err := runner.RunOnce(context.Background())
			if err != nil || order == nil || order.State != traust.RunOrderStateRefused {
				t.Fatalf("order=%+v err=%v", order, err)
			}
			if _, result := server.snapshot(); result.Reason == nil || *result.Reason != traust.RunOrderReasonBadRequest {
				t.Fatalf("result=%+v", result)
			}
		})
	}
}

func TestVocabularyRunnerRefusesARepositoryURLCoreWouldNotParse(t *testing.T) {
	for _, tc := range []struct {
		name string
		url  string
	}{
		{"not https", "http://github.com/example/app"},
		{"host off the allowlist", "https://gitlab.com/example/app"},
		{"credentials", "https://secret@github.com/example/app"},
		{"port", "https://github.com:443/example/app"},
		{"one path segment", "https://github.com/app"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			runner, server := vocabularyRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
				t.Fatal("executor ran an order with an unresolvable repository")
				return ProducedReport{}, nil
			}))
			server.repository.Url = tc.url
			order, err := runner.RunOnce(context.Background())
			if err != nil || order == nil || order.State != traust.RunOrderStateRefused {
				t.Fatalf("order=%+v err=%v", order, err)
			}
			if _, result := server.snapshot(); result.Reason == nil || *result.Reason != traust.RunOrderReasonBadRequest {
				t.Fatalf("result=%+v", result)
			}
		})
	}
}

func TestVocabularyRunnerClosesOnlyOrderSpecificRepositoryErrors(t *testing.T) {
	for _, tc := range []struct {
		name   string
		status int
		change func(*queuedServer)
	}{
		{"missing", http.StatusNotFound, func(server *queuedServer) { server.repository = nil }},
		{"gone", http.StatusGone, func(server *queuedServer) { server.repositoryStatus = http.StatusGone }},
		{"unauthorized", http.StatusUnauthorized, func(server *queuedServer) { server.repositoryStatus = http.StatusUnauthorized }},
		{"forbidden", http.StatusForbidden, func(server *queuedServer) { server.repositoryStatus = http.StatusForbidden }},
		{"invalid request", http.StatusUnprocessableEntity, func(server *queuedServer) { server.repositoryStatus = http.StatusUnprocessableEntity }},
		{"rate limited", http.StatusTooManyRequests, func(server *queuedServer) { server.repositoryStatus = http.StatusTooManyRequests }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			runner, server := vocabularyRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
				t.Fatal("executor ran for a repository the runner cannot resolve")
				return ProducedReport{}, nil
			}))
			tc.change(server)
			order, err := runner.RunOnce(context.Background())
			var problem *traust.ProblemError
			if !errors.As(err, &problem) || order == nil {
				t.Fatalf("order=%+v err=%v", order, err)
			}
			wantClosed := problem.StatusCode == http.StatusNotFound || problem.StatusCode == http.StatusGone ||
				problem.StatusCode == http.StatusUnprocessableEntity
			events, result := server.snapshot()
			if slices.Contains(events, "complete") != wantClosed {
				t.Fatalf("repository error closed incorrectly: events=%v result=%+v", events, result)
			}
			if wantClosed && (order.State != traust.RunOrderStateFailed || result.Outcome != traust.OutcomeFailed) {
				t.Fatalf("order-specific repository error was not closed failed: order=%+v result=%+v", order, result)
			}
			if wantClosed {
				// A definitively failed read keeps its diagnosis on the completed row: the
				// 404 names not_found, every other definitive refusal stays runner-crashed
				// (the read failed without core naming a reason). Either way the row is
				// attributable — never a bare failed.
				wantReason := traust.RunOrderReasonRunnerCrashed
				if tc.status == http.StatusNotFound {
					wantReason = traust.RunOrderReasonNotFound
				}
				if result.Reason == nil || *result.Reason != wantReason {
					t.Fatalf("order-specific repository error lost its reason: result=%+v want=%s", result, wantReason)
				}
			}
			if !wantClosed && order.State != traust.RunOrderStateLeased {
				t.Fatalf("host repository error lost reconciliation identity: %+v", order)
			}
		})
	}
}

func TestVocabularyRunnerStopsOnAFailedRepositoryRead(t *testing.T) {
	runner, server := vocabularyRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
		t.Fatal("executor ran after the repository read failed")
		return ProducedReport{}, nil
	}))
	server.repositoryStatus = http.StatusServiceUnavailable
	order, err := runner.RunOnce(context.Background())
	if err == nil || order == nil {
		t.Fatalf("order=%+v err=%v", order, err)
	}
	events, _ := server.snapshot()
	if slices.Contains(events, "complete") {
		t.Fatalf("completed after a failed repository read: %v", events)
	}
}

// TestRunnerLogsRefusedAtWarnAndFailedAtError pins the log-level split: a refused order is a
// routine terminal state the runner closes on its own, while a failed one names a problem a
// human must act on.
func TestRunnerLogsRefusedAtWarnAndFailedAtError(t *testing.T) {
	for _, tc := range []struct {
		name       string
		change     func(*queuedServer)
		wantState  traust.RunOrderState
		wantLevel  string
		wantReason string
	}{
		{"refused pack mismatch", func(server *queuedServer) { server.wrongDigest = true },
			traust.RunOrderStateRefused, "WARN", "bad_request"},
		{"failed repository read", func(server *queuedServer) { server.repository = nil },
			traust.RunOrderStateFailed, "ERROR", "not_found"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			runner, server := vocabularyRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
				t.Fatal("executor ran for an order that closes before execution")
				return ProducedReport{}, nil
			}))
			tc.change(server)
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			var claimed bool
			server.onClaim = func() {
				// The first claim returns the order under test; let the runner finish
				// and log it before the queue runs dry and the loop stops.
				if claimed {
					server.empty = true
					cancel()
				}
				claimed = true
			}
			var logs bytes.Buffer
			runner.config.Logger = slog.New(slog.NewJSONHandler(&logs, nil))
			if err := runner.Run(ctx); !errors.Is(err, context.Canceled) {
				t.Fatalf("runner stopped with %v", err)
			}
			var entry struct {
				Level  string               `json:"level"`
				State  traust.RunOrderState `json:"state"`
				Reason string               `json:"reason"`
			}
			if err := json.Unmarshal(bytes.TrimSpace(logs.Bytes()), &entry); err != nil {
				t.Fatalf("expected one completed-order diagnostic: %q: %v", logs.String(), err)
			}
			if entry.Level != tc.wantLevel || entry.State != tc.wantState || entry.Reason != tc.wantReason {
				t.Fatalf("terminal diagnostic logged wrong: level=%s state=%s reason=%s, want %s/%s/%s: %s",
					entry.Level, entry.State, entry.Reason, tc.wantLevel, tc.wantState, tc.wantReason, logs.String())
			}
		})
	}
}

func TestQueuedRunnerRejectsEmbeddedCommandParameter(t *testing.T) {
	runner, _ := queuedRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
		return ProducedReport{}, nil
	}))
	pack, err := LoadPack([]byte(strings.Replace(ledgerPack, "{run_id}", "run-{run_id}", 1)))
	if err != nil {
		t.Fatal(err)
	}
	config := runner.config
	config.Pack = pack
	if _, err := NewRunner(config); err == nil {
		t.Fatal("embedded command parameter accepted")
	}
}

func TestQueuedRunnerRejectsARoleParameterWithoutAModelRole(t *testing.T) {
	runner, _ := queuedRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
		return ProducedReport{}, nil
	}))
	pack, err := LoadPack([]byte(strings.Replace(ledgerPack, "{run_id}", "{role}", 1)))
	if err != nil {
		t.Fatal(err)
	}
	config := runner.config
	config.Pack = pack
	if _, err := NewRunner(config); err == nil {
		t.Fatal("a command naming {role} without a model role accepted")
	}
}

func TestQueuedRunnerRequiresForgeHostsForARepoURLCommand(t *testing.T) {
	runner, _ := queuedRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
		return ProducedReport{}, nil
	}))
	pack, err := LoadPack([]byte(strings.Replace(ledgerPack, "{run_id}", "{repo_url}", 1)))
	if err != nil {
		t.Fatal(err)
	}
	config := runner.config
	config.Pack = pack
	if _, err := NewRunner(config); err == nil {
		t.Fatal("a command naming {repo_url} without forge hosts accepted")
	}
	config.ForgeHosts = []string{"github.com"}
	if _, err := NewRunner(config); err != nil {
		t.Fatalf("a command naming {repo_url} with forge hosts rejected: %v", err)
	}
}

func TestQueuedRunnerRequiresForgeHostsForAPartitionedAsset(t *testing.T) {
	runner, _ := queuedRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
		return ProducedReport{}, nil
	}))
	// A partitioned asset whose command does not name {repo_url}: its orders are still
	// repository-bound, so resolve checks each one's repository against the forge hosts.
	pack, err := LoadPack([]byte(strings.Replace(ledgerPack, "trigger: {every: 7d}", "trigger: {every: 7d}\n    partitions: dynamic:repo", 1)))
	if err != nil {
		t.Fatal(err)
	}
	config := runner.config
	config.Pack = pack
	if _, err := NewRunner(config); err == nil {
		t.Fatal("a partitioned asset without forge hosts accepted")
	}
	config.ForgeHosts = []string{"github.com"}
	if _, err := NewRunner(config); err != nil {
		t.Fatalf("a partitioned asset with forge hosts rejected: %v", err)
	}
}

func TestHTTPSRepoURLMirrorsCoresValidation(t *testing.T) {
	allowed := []string{"GitHub.com"}
	for _, tc := range []struct {
		name string
		raw  string
		want string
	}{
		{"https", "https://github.com/example/app", "https://github.com/example/app"},
		{"uppercase host", "https://GitHub.COM/example/app", "https://github.com/example/app"},
		{"git suffix", "https://github.com/example/app.git", "https://github.com/example/app"},
		{"empty userinfo", "https://@github.com/example/app", "https://github.com/example/app"},
		{"colon-only userinfo", "https://:@github.com/example/app", "https://github.com/example/app"},
		{"zero port", "https://github.com:0/example/app", "https://github.com/example/app"},
		{"trailing slash", "https://github.com/example/app/", "https://github.com/example/app"},
		{"deep path", "https://github.com/group/example/app", "https://github.com/group/example/app"},
		{"http", "http://github.com/example/app", ""},
		{"no scheme", "github.com/example/app", ""},
		{"credentials", "https://user:pw@github.com/example/app", ""},
		{"bare username", "https://user@github.com/example/app", ""},
		{"bare password", "https://:pw@github.com/example/app", ""},
		{"port", "https://github.com:443/example/app", ""},
		{"query", "https://github.com/example/app?x=1", ""},
		{"fragment", "https://github.com/example/app#f", ""},
		{"no path", "https://github.com", ""},
		{"one path segment", "https://github.com/app", ""},
		{"dot segment", "https://github.com/../app", ""},
		{"parent segment", "https://github.com/../app", ""},
		{"dash-led segment", "https://github.com/-app/x", ""},
		{"space in segment", "https://github.com/ex ample/x", ""},
		{"percent escape", "https://github.com/ex%20ample/x", ""},
		{"non-ascii", "https://github.com/example/äpp", ""},
		{"padded", " https://github.com/example/app", ""},
		{"dash-led", "-https://github.com/example/app", ""},
		{"empty host", "https:///example/app", ""},
		{"ipv6 host", "https://[::1]/example/app", ""},
		{"host off the allowlist", "https://gitlab.com/example/app", ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := httpsRepoURL(tc.raw, allowed)
			if tc.want == "" {
				if err == nil {
					t.Fatalf("accepted %q as %q", tc.raw, got)
				}
				return
			}
			if err != nil || got != tc.want {
				t.Fatalf("got %q err=%v, want %q", got, err, tc.want)
			}
		})
	}
}

func TestQueuedRunnerRejectsInvalidConfiguration(t *testing.T) {
	runner, _ := queuedRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
		return ProducedReport{}, nil
	}))
	for _, tc := range []struct {
		name   string
		change func(*RunnerConfig)
	}{
		{"missing client", func(cfg *RunnerConfig) { cfg.Client = nil }},
		{"missing pack", func(cfg *RunnerConfig) { cfg.Pack = nil }},
		{"missing executor", func(cfg *RunnerConfig) { cfg.Executor = nil }},
		{"missing logger", func(cfg *RunnerConfig) { cfg.Logger = nil }},
		{"missing assets", func(cfg *RunnerConfig) { cfg.Assets = nil }},
		{"duplicate asset", func(cfg *RunnerConfig) { cfg.Assets = append(cfg.Assets, cfg.Assets[0]) }},
		{"unknown asset", func(cfg *RunnerConfig) { cfg.Assets = []string{"unknown"} }},
		{"zero lease", func(cfg *RunnerConfig) { cfg.LeaseSeconds = 0 }},
		{"oversized lease", func(cfg *RunnerConfig) { cfg.LeaseSeconds = 86401 }},
		{"slow heartbeat", func(cfg *RunnerConfig) { cfg.HeartbeatInterval = 30 * time.Second }},
		{"negative timeout", func(cfg *RunnerConfig) { cfg.RequestTimeout = -time.Second }},
		{"missing poll interval", func(cfg *RunnerConfig) { cfg.PollInterval = 0 }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			cfg := runner.config
			tc.change(&cfg)
			if _, err := NewRunner(cfg); err == nil {
				t.Fatal("invalid runner configuration accepted")
			}
		})
	}
}

func TestQueuedRunnerStopsOnLostLease(t *testing.T) {
	stopped := make(chan struct{})
	var server *queuedServer
	runner, mock := queuedRunner(t, executorFunc(func(ctx context.Context, _ Execution) (ProducedReport, error) {
		server.mu.Lock()
		server.heartbeatStatus = http.StatusConflict
		server.mu.Unlock()
		<-ctx.Done()
		close(stopped)
		return ProducedReport{}, ctx.Err()
	}))
	server = mock
	order, err := runner.RunOnce(context.Background())
	var lost *LeaseLostError
	if order == nil || !errors.As(err, &lost) {
		t.Fatalf("order=%+v err=%v", order, err)
	}
	select {
	case <-stopped:
	default:
		t.Fatal("executor was not stopped before returning")
	}
	events, _ := server.snapshot()
	if strings.Contains(strings.Join(events, ","), "complete") {
		t.Fatalf("completed after lease loss: %v", events)
	}
}

func TestQueuedRunnerRetainsOrderOnPreExecutionHeartbeatFailure(t *testing.T) {
	runner, server := queuedRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
		t.Fatal("executor started without a live lease")
		return ProducedReport{}, nil
	}))
	runner.config.HeartbeatInterval = 10 * time.Second
	server.heartbeatStatus = http.StatusConflict
	server.heartbeatFailAt = 2
	order, err := runner.RunOnce(context.Background())
	var lost *LeaseLostError
	if order == nil || !errors.As(err, &lost) || order.Id == nil {
		t.Fatalf("order=%+v err=%v", order, err)
	}
	events, _ := server.snapshot()
	if slices.Contains(events, "complete") {
		t.Fatalf("completed after lease loss: %v", events)
	}
}

func TestQueuedRunnerDoesNotExecuteMalformedManifest(t *testing.T) {
	runner, server := queuedRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
		t.Fatal("executor ran with malformed manifest")
		return ProducedReport{}, nil
	}))
	server.malformed = true
	_, err := runner.RunOnce(context.Background())
	if err == nil {
		t.Fatal("malformed manifest was accepted")
	}
	events, _ := server.snapshot()
	if strings.Contains(strings.Join(events, ","), "complete") {
		t.Fatalf("completed after malformed manifest: %v", events)
	}
}

func TestQueuedRunnerStopsOnCoreReadFailure(t *testing.T) {
	for _, tc := range []struct {
		name         string
		change       func(*queuedServer)
		manifestRead bool
	}{
		{"claim", func(server *queuedServer) { server.claimStatus = http.StatusServiceUnavailable }, false},
		{"manifest", func(server *queuedServer) { server.manifestStatus = http.StatusServiceUnavailable }, true},
		{"manifest without body", func(server *queuedServer) { server.manifestStatus = http.StatusNoContent }, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			runner, server := queuedRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
				t.Fatal("executor ran after Core failed")
				return ProducedReport{}, nil
			}))
			tc.change(server)
			_, err := runner.RunOnce(context.Background())
			if err == nil {
				t.Fatal("Core failure was ignored")
			}
			if server.manifestStatus == http.StatusServiceUnavailable {
				var problem *traust.ProblemError
				if !errors.As(err, &problem) || problem.StatusCode != http.StatusServiceUnavailable {
					t.Fatalf("manifest failure was not typed: %v", err)
				}
			}
			events, _ := server.snapshot()
			if len(events) == 0 || events[0] != "claim" || tc.manifestRead != slices.Contains(events, "manifest") || slices.Contains(events, "complete") {
				t.Fatalf("unexpected calls after Core failure: %v", events)
			}
			for _, event := range events[1:] {
				if (tc.manifestRead && event != "manifest" && event != "heartbeat") || (!tc.manifestRead && event != "claim") {
					t.Fatalf("runner continued after Core failure: %v", events)
				}
			}
		})
	}
}

func TestQueuedRunnerCompletesFailedAfterInvalidExecution(t *testing.T) {
	for _, tc := range []struct {
		name     string
		produced ProducedReport
		err      error
	}{
		{"missing result", ProducedReport{Payload: []byte("command output")}, nil},
		{"executor crash", ProducedReport{}, errors.New("process failed")},
	} {
		t.Run(tc.name, func(t *testing.T) {
			runner, server := queuedRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
				return tc.produced, tc.err
			}))
			order, err := runner.RunOnce(context.Background())
			if err == nil || order == nil || order.State != traust.RunOrderStateFailed {
				t.Fatalf("order=%+v err=%v", order, err)
			}
			events, result := server.snapshot()
			assertQueuedLifecycle(t, events, 2)
			if result.Reason == nil || *result.Reason != "runner-crashed" {
				t.Fatalf("events=%v result=%+v", events, result)
			}
		})
	}
}

func TestQueuedRunnerRefusesAnOrderNotPinnedToExactlyItsPack(t *testing.T) {
	pinned := func(digests ...string) *traust.OrderManifest {
		config := []traust.ConfigDigest{{Name: "pack:other", Digest: strings.Repeat("1", 64)}}
		for _, digest := range digests {
			config = append(config, traust.ConfigDigest{Name: "pack:security", Digest: digest})
		}
		return &traust.OrderManifest{"config": config}
	}
	for _, tc := range []struct {
		name     string
		manifest func(pack string) *traust.OrderManifest
	}{
		{"no config", func(string) *traust.OrderManifest { return &traust.OrderManifest{} }},
		{"another pack only", func(string) *traust.OrderManifest { return pinned() }},
		{"this pack twice", func(pack string) *traust.OrderManifest { return pinned(pack, pack) }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			runner, server := queuedRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
				t.Fatal("executor ran an order not pinned to its pack")
				return ProducedReport{}, nil
			}))
			server.manifest = tc.manifest(server.pack.digest)
			order, err := runner.RunOnce(context.Background())
			if err != nil || order == nil || order.State != traust.RunOrderStateRefused {
				t.Fatalf("order=%+v err=%v", order, err)
			}
			if _, result := server.snapshot(); result.Reason == nil || *result.Reason != traust.RunOrderReasonBadRequest {
				t.Fatalf("result=%+v", result)
			}
		})
	}
}

func TestQueuedRunnerClosesOnlyOrderSpecificManifestErrors(t *testing.T) {
	for _, status := range []int{http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusUnprocessableEntity, http.StatusTooManyRequests} {
		t.Run(http.StatusText(status), func(t *testing.T) {
			runner, server := queuedRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
				t.Error("executor ran after the manifest read failed")
				return ProducedReport{}, nil
			}))
			server.manifestStatus = status
			order, err := runner.RunOnce(context.Background())
			var problem *traust.ProblemError
			events, result := server.snapshot()
			if !errors.As(err, &problem) || problem.StatusCode != status || order == nil {
				t.Fatalf("manifest HTTP error lost: order=%+v err=%v", order, err)
			}
			wantClosed := status == http.StatusNotFound || status == http.StatusUnprocessableEntity
			if slices.Contains(events, "complete") != wantClosed {
				t.Fatalf("manifest error was reclassified or closed incorrectly: events=%v result=%+v", events, result)
			}
			if wantClosed {
				// The closed row keeps the read's diagnosis: not_found for a 404
				// manifest, runner-crashed for any other definitive refusal.
				wantReason := traust.RunOrderReasonRunnerCrashed
				if status == http.StatusNotFound {
					wantReason = traust.RunOrderReasonNotFound
				}
				if result.Reason == nil || *result.Reason != wantReason {
					t.Fatalf("order-specific manifest error lost its reason: result=%+v want=%s", result, wantReason)
				}
			}
			if wantClosed && (order.State != traust.RunOrderStateFailed || result.Outcome != traust.OutcomeFailed) {
				t.Fatalf("order-specific manifest error left leased: order=%+v result=%+v", order, result)
			}
			if !wantClosed && order.State != traust.RunOrderStateLeased {
				t.Fatalf("host manifest error lost reconciliation identity: %+v", order)
			}
		})
	}
}
