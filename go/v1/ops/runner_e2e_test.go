//go:build e2e && (darwin || linux)

package ops

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"os"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/traust-security/traust-sdk/go/v1/enums"
	"github.com/traust-security/traust-sdk/go/v1/skills/skillstest"
	"github.com/traust-security/traust-sdk/go/v1/traust"
	"github.com/traust-security/traust-sdk/go/v1/types"
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
		Logger: slog.New(slog.NewTextHandler(os.Stderr, nil)),
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
		Logger: slog.New(slog.NewTextHandler(os.Stderr, nil)),
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

func TestE2ERecordingPreservesReportFields(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	person := e2eClient(t, "TRAUST_E2E_TOKEN_FILE")
	worker := e2eClient(t, "TRAUST_E2E_RUNNER_TOKEN_FILE")
	repository, branch, sha := "https://e2e-forge.test/e2e/fixture", "main", strings.Repeat("a", 40)
	label := "Recording integrity"
	onboarded, err := person.Onboard(ctx, traust.OnboardingRequest{
		Product: "recording-integrity", Label: &label,
		Repositories: &[]traust.RepositorySpec{{Url: repository}},
	})
	if err != nil {
		t.Fatal(err)
	}
	repo := (*onboarded.RepositoryIds)[0]
	kind := traust.Branch
	ref, err := person.RegisterRef(ctx, repo, traust.RegisterRefJSONRequestBody{
		Sha: sha, Name: &branch, Kind: &kind,
	})
	if err != nil {
		t.Fatal(err)
	}

	document := skillstest.FixtureReport()
	document.Metadata.Repository, document.Metadata.Commit = &repository, &sha
	id := "FIXTURE-aaaaaaa-001"
	weakness := enums.FindingKindWeakness
	document.Findings = []types.ReportFinding{{
		Id: id, Kind: &weakness, Severity: enums.SeverityMedium, Title: "Untrusted input to the model",
		Cwes: []string{"CWE-20"}, Locations: []types.ReportLocation{{Path: "fixture.py"}},
		Description:        "The fixture lets untrusted input into a model instruction.",
		Remediation:        "Keep user input separate from instructions.",
		OwaspLlmReferences: []enums.OwaspLlmCategory{enums.OwaspLlmCategoryLLM01},
	}}
	document.ExecutiveSummary.SeverityCounts.Medium = 1
	document.FindingsSummary[2].Count, document.FindingsSummary[2].FindingIds = 1, []string{id}
	document.RemediationRoadmap[0].Addresses = []string{id}
	payload, err := json.Marshal(document)
	if err != nil {
		t.Fatal(err)
	}
	run := RecordingRun{
		Client: worker,
		Start: traust.StartRecordingRequest{
			LeaseSeconds: 30,
			Order:        traust.OrderRequest{Asset: "audit", Priority: 1, RepositoryId: &repo},
		},
		Request: RecordingRequest{Assessment: &traust.AssessmentRecordRequest{
			RepositoryRefId: *ref.Id, AssessorVersion: "1.2.3",
			Provenance: traust.Provenance{Source: traust.SourceTypeVulnScanReport},
		}},
		Executor: executorFunc(func(context.Context, Execution) (ProducedReport, error) {
			return ProducedReport{Payload: payload, Result: traust.JobResult{Outcome: traust.OutcomeSucceeded}}, nil
		}),
	}
	result, err := run.Run(ctx)
	if err != nil {
		t.Fatal(err)
	}
	recorded := result.Receipt.Assessment
	if recorded == nil || result.Order.State != traust.RunOrderStateSucceeded {
		t.Fatalf("recording did not succeed: %+v", result)
	}
	raw, err := person.GetArtifactContent(ctx, recorded.ArtifactId)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(raw, payload) {
		t.Fatal("stored report differs from the SDK payload")
	}
	findingID, ok := recorded.Intake.Findings[id]
	if !ok {
		t.Fatalf("no native finding bound to report label %s", id)
	}
	finding, err := person.GetFinding(ctx, findingID)
	if err != nil {
		t.Fatal(err)
	}
	if finding.Labels == nil || !slices.ContainsFunc(*finding.Labels, func(label traust.FindingLabel) bool {
		return label.Taxonomy == "owasp-llm" && label.Label == "LLM01"
	}) {
		t.Fatal("finding lost its OWASP LLM label")
	}

	t.Run("adapter envelope is not an authored report", func(t *testing.T) {
		raw, err := os.ReadFile("../storage/testdata/adapter-result-rich.test.json")
		if err != nil {
			t.Fatal(err)
		}
		adapter, err := types.ParseAdapterResultArtifact(raw)
		if err != nil {
			t.Fatal(err)
		}
		run.Executor = executorFunc(func(context.Context, Execution) (ProducedReport, error) {
			return ProducedReport{Payload: adapter.Payload(), Result: traust.JobResult{Outcome: traust.OutcomeSucceeded}}, nil
		})
		rejected, err := run.Run(ctx)
		problem := causeProblem(err)
		if problem == nil || problem.StatusCode != 422 || rejected == nil || rejected.Order.Id == nil {
			t.Fatalf("expected a schema rejection from Core, got result=%+v err=%v", rejected, err)
		}
		order, err := person.GetOrder(ctx, *rejected.Order.Id)
		if err != nil || order.State != traust.RunOrderStateFailed ||
			order.Reason == nil || *order.Reason != string(traust.RunOrderReasonSessionOutputInvalid) {
			t.Fatalf("rejected record's stored order=%+v err=%v", order, err)
		}
		assessments, err := person.ListOrderAssessments(ctx, *order.Id)
		if err != nil || len(assessments) != 0 {
			t.Fatalf("rejected record left assessments=%+v err=%v", assessments, err)
		}
	})
}
