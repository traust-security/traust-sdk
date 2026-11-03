//go:build e2e && (darwin || linux)

package ops

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/traust-security/traust-sdk/go/v1/traust"
)

// Runs against traust-server built from core at CORE_REF; `make e2e` starts it.
func e2eClient(t *testing.T, tokens string) *traust.Client {
	t.Helper()
	url, path := os.Getenv("TRAUST_E2E_URL"), os.Getenv(tokens)
	if url == "" || path == "" {
		t.Fatalf("TRAUST_E2E_URL and %s are required; run make e2e", tokens)
	}
	client, err := traust.NewClient(traust.Config{BaseURL: url, Tokens: traust.FileTokenSource{Path: path}})
	if err != nil {
		t.Fatal(err)
	}
	return client
}

func TestE2ERunnerRunsASubmittedPackCommand(t *testing.T) {
	raw, err := os.ReadFile("testdata/local-command.yaml")
	if err != nil {
		t.Fatal(err)
	}
	pack, err := LoadPack(raw)
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	person := e2eClient(t, "TRAUST_E2E_TOKEN_FILE")
	want, err := person.Submit(ctx, traust.SubmitRequest{
		Order:  traust.OrderRequest{Asset: "command_result", Priority: 1},
		Config: &[]traust.ConfigDigest{pack.Pin()},
	})
	if err != nil {
		t.Fatal(err)
	}

	runner, err := NewRunner(RunnerConfig{
		Client: e2eClient(t, "TRAUST_E2E_RUNNER_TOKEN_FILE"), Pack: pack, Executor: ProcessExecutor{},
		Assets: []string{"command_result"}, LeaseSeconds: 30, PollInterval: time.Second,
	})
	if err != nil {
		t.Fatal(err)
	}
	order, err := runner.RunOnce(ctx)
	if err != nil || order == nil || *order.Id != *want.Id || order.State != traust.RunOrderStateSucceeded {
		t.Fatalf("order=%+v err=%v", order, err)
	}
	if idle, err := runner.RunOnce(ctx); err != nil || idle != nil {
		t.Fatalf("drained queue returned order=%+v err=%v", idle, err)
	}
}

// A repository-bound, routed order of an asset with a model role fills every parameter: the
// order's id and rule, the claiming worker, the pack's role, and the repository's URL read
// through core and checked against the runner's forge hosts.
func TestE2ERunnerResolvesEveryCommandParameter(t *testing.T) {
	raw, err := os.ReadFile("testdata/vocabulary-command.yaml")
	if err != nil {
		t.Fatal(err)
	}
	pack, err := LoadPack(raw)
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	person := e2eClient(t, "TRAUST_E2E_TOKEN_FILE")
	label := "Smoke"
	// The harness's e2eforge mock (go/internal/e2eforge, wired in go/Makefile's e2e target)
	// answers for this repository: core reads its visibility there at onboarding, and the
	// runner checks the host against the allowlist below — no live forge involved.
	repository := "https://e2e-forge.test/e2e/fixture"
	onboarded, err := person.Onboard(ctx, traust.OnboardingRequest{
		Product: "smoke-vocabulary", Label: &label,
		Repositories: &[]traust.RepositorySpec{{Url: repository}},
	})
	if err != nil {
		t.Fatal(err)
	}
	repositoryID := (*onboarded.RepositoryIds)[0]
	rule := "rule-e2e"
	want, err := person.Submit(ctx, traust.SubmitRequest{
		Order:  traust.OrderRequest{Asset: "command_params", Priority: 1, RepositoryId: &repositoryID, Rule: &rule},
		Config: &[]traust.ConfigDigest{pack.Pin()},
	})
	if err != nil {
		t.Fatal(err)
	}

	var argv []string
	runner, err := NewRunner(RunnerConfig{
		Client: e2eClient(t, "TRAUST_E2E_RUNNER_TOKEN_FILE"), Pack: pack,
		Executor: executorFunc(func(_ context.Context, execution Execution) (ProducedReport, error) {
			argv = append([]string(nil), execution.Argv...)
			return ProducedReport{Result: traust.JobResult{Outcome: traust.OutcomeSucceeded}}, nil
		}),
		Assets: []string{"command_params"}, ForgeHosts: []string{"e2e-forge.test"},
		LeaseSeconds: 30, PollInterval: time.Second,
	})
	if err != nil {
		t.Fatal(err)
	}
	order, err := runner.RunOnce(ctx)
	if err != nil || order == nil || *order.Id != *want.Id || order.State != traust.RunOrderStateSucceeded {
		t.Fatalf("order=%+v err=%v", order, err)
	}
	if len(argv) != 8 || argv[0] != "python3" || argv[3] != want.Id.String() || argv[4] == "" ||
		argv[5] != "smoke-role" || argv[6] != rule || argv[7] != repository {
		t.Fatalf("resolved argv: %v", argv)
	}
	if idle, err := runner.RunOnce(ctx); err != nil || idle != nil {
		t.Fatalf("drained queue returned order=%+v err=%v", idle, err)
	}
}
