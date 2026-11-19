//go:build core_reclaim

package ops

import (
	"maps"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/traust-security/traust-sdk/go/v1/traust"
)

// Core's test_sdk_claim_reclaim_recording seeds the queued order and checks the durable result.
func TestCoreClaimReclaim(t *testing.T) {
	client, err := traust.NewClient(traust.Config{
		BaseURL: os.Getenv("TRAUST_E2E_URL"), Tokens: traust.FileTokenSource{Path: os.Getenv("TRAUST_E2E_RUNNER_TOKEN_FILE")},
	})
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
		Request: RecordingRequest{Assessment: &traust.AssessmentRecordRequest{
			RepositoryRefId: refID, AssessorVersion: "1.2.3",
			Provenance: traust.Provenance{Source: traust.SourceTypeVulnScanReport},
		}},
	}
	claim := func(id *uuid.UUID, seconds int) *lease {
		t.Helper()
		assets := []string{"audit"}
		order, err := client.Claim(t.Context(), traust.ClaimRequest{Assets: &assets, OrderId: id, LeaseSeconds: seconds})
		if err != nil {
			t.Fatal(err)
		}
		l, err := hold(client, order, seconds, time.Second, 10*time.Second)
		if err != nil {
			t.Fatal(err)
		}
		return l
	}
	old := claim(nil, 8)
	first, err := run.record(t.Context(), old, payload, nil)
	if err != nil || first.Assessment == nil || len(first.Assessment.Intake.Findings) != 2 {
		t.Fatalf("first attempt receipt: %+v, %v", first, err)
	}
	// Simulate a lost response followed by worker death: neither heartbeat nor complete.
	time.Sleep(time.Until(*old.origin.LeaseExpiresAt) + 10*time.Millisecond)
	id := old.id()
	current := claim(&id, 60)
	if current.id() != id || !current.origin.StartedAt.After(*old.origin.StartedAt) {
		t.Fatal("expected a new claim of the original order")
	}
	refused := func(l *lease, data []byte) {
		t.Helper()
		receipt, err := run.record(t.Context(), l, data, nil)
		if problem := causeProblem(err); problem == nil || problem.StatusCode != http.StatusConflict || receipt.Assessment != nil {
			t.Fatalf("expected conflict without receipt, got %+v, %v", receipt, err)
		}
	}
	refused(old, payload)
	refused(current, append([]byte("\n"), payload...))
	repeated, err := run.record(t.Context(), current, payload, nil)
	if err != nil || repeated.Assessment == nil {
		t.Fatalf("current attempt replay: %+v, %v", repeated, err)
	}
	a, b := first.Assessment, repeated.Assessment
	if b.AssessmentId != a.AssessmentId || b.ArtifactId != a.ArtifactId || !maps.Equal(b.Intake.Findings, a.Intake.Findings) || (b.Intake.Created != nil && len(*b.Intake.Created) != 0) {
		t.Fatal("reclaim did not replay the original assessment, artifact and finding IDs")
	}
	if _, err := current.complete(t.Context(), traust.JobResult{Outcome: traust.OutcomeSucceeded}); err != nil {
		t.Fatal(err)
	}
}
