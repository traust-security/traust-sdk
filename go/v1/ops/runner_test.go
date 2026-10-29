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
	"github.com/traust-security/traust-sdk/go/v1/api"
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
	mu              sync.Mutex
	order           api.RunOrder
	pack            *Pack
	empty           bool
	wrongDigest     bool
	malformed       bool
	heartbeatStatus int
	claimStatus     int
	manifestStatus  int
	manifestDelay   time.Duration
	heartbeatCalls  int
	heartbeatFailAt int
	events          []string
	result          api.JobResult
	onClaim         func()
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
			_ = json.NewEncoder(w).Encode(api.Problem{Title: "claim unavailable", Status: server.claimStatus})
			return
		}
		var body api.ClaimRequest
		if err := json.NewDecoder(request.Body).Decode(&body); err != nil || body.Assets == nil || len(*body.Assets) != 1 || (*body.Assets)[0] != "ledger_audit" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if server.empty {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		_ = json.NewEncoder(w).Encode(server.order)
	case strings.HasSuffix(request.URL.Path, "/content"):
		server.events = append(server.events, "manifest")
		if server.manifestDelay > 0 {
			time.Sleep(server.manifestDelay)
		}
		if server.manifestStatus != 0 {
			w.WriteHeader(server.manifestStatus)
			_ = json.NewEncoder(w).Encode(api.Problem{Title: "manifest unavailable", Status: server.manifestStatus})
			return
		}
		if server.malformed {
			_, _ = w.Write([]byte(`{"artifact":"wrong"}`))
			return
		}
		digest := server.pack.digest
		if server.wrongDigest {
			digest = strings.Repeat("0", 64)
		}
		_ = json.NewEncoder(w).Encode(map[string]any{
			"artifact": "run-manifest",
			"config":   []map[string]string{{"name": "pack:security", "digest": digest}},
		})
	case strings.HasSuffix(request.URL.Path, "/heartbeat"):
		server.events = append(server.events, "heartbeat")
		server.heartbeatCalls++
		if server.heartbeatStatus != 0 && (server.heartbeatFailAt == 0 || server.heartbeatCalls >= server.heartbeatFailAt) {
			w.Header().Set("Content-Type", "application/problem+json")
			w.WriteHeader(server.heartbeatStatus)
			_ = json.NewEncoder(w).Encode(api.Problem{Title: "lease lost", Status: server.heartbeatStatus})
			return
		}
		expires := time.Now().Add(time.Minute)
		server.order.LeaseExpiresAt = &expires
		_ = json.NewEncoder(w).Encode(server.order)
	case strings.HasSuffix(request.URL.Path, "/complete"):
		server.events = append(server.events, "complete")
		var body api.CompleteRequest
		if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		server.result = body.Result
		server.order.State = api.RunOrderState(body.Result.Outcome)
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
	server := &queuedServer{pack: pack, order: api.RunOrder{
		Id: &id, ArtifactId: &artifactID, Asset: "ledger_audit", State: api.RunOrderStateLeased,
		StartedAt: &started, LeaseExpiresAt: &expires, RequestedBy: "human:operator",
	}}
	httpServer := httptest.NewServer(server)
	t.Cleanup(httpServer.Close)
	tokenFile := filepath.Join(t.TempDir(), "token")
	if err := os.WriteFile(tokenFile, []byte("test-token"), 0o600); err != nil {
		t.Fatal(err)
	}
	client, err := api.NewBoundClient(api.ClientConfig{BaseURL: httpServer.URL, Tokens: api.FileTokenSource{Path: tokenFile}})
	if err != nil {
		t.Fatal(err)
	}
	runner, err := NewRunner(RunnerConfig{
		Client: client, Pack: pack, Executor: executor, Assets: []string{"ledger_audit"},
		LeaseSeconds: 30, HeartbeatInterval: 5 * time.Millisecond, RequestTimeout: time.Second, PollInterval: time.Millisecond,
	})
	if err != nil {
		t.Fatal(err)
	}
	return runner, server
}

func (server *queuedServer) snapshot() ([]string, api.JobResult) {
	server.mu.Lock()
	defer server.mu.Unlock()
	return append([]string(nil), server.events...), server.result
}

func TestQueuedRunnerClaimsExecutesAndCompletesLedgerAudit(t *testing.T) {
	var argv []string
	runner, server := queuedRunner(t, executorFunc(func(_ context.Context, execution Execution) (ProducedReport, error) {
		argv = append([]string(nil), execution.Argv...)
		return ProducedReport{
			Payload: []byte("command output need not be JSON"),
			Result:  api.JobResult{Outcome: api.OutcomeDegraded, Degraded: &map[string]int{"unverified": 1}},
		}, nil
	}))
	completed, err := runner.RunOnce(context.Background())
	if err != nil || completed == nil || completed.State != api.RunOrderStateDegraded {
		t.Fatalf("completed=%+v err=%v", completed, err)
	}
	events, result := server.snapshot()
	if strings.Join(events, ",") != "claim,manifest,heartbeat,heartbeat,complete" || result.Outcome != api.OutcomeDegraded {
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
	if err != nil || order == nil || order.State != api.RunOrderStateRefused {
		t.Fatalf("order=%+v err=%v", order, err)
	}
	events, _ := server.snapshot()
	if strings.Join(events, ",") != "claim,manifest,heartbeat,complete" {
		t.Fatalf("events=%v", events)
	}
}

func TestQueuedRunnerRefusesUnknownClaimedAsset(t *testing.T) {
	runner, server := queuedRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
		t.Fatal("executor ran an unconfigured asset")
		return ProducedReport{}, nil
	}))
	server.order.Asset = "unknown"
	order, err := runner.RunOnce(context.Background())
	if err != nil || order == nil || order.State != api.RunOrderStateRefused {
		t.Fatalf("order=%+v err=%v", order, err)
	}
	events, result := server.snapshot()
	if strings.Join(events, ",") != "claim,heartbeat,complete" || result.Reason == nil || *result.Reason != "asset-not-in-pack" {
		t.Fatalf("events=%v result=%+v", events, result)
	}
}

func TestQueuedRunnerNeverStartsAfterManifestUsesLease(t *testing.T) {
	runner, server := queuedRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
		t.Fatal("executor started after the lease expired")
		return ProducedReport{}, nil
	}))
	expires := time.Now().Add(200 * time.Millisecond)
	server.order.LeaseExpiresAt = &expires
	server.manifestDelay = 300 * time.Millisecond
	order, err := runner.RunOnce(context.Background())
	var lost *LeaseLostError
	if order == nil || !errors.As(err, &lost) {
		t.Fatalf("order=%+v err=%v", order, err)
	}
	events, _ := server.snapshot()
	if strings.Join(events, ",") != "claim,manifest" {
		t.Fatalf("work continued after lease expiry: %v", events)
	}
}

func TestQueuedRunnerRefusesRepositoryBoundFleetOrder(t *testing.T) {
	runner, server := queuedRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
		t.Fatal("executor ran a repository-bound fleet order")
		return ProducedReport{}, nil
	}))
	repositoryID := uuid.New()
	server.order.RepositoryId = &repositoryID
	order, err := runner.RunOnce(context.Background())
	if err != nil || order == nil || order.State != api.RunOrderStateRefused {
		t.Fatalf("order=%+v err=%v", order, err)
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
	runner, server := queuedRunner(t, executorFunc(func(ctx context.Context, _ Execution) (ProducedReport, error) {
		<-ctx.Done()
		close(stopped)
		return ProducedReport{}, ctx.Err()
	}))
	server.heartbeatStatus = http.StatusConflict
	server.heartbeatFailAt = 2
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
	} {
		t.Run(tc.name, func(t *testing.T) {
			runner, server := queuedRunner(t, executorFunc(func(context.Context, Execution) (ProducedReport, error) {
				t.Fatal("executor ran after Core failed")
				return ProducedReport{}, nil
			}))
			tc.change(server)
			if _, err := runner.RunOnce(context.Background()); err == nil {
				t.Fatal("Core failure was ignored")
			}
			events, _ := server.snapshot()
			if len(events) == 0 || events[0] != "claim" || tc.manifestRead != (len(events) > 1) {
				t.Fatalf("unexpected calls after Core failure: %v", events)
			}
			for _, event := range events[1:] {
				if event != "manifest" {
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
			if err == nil || order == nil || order.State != api.RunOrderStateFailed {
				t.Fatalf("order=%+v err=%v", order, err)
			}
			events, result := server.snapshot()
			if strings.Join(events, ",") != "claim,manifest,heartbeat,heartbeat,complete" || result.Reason == nil || *result.Reason != "runner-crashed" {
				t.Fatalf("events=%v result=%+v", events, result)
			}
		})
	}
}
