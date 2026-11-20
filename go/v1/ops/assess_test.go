package ops

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/google/uuid"

	"github.com/traust-security/traust-sdk/go/v1/traust"
	"github.com/traust-security/traust-sdk/go/v1/traust/traustest"
)

func TestAssessDefaultRefNeedsItsDependencies(t *testing.T) {
	core := traustest.New(t)
	for name, build := range map[string]func() (RecordingPreparer, error){
		"client": func() (RecordingPreparer, error) { return AssessDefaultRef(nil, traust.SourceTypeVulnScanReport, "v1") },
		"source": func() (RecordingPreparer, error) { return AssessDefaultRef(core.Client(), "", "v1") },
		"version": func() (RecordingPreparer, error) {
			return AssessDefaultRef(core.Client(), traust.SourceTypeVulnScanReport, "")
		},
	} {
		if _, err := build(); err == nil {
			t.Fatalf("no %s: built a preparer", name)
		}
	}
}

func TestAssessDefaultRefRecordsTheCommitCoreNames(t *testing.T) {
	repo, ref := uuid.New(), uuid.New()
	tests := []struct {
		name       string
		order      traust.RunOrder
		defaultRef http.Handler
		refused    traust.RunOrderReason
		fails      bool
	}{
		{name: "the repository's default ref", order: traust.RunOrder{RepositoryId: &repo},
			defaultRef: traustest.JSON(http.StatusOK, traust.RepositoryRef{Id: &ref, RepositoryId: repo, Sha: "4c7f0b9"})},
		{name: "an order without a repository", order: traust.RunOrder{}, refused: traust.RunOrderReasonBadRequest},
		{name: "a repository without a default ref", order: traust.RunOrder{RepositoryId: &repo},
			defaultRef: traustest.Problem(http.StatusNotFound, "no default ref"), refused: traust.RunOrderReasonNotFound},
		{name: "another repository's ref", order: traust.RunOrder{RepositoryId: &repo},
			defaultRef: traustest.JSON(http.StatusOK, traust.RepositoryRef{Id: &ref, RepositoryId: uuid.New(), Sha: "4c7f0b9"}), fails: true},
		{name: "core unavailable", order: traust.RunOrder{RepositoryId: &repo},
			defaultRef: traustest.Problem(http.StatusServiceUnavailable, "down"), fails: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			core := traustest.New(t)
			core.Handle("GET /v1/repositories/{id}", traustest.JSON(http.StatusOK, traust.Repository{Id: &repo, Url: "https://forge.example/org/app"}))
			if tt.defaultRef != nil {
				core.Handle("GET /v1/repositories/{id}/refs/current-default", tt.defaultRef)
			}
			preparer, err := AssessDefaultRef(core.Client(), traust.SourceTypeVulnScanReport, "scanner-1")
			if err != nil {
				t.Fatal(err)
			}
			request, err := preparer.PrepareRecording(t.Context(), tt.order)
			var refusal *PreparationRefusedError
			switch {
			case tt.refused != "":
				if !errors.As(err, &refusal) || refusal.Reason != tt.refused {
					t.Fatalf("got %v, want a %s refusal", err, tt.refused)
				}
			case tt.fails:
				if err == nil || errors.As(err, &refusal) {
					t.Fatalf("got %v, want an error that is not a refusal", err)
				}
			default:
				var subject traust.SubjectRef
				if err != nil || request.Assessment == nil || request.Assessment.RepositoryRefId != ref ||
					request.Assessment.Provenance.Source != traust.SourceTypeVulnScanReport || json.Unmarshal(request.Input, &subject) != nil ||
					subject.Sha != "4c7f0b9" || subject.RepositoryUrl != "https://forge.example/org/app" {
					t.Fatalf("request %+v err %v, want an assessment of the default ref with its subject", request, err)
				}
			}
		})
	}
}

func TestTriageLatestJudgesTheRepositorysLatestAssessment(t *testing.T) {
	repo, assessment := uuid.New(), uuid.New()
	tests := []struct {
		name    string
		order   traust.RunOrder
		latest  http.Handler
		refused traust.RunOrderReason
		fails   bool
	}{
		{name: "the latest assessment", order: traust.RunOrder{RepositoryId: &repo}, latest: traustest.JSON(http.StatusOK, traust.Assessment{Id: &assessment})},
		{name: "an order without a repository", order: traust.RunOrder{}, refused: traust.RunOrderReasonBadRequest},
		{name: "nothing recorded yet", order: traust.RunOrder{RepositoryId: &repo}, latest: traustest.Problem(http.StatusNotFound, "none"), refused: traust.RunOrderReasonNotFound},
		{name: "core unavailable", order: traust.RunOrder{RepositoryId: &repo}, latest: traustest.Problem(http.StatusServiceUnavailable, "down"), fails: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			core := traustest.New(t)
			if tt.latest != nil {
				core.Handle("GET /v1/repositories/{id}/assessments/latest", tt.latest)
			}
			preparer, err := TriageLatest(core.Client(), traust.AssessmentKindSecureCode1)
			if err != nil {
				t.Fatal(err)
			}
			request, err := preparer.PrepareRecording(t.Context(), tt.order)
			var refusal *PreparationRefusedError
			switch {
			case tt.refused != "":
				if !errors.As(err, &refusal) || refusal.Reason != tt.refused {
					t.Fatalf("got %v, want a %s refusal", err, tt.refused)
				}
			case tt.fails:
				if err == nil || errors.As(err, &refusal) {
					t.Fatalf("got %v, want an error that is not a refusal", err)
				}
			default:
				if err != nil || request.Triage == nil || request.AssessmentID != assessment || request.Triage.Provenance.Source != traust.SourceTypeTriageReport {
					t.Fatalf("request %+v err %v, want a triage of the latest assessment", request, err)
				}
				if q := core.Requests()[0].Query.Get("kind"); q != string(traust.AssessmentKindSecureCode1) {
					t.Fatalf("asked for kind %q", q)
				}
			}
		})
	}
	if _, err := TriageLatest(nil, traust.AssessmentKindSecureCode1); err == nil {
		t.Fatal("TriageLatest built without a client")
	}
}
