//go:build core_rollback

package ops

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/traust-security/traust-sdk/go/v1/traust"
)

// Core's test_recording_runner_fails_the_order_after_intake_rollback owns the server and fault.
func TestCoreIntakeRollback(t *testing.T) {
	client, err := traust.NewClient(traust.Config{
		BaseURL: os.Getenv("TRAUST_E2E_URL"), Tokens: traust.FileTokenSource{Path: os.Getenv("TRAUST_E2E_RUNNER_TOKEN_FILE")},
	})
	if err != nil {
		t.Fatal(err)
	}
	repoID, err := uuid.Parse(os.Getenv("TRAUST_E2E_REPOSITORY_ID"))
	if err != nil {
		t.Fatal(err)
	}
	refID, err := uuid.Parse(os.Getenv("TRAUST_E2E_REF_ID"))
	if err != nil {
		t.Fatal(err)
	}
	payload, err := os.ReadFile(os.Getenv("TRAUST_E2E_REPORT_FILE"))
	if err != nil {
		t.Fatal(err)
	}
	run := RecordingRun{
		Client: client,
		Start: traust.StartRecordingRequest{
			Order: traust.OrderRequest{Asset: "audit", Priority: 1, RepositoryId: &repoID}, LeaseSeconds: 30,
		},
		Request: RecordingRequest{Assessment: &traust.AssessmentRecordRequest{
			RepositoryRefId: refID, AssessorVersion: "1.2.3",
			Provenance: traust.Provenance{Source: traust.SourceTypeVulnScanReport},
		}},
		Executor: executorFunc(func(context.Context, Execution) (ProducedReport, error) {
			return ProducedReport{Payload: payload, Result: traust.JobResult{Outcome: traust.OutcomeSucceeded}}, nil
		}),
	}
	result, err := run.Run(t.Context())
	problem := causeProblem(err)
	if problem == nil || problem.StatusCode != http.StatusUnprocessableEntity || problem.Problem == nil ||
		problem.Problem.Instance != "/v1/assessment-records" || problem.Problem.Detail != "intake failed after writing" {
		t.Fatalf("expected the post-intake refusal, got %v", err)
	}
	if result == nil || result.Order.State != traust.RunOrderStateFailed || result.Receipt.Assessment != nil ||
		result.Result.Outcome != traust.OutcomeFailed || result.Result.Reason == nil || *result.Result.Reason != traust.RunOrderReasonSessionOutputInvalid {
		t.Fatalf("unexpected recording result: %+v", result)
	}
}
