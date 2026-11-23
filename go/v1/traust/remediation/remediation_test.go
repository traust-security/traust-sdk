package remediation_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"

	"github.com/traust-security/traust-sdk/go/v1/traust"
	"github.com/traust-security/traust-sdk/go/v1/traust/remediation"
)

// The re-keyed read names the work item in the path, and an effort core
// refuses (here: none for that work item) reads as *ProblemError, never a
// zero-value effort.
func TestGetRemediationNamesTheWorkItemAndRefusesAbsence(t *testing.T) {
	var path string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path = r.URL.Path
		w.Header().Set("Content-Type", "application/problem+json")
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(traust.Problem{
			Type: "about:blank", Title: http.StatusText(http.StatusNotFound), Status: http.StatusNotFound,
			Detail: "no remediation effort names this work item", RequestId: "req-test",
		})
	}))
	defer server.Close()

	client, err := traust.NewClient(traust.Config{BaseURL: server.URL, Tokens: traust.StaticToken("t")})
	if err != nil {
		t.Fatal(err)
	}
	id := uuid.New()
	_, err = remediation.New(client).GetRemediation(context.Background(), id)
	var problemErr *traust.ProblemError
	if !errors.As(err, &problemErr) {
		t.Fatalf("got %v (%T), want *ProblemError", err, err)
	}
	if path != "/v1/remediations/"+id.String() {
		t.Fatalf("requested %s, want /v1/remediations/%s", path, id)
	}
	if problemErr.StatusCode != http.StatusNotFound {
		t.Fatalf("status %d, want 404", problemErr.StatusCode)
	}
}

func TestPrepareVerificationPreservesTypedRefusals(t *testing.T) {
	for _, tc := range []struct {
		name        string
		contentType string
		body        string
		problem     bool
	}{
		{
			name:        "core problem instead of validation errors",
			contentType: "application/problem+json",
			body:        `{"type":"about:blank","title":"Unprocessable Entity","status":422,"detail":"verification requires a closed work item","request_id":"req-verification"}`,
			problem:     true,
		},
		{
			name:        "malformed problem document",
			contentType: "application/problem+json",
			body:        `{"title":"Unprocessable Entity","detail":`,
		},
		{
			name:        "non-JSON refusal",
			contentType: "text/plain",
			body:        "verification refused\n",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", tc.contentType)
				w.WriteHeader(http.StatusUnprocessableEntity)
				_, _ = w.Write([]byte(tc.body))
			}))
			defer server.Close()

			client, err := traust.NewClient(traust.Config{BaseURL: server.URL, Tokens: traust.StaticToken("test-token")})
			if err != nil {
				t.Fatal(err)
			}
			got, err := remediation.New(client).PrepareVerification(context.Background(), uuid.New())
			var refusal *traust.ProblemError
			if !errors.As(err, &refusal) || got != nil {
				t.Fatalf("got=%v err=%v (%T), want nil and *ProblemError", got, err, err)
			}
			if refusal.StatusCode != http.StatusUnprocessableEntity || refusal.Body != tc.body {
				t.Fatalf("status=%d body=%q, want 422 and %q", refusal.StatusCode, refusal.Body, tc.body)
			}
			if !tc.problem {
				if refusal.Problem != nil {
					t.Fatalf("problem=%+v, want no decoded problem", refusal.Problem)
				}
				return
			}
			document := refusal.Problem
			if document == nil || document.Type != "about:blank" || document.Title != "Unprocessable Entity" ||
				document.Status != http.StatusUnprocessableEntity || document.Detail != "verification requires a closed work item" ||
				document.RequestId != "req-verification" {
				t.Fatalf("problem=%+v, want the original core refusal document", document)
			}
		})
	}
}

func TestMalformedRefusalCannotLeakBodyOrForgeLogLines(t *testing.T) {
	const body = "{invalid\nsensitive-refusal-marker\r\nFORGED EVENT\x1b[31m"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = w.Write([]byte(body))
	}))
	defer server.Close()
	client, err := traust.NewClient(traust.Config{BaseURL: server.URL, Tokens: traust.StaticToken("test-token")})
	if err != nil {
		t.Fatal(err)
	}
	_, err = remediation.New(client).PrepareVerification(context.Background(), uuid.New())
	var refusal *traust.ProblemError
	if !errors.As(err, &refusal) || refusal.StatusCode != http.StatusUnprocessableEntity || refusal.Body != body {
		t.Fatalf("original refusal is unavailable for explicit inspection: %v", err)
	}
	formatted := err.Error()
	if strings.Contains(formatted, "sensitive-refusal-marker") || strings.Contains(formatted, "FORGED EVENT") || strings.ContainsAny(formatted, "\r\n\x1b") {
		t.Fatalf("formatted refusal leaked untrusted response text: %q", formatted)
	}
}

// The generated client's response types decode core's problem media, whatever its parameters.
func TestWithResponseDecodesCoreProblemMediaAndValidationErrors(t *testing.T) {
	for _, contentType := range []string{"application/problem+json", "application/problem+json; charset=utf-8", "application/json"} {
		t.Run(contentType, func(t *testing.T) {
			code := traust.RoleDenied
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", contentType)
				w.WriteHeader(http.StatusUnprocessableEntity)
				_ = json.NewEncoder(w).Encode(traust.Problem{
					Type: "about:blank", Title: http.StatusText(http.StatusUnprocessableEntity), Status: http.StatusUnprocessableEntity,
					Detail: "typed refusal", RequestId: "req-with-response", Code: &code,
					Issues: &[]traust.DenyIssue{{Code: code, Detail: "typed issue"}},
				})
			}))
			defer server.Close()
			client, err := traust.NewClient(traust.Config{BaseURL: server.URL, Tokens: traust.StaticToken("test-token")})
			if err != nil {
				t.Fatal(err)
			}
			verification, err := client.Raw().PrepareVerificationWithResponse(context.Background(), uuid.New())
			if err != nil || verification == nil || verification.StatusCode() != http.StatusUnprocessableEntity || verification.JSON422 == nil {
				t.Fatalf("problem media did not populate typed 422: response=%+v err=%v", verification, err)
			}
			problem := verification.JSON422
			if problem.Detail != "typed refusal" || problem.RequestId != "req-with-response" ||
				problem.Code == nil || *problem.Code != code || problem.Issues == nil || len(*problem.Issues) != 1 {
				t.Fatalf("typed problem was decoded as an empty or validation-error body: %+v", problem)
			}
		})
	}
}
