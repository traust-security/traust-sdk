//go:build e2e && (darwin || linux)

package ops

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/traust-security/traust-sdk/go/v1/traust"
	"github.com/traust-security/traust-sdk/go/v1/types"
)

// The fixture starts at SCI's ProducedReport boundary, not at a scanner's stdout.
// Finding shapes: source-code-intelligence-worker@70c3277 tests/entrypoint-scan-normalize.sh;
// scanner_correlation: the worker's pinned traust@ed72049e secure-code-audit skill (opengrep).
// Both worker@70c3277 and @619a80c emit jq -c report lines; only the former repairs malformed fields.
func TestE2ESCIQueuedScanFeedsTriage(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	person := e2eClient(t, "TRAUST_E2E_TOKEN_FILE")
	worker := e2eClient(t, "TRAUST_E2E_RUNNER_TOKEN_FILE")
	raw, err := os.ReadFile("testdata/sci-scan-report.json")
	if err != nil {
		t.Fatal(err)
	}
	var line bytes.Buffer
	if err := json.Compact(&line, raw); err != nil {
		t.Fatal(err)
	}
	var report types.Report
	if err := json.Unmarshal(raw, &report); err != nil {
		t.Fatal(err)
	}
	const repository = "https://e2e-forge.test/e2e/fixture"
	sha := strings.Repeat("a", 40) // The e2eforge default branch, independent of what the report claims.
	if report.Metadata.Repository == nil || *report.Metadata.Repository != repository || report.Metadata.Commit == nil || *report.Metadata.Commit != sha {
		t.Fatal("scan fixture must identify the e2e forge checkout")
	}
	label := "SCI recording"
	onboarded, err := person.Onboard(ctx, traust.OnboardingRequest{
		Product: "sci-recording", Label: &label, Repositories: &[]traust.RepositorySpec{{Url: repository}},
	})
	if err != nil {
		t.Fatal(err)
	}
	if onboarded.RepositoryIds == nil || len(*onboarded.RepositoryIds) != 1 {
		t.Fatalf("onboarded repositories: %v", onboarded.RepositoryIds)
	}
	repo := (*onboarded.RepositoryIds)[0]
	ref, err := person.MarkDefaultBranch(ctx, repo, traust.DefaultRefNaming{Branch: "main", Sha: sha})
	if err != nil || ref == nil || ref.Id == nil {
		t.Fatalf("naming the fixture head: ref=%v err=%v", ref, err)
	}
	pack, err := LoadPack([]byte(`pack: sci
assets:
  sci_scan:
    run: [/usr/local/bin/entrypoint.sh]
    partitions: dynamic:repo
    runner: {labels: [llm]}
`))
	if err != nil {
		t.Fatal(err)
	}
	request := traust.SubmitRequest{
		Order:  traust.OrderRequest{Asset: "sci_scan", RepositoryId: &repo, Priority: 1},
		Config: &[]traust.ConfigDigest{pack.Pin()},
	}
	queued, err := person.Submit(ctx, request)
	if err != nil {
		t.Fatal(err)
	}
	e2eAdmit(t, *queued)
	// source-code-intelligence@ca42045f (also @d378b860): scan_input.go prepares the
	// subject, runner.go extracts the report line; the pod never holds the Core client.
	prepare := recordingPreparerFunc(func(ctx context.Context, order traust.RunOrder) (RecordingRequest, error) {
		if order.RepositoryId == nil || *order.RepositoryId != repo {
			return RecordingRequest{}, fmt.Errorf("unexpected scan repository: %v", order.RepositoryId)
		}
		head, err := worker.GetCurrentDefaultRef(ctx, repo)
		if err != nil {
			return RecordingRequest{}, err
		}
		if head.Id == nil || *head.Id != *ref.Id || head.Sha != sha {
			return RecordingRequest{}, fmt.Errorf("unexpected scan head: %v", head)
		}
		subject, err := json.Marshal(traust.SubjectRef{RepositoryId: repo, RepositoryRefId: *head.Id, RepositoryUrl: repository, Sha: head.Sha})
		return RecordingRequest{Assessment: &traust.AssessmentRecordRequest{
			RepositoryRefId: *head.Id, AssessorVersion: "fixture-worker",
			Provenance: traust.Provenance{Source: traust.SourceTypeVulnScanReport},
		}, Input: subject}, err
	})
	success := traust.JobResult{Outcome: traust.OutcomeSucceeded}
	produced := ProducedReport{Payload: line.Bytes(), Result: success}
	runner, err := NewRunner(RunnerConfig{
		Client: worker, Pack: pack, Assets: []string{"sci_scan"}, ForgeHosts: []string{"e2e-forge.test"},
		LeaseSeconds: 30, PollInterval: time.Second, Logger: slog.Default(),
		Recordings: map[string]RecordingPreparer{"sci_scan": prepare},
		Executor: executorFunc(func(_ context.Context, execution Execution) (ProducedReport, error) {
			var subject traust.SubjectRef
			if err := json.Unmarshal(execution.Input, &subject); err != nil {
				return ProducedReport{}, err
			}
			if subject.RepositoryId != repo || subject.RepositoryRefId != *ref.Id || subject.RepositoryUrl != repository || subject.Sha != sha {
				return ProducedReport{}, fmt.Errorf("unexpected scan checkout: %+v", subject)
			}
			return produced, nil
		}),
	})
	if err != nil {
		t.Fatal(err)
	}
	order, err := runner.RunOnce(ctx)
	if err != nil || order == nil || order.Id == nil || queued.Id == nil || *order.Id != *queued.Id || order.State != traust.RunOrderStateSucceeded {
		t.Fatalf("queued scan: order=%v err=%v", order, err)
	}
	assessments, err := person.ListOrderAssessments(ctx, *order.Id)
	if err != nil || len(assessments) != 1 || assessments[0].Id == nil || assessments[0].RepositoryRefId != *ref.Id {
		t.Fatalf("queued scan assessments: %v err=%v", assessments, err)
	}
	id := *assessments[0].Id
	artifacts, err := person.ListAssessmentArtifacts(ctx, id, nil)
	if err != nil || len(artifacts) != 1 || artifacts[0].Id == nil {
		t.Fatalf("scan artifacts: %v err=%v", artifacts, err)
	}
	stored, err := person.GetArtifactContent(ctx, *artifacts[0].Id)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(stored, line.Bytes()) {
		t.Fatal("stored report differs from the worker line (including the range, effort and scanner decisions)")
	}
	bindings, err := person.GetAssessmentIntake(ctx, id)
	if err != nil {
		t.Fatal(err)
	}
	triage, err := worker.PrepareTriage(ctx, id)
	if err != nil {
		t.Fatal(err)
	}
	if triage.Subject.AssessmentId == nil || *triage.Subject.AssessmentId != id || triage.Subject.RepositoryRefId != *ref.Id || triage.Subject.RepositoryId != repo || triage.Subject.Sha != sha {
		t.Fatalf("triage prepared a different subject: %+v", triage.Subject)
	}
	if len(triage.Items) != 2 || len(bindings.Findings) != 2 || len(report.Findings) != 2 {
		t.Fatalf("expected just the two reported findings, not the dismissed scanner fact: items=%d bindings=%v", len(triage.Items), bindings.Findings)
	}
	for _, f := range report.Findings {
		t.Run(f.Id, func(t *testing.T) {
			fid, ok := bindings.Findings[f.Id]
			index := slices.IndexFunc(triage.Items, func(item traust.TriageItem) bool { return item.Finding.Id != nil && *item.Finding.Id == fid })
			if !ok || index < 0 {
				t.Fatalf("missing triage item for report label %s", f.Id)
			}
			item := triage.Items[index]
			if f.Fingerprint == nil || item.Finding.Fingerprint != *f.Fingerprint || item.Finding.FingerprintAlgo != "v3" || item.Finding.Title != f.Title {
				t.Fatalf("worker and intake disagree on finding identity: %+v", item.Finding)
			}
			if item.Occurrence.AssessmentId != id || item.Occurrence.FindingId != fid || item.Occurrence.Severity != traust.ScannerSeverity(f.Severity) || item.Occurrence.Id == nil {
				t.Fatalf("unexpected triage occurrence: %+v", item.Occurrence)
			}
			if item.Locations == nil || len(*item.Locations) != 1 || len(f.Locations) != 1 {
				t.Fatalf("expected one location for %s", f.Id)
			}
			loc := (*item.Locations)[0]
			if loc.OccurrenceId != *item.Occurrence.Id || loc.Path != f.Locations[0].Path {
				t.Fatalf("wrong occurrence or path: %+v", loc)
			}
			// The report keeps the whole range; native report intake indexes its first line.
			if f.Locations[0].Lines != nil {
				if loc.LineStart == nil || *loc.LineStart != 104 {
					t.Fatalf("expected indexed first line 104, got %+v", loc)
				}
			} else if loc.LineStart != nil || loc.LineEnd != nil {
				t.Fatalf("repo:maintenance acquired a synthetic line: %+v", loc)
			}
		})
	}

	runAnother := func(t *testing.T, payload string, result traust.JobResult) (*traust.RunOrder, error) {
		t.Helper()
		produced = ProducedReport{Payload: []byte(payload), Result: result}
		queued, err := person.Submit(ctx, request)
		if err != nil {
			t.Fatal(err)
		}
		e2eAdmit(t, *queued)
		got, err := runner.RunOnce(ctx)
		if got == nil || got.Id == nil || queued.Id == nil || *got.Id != *queued.Id {
			t.Fatalf("wrong order: queued=%+v got=%+v err=%v", queued, got, err)
		}
		return got, err
	}
	checkNoNewAudit := func(t *testing.T, scan *traust.RunOrder) {
		t.Helper()
		left, err := person.ListOrderAssessments(ctx, *scan.Id)
		if err != nil || len(left) != 0 {
			t.Fatalf("unsuccessful audit left assessments=%+v err=%v", left, err)
		}
		// SCI selects by repository and kind, not by the last known good assessment ID.
		latest, err := worker.GetLatestRecordedAssessment(ctx, repo, &traust.GetLatestRecordedAssessmentParams{Kind: traust.AssessmentKindSecureCode1})
		if err != nil || latest == nil || latest.Id == nil || *latest.Id != id || latest.Status != "completed" {
			t.Fatalf("latest audit=%+v err=%v, want previous completed audit %s", latest, err, id)
		}
		next, err := worker.PrepareTriage(ctx, *latest.Id)
		if err != nil || next == nil || next.Subject.AssessmentId == nil || *next.Subject.AssessmentId != id || next.Subject.Sha != sha || len(next.Items) != len(triage.Items) {
			t.Fatalf("wrong next triage input: %+v err=%v", next, err)
		}
	}
	for _, result := range []traust.JobResult{
		{Outcome: traust.OutcomeFailed, Reason: ptr(traust.RunOrderReasonRunnerCrashed)},
		{Outcome: traust.OutcomeDegraded, Degraded: &map[string]int{"network-skipped": 1}},
	} {
		t.Run(string(result.Outcome)+" with valid report", func(t *testing.T) {
			got, err := runAnother(t, line.String(), result)
			if err != nil || got.State != traust.RunOrderState(result.Outcome) ||
				(result.Reason != nil && (got.Reason == nil || *got.Reason != string(*result.Reason))) {
				t.Fatalf("scan=%+v err=%v, want %+v", got, err, result)
			}
			checkNoNewAudit(t, got)
		})
	}
	t.Run("report commit is not the checkout", func(t *testing.T) {
		wrong := strings.Replace(line.String(), `"commit":"`+sha+`"`, `"commit":"`+strings.Repeat("b", 40)+`"`, 1)
		if wrong == line.String() {
			t.Fatal("fixture did not carry the expected commit")
		}
		got, err := runAnother(t, wrong, success)
		problem := causeProblem(err)
		if problem == nil || problem.StatusCode != 422 || got.State != traust.RunOrderStateFailed ||
			got.Reason == nil || *got.Reason != string(traust.RunOrderReasonSessionOutputInvalid) {
			t.Fatalf("mismatched commit: order=%+v err=%v", got, err)
		}
		checkNoNewAudit(t, got)
	})

	// The worker instructs audits to omit validation_status. A historical
	// not_verified value is readable, but is not a new validity decision.
	t.Run("historical validity refuses the order", func(t *testing.T) {
		const firstFinding = `"findings":[{`
		valid := line.String()
		if strings.Count(valid, firstFinding) != 1 {
			t.Fatal("expected one compact findings array in the accepted payload")
		}
		got, err := runAnother(t, strings.Replace(valid, firstFinding, firstFinding+`"validation_status":"not_verified",`, 1), success)
		problem := causeProblem(err)
		if problem == nil || problem.StatusCode != 422 || problem.Problem == nil ||
			!strings.Contains(problem.Problem.Detail, "validity") || !strings.Contains(problem.Problem.Detail, "not_verified") ||
			got.State != traust.RunOrderStateFailed ||
			got.Reason == nil || *got.Reason != string(traust.RunOrderReasonSessionOutputInvalid) {
			t.Fatalf("historical validity: order=%+v err=%v", got, err)
		}
		checkNoNewAudit(t, got)
	})
}
