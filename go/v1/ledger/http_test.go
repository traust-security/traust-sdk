package ledger_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/traust-security/traust-sdk/go/v1/ledger"
	"github.com/traust-security/traust-sdk/go/v1/ledger/querytest"
)

func mustNoErr(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func TestHTTPClient_BatchSubmitLane(t *testing.T) {
	var gotPath, gotLayer string
	var body ledger.BatchSubmitInput

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		gotLayer = r.URL.Query().Get("layer_id")
		raw, _ := io.ReadAll(r.Body)
		mustNoErr(t, json.Unmarshal(raw, &body))
		_, _ = w.Write([]byte(`{"id":"batch-ok","status":"accepted"}`))
	}))
	defer srv.Close()

	client := ledger.NewHTTPClient(srv.URL, ledger.WithBearerToken("tok-123"))

	resp, err := client.BatchSubmit(context.Background(), "repo-a", ledger.BatchSubmitInput{
		SourceRef:  "findings/repo-a/triage.json",
		RecordedAt: "2026-01-16T00:00:00+00:00",
		Events: []map[string]interface{}{
			{"finding_ref": "FIND-001"},
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if gotPath != "/v1/ledger/layer/submit" || gotLayer != "repo-a" {
		t.Fatalf("expected batch submit path, got %s", gotPath)
	}
	if body.SourceRef != "findings/repo-a/triage.json" {
		t.Fatalf("expected source_ref in body, got %s", body.SourceRef)
	}
	if resp.ID != "batch-ok" {
		t.Fatalf("expected ID batch-ok, got %s", resp.ID)
	}
}

func TestHTTPClient_EventLane(t *testing.T) {
	var gotPath, gotAuth string
	var envelope map[string]json.RawMessage
	var event struct {
		LayerID    string `json:"layer_id"`
		FindingRef string `json:"finding_ref"`
		Verdict    string `json:"verdict"`
		RecordedAt string `json:"recorded_at"`
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		gotAuth = r.Header.Get("Authorization")
		body, _ := io.ReadAll(r.Body)
		mustNoErr(t, json.Unmarshal(body, &envelope))
		mustNoErr(t, json.Unmarshal(envelope["event"], &event))
		_, _ = w.Write([]byte(`{"id":"cs-1","status":"accepted"}`))
	}))
	defer srv.Close()

	client := ledger.NewHTTPClient(srv.URL, ledger.WithBearerToken("user-tok"))

	_, err := client.SubmitCountersign(context.Background(), ledger.CountersignInput{
		EventMeta: ledger.EventMeta{
			LayerID:    "repo-a",
			RecordedAt: "2026-01-16T00:00:00+00:00",
		},
		FindingRef:    "FIND-001",
		Verdict:       "true_positive",
		Justification: "Reviewed source and confirmed exploit path.",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if gotPath != "/v1/ledger/events" {
		t.Fatalf("expected /v1/ledger/events, got %s", gotPath)
	}
	if string(envelope["kind"]) != `"countersign"` {
		t.Fatalf("expected kind countersign in envelope, got %s", envelope["kind"])
	}
	if len(envelope) != 2 {
		t.Fatalf("envelope carries only kind and event, got %v", envelope)
	}
	if event.LayerID != "repo-a" {
		t.Fatalf("expected layer_id repo-a in event, got %s", event.LayerID)
	}
	if gotAuth != "Bearer user-tok" {
		t.Fatalf("expected Bearer user-tok, got %s", gotAuth)
	}
}

func TestHTTPClient_AllEventKindsRoute(t *testing.T) {
	var gotPath string
	var envelope struct {
		Kind string `json:"kind"`
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		body, _ := io.ReadAll(r.Body)
		mustNoErr(t, json.Unmarshal(body, &envelope))
		_, _ = w.Write([]byte(`{"id":"x","status":"accepted"}`))
	}))
	defer srv.Close()

	client := ledger.NewHTTPClient(srv.URL)
	meta := ledger.EventMeta{LayerID: "repo-a", RecordedAt: "2026-01-16T00:00:00+00:00"}

	tests := []struct {
		kind string
		call func() error
	}{
		{"countersign", func() error {
			_, err := client.SubmitCountersign(context.Background(), ledger.CountersignInput{
				EventMeta: meta, FindingRef: "FIND-001", Verdict: "true_positive",
				Justification: "Reviewed source and confirmed exploit path.",
			})
			return err
		}},
		{"severity", func() error {
			_, err := client.SubmitSeverity(context.Background(), ledger.SeverityInput{
				EventMeta: meta, FindingRef: "FIND-001", Severity: "high",
				Rationale: "Reviewed source and confirmed exploit path.",
			})
			return err
		}},
	}

	for _, tt := range tests {
		if err := tt.call(); err != nil {
			t.Fatalf("kind %s: unexpected error: %v", tt.kind, err)
		}
		if gotPath != "/v1/ledger/events" {
			t.Fatalf("%s: expected /v1/ledger/events, got %s", tt.kind, gotPath)
		}
		if envelope.Kind != tt.kind {
			t.Fatalf("expected kind %s in envelope, got %s", tt.kind, envelope.Kind)
		}
	}
}

func TestHTTPClient_ResolveAndFingerprint(t *testing.T) {
	var paths []string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paths = append(paths, r.URL.Path)
		switch r.URL.Path {
		case "/v1/ledger/layer/resolve":
			_, _ = w.Write([]byte(`{"resolved":true,"key":"k"}`))
		case "/v1/ledger/fingerprint":
			_, _ = w.Write([]byte(`{"findings":[],"stamped_count":0}`))
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer srv.Close()

	client := ledger.NewHTTPClient(srv.URL)

	if _, err := client.ResolveReviewItem(context.Background(), "repo-a", ledger.ResolveInput{
		Key: "k", Decision: "confirmed",
	}); err != nil {
		t.Fatalf("resolve: %v", err)
	}
	if _, err := client.ComputeFingerprints(context.Background(), ledger.FingerprintInput{
		Findings: []map[string]interface{}{{"id": "FIND-001"}},
	}); err != nil {
		t.Fatalf("fingerprint: %v", err)
	}

	if len(paths) != 2 {
		t.Fatalf("expected 2 paths, got %v", paths)
	}
}

func TestHTTPClient_SignLayer(t *testing.T) {
	var gotPath, gotAuth, gotMethod string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		gotMethod = r.Method
		gotAuth = r.Header.Get("Authorization")
		_, _ = w.Write([]byte(`{"status":"signed","method":"cosign","layer_id":"repo-a"}`))
	}))
	defer srv.Close()

	client := ledger.NewHTTPClient(srv.URL, ledger.WithBearerToken("sign-tok"))

	resp, err := client.SignLayer(context.Background(), "repo-a", ledger.SignOpts{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotPath != "/v1/ledger/layer/sign" {
		t.Fatalf("expected sign path, got %s", gotPath)
	}
	if gotMethod != "POST" {
		t.Fatalf("expected POST, got %s", gotMethod)
	}
	if gotAuth != "Bearer sign-tok" {
		t.Fatalf("expected Bearer sign-tok, got %s", gotAuth)
	}
	if resp.Status != "signed" || resp.LayerID != "repo-a" {
		t.Fatalf("unexpected response: %+v", resp)
	}
}

func TestHTTPClient_SignLayerRekor(t *testing.T) {
	var gotQuery string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotQuery = r.URL.RawQuery
		_, _ = w.Write([]byte(`{"status":"signed","method":"cosign","layer_id":"repo-a"}`))
	}))
	defer srv.Close()

	client := ledger.NewHTTPClient(srv.URL)
	_, err := client.SignLayer(context.Background(), "repo-a", ledger.SignOpts{Rekor: true})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotQuery != "layer_id=repo-a&rekor=true" {
		t.Fatalf("expected rekor=true query, got %q", gotQuery)
	}
}

func TestHTTPClient_Whoami(t *testing.T) {
	var gotPath, gotAuth string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		gotAuth = r.Header.Get("Authorization")
		_, _ = w.Write([]byte(`{"kind":"human","identity":"alice","identity_verified":true}`))
	}))
	defer srv.Close()

	client := ledger.NewHTTPClient(srv.URL, ledger.WithBearerToken("tok-123"))
	actor, err := client.Whoami(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotPath != "/v1/ledger/whoami" {
		t.Fatalf("expected /v1/ledger/whoami, got %s", gotPath)
	}
	if gotAuth != "Bearer tok-123" {
		t.Fatalf("expected Bearer tok-123, got %s", gotAuth)
	}
	if actor.Identity == nil || *actor.Identity != "alice" {
		t.Fatalf("expected identity alice, got %v", actor.Identity)
	}
	if actor.IdentityVerified == nil || !*actor.IdentityVerified {
		t.Fatalf("expected identity_verified true, got %v", actor.IdentityVerified)
	}
}

func TestHTTPClient_StatusError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(422)
		_, _ = w.Write([]byte(`{"error":"validation failed"}`))
	}))
	defer srv.Close()

	client := ledger.NewHTTPClient(srv.URL)
	_, err := client.BatchSubmit(context.Background(), "repo-a", ledger.BatchSubmitInput{
		SourceRef: "ref", RecordedAt: "2026-01-16T00:00:00+00:00",
	})
	if err == nil {
		t.Fatal("expected error for 422")
	}

	var se *ledger.StatusError
	if !errors.As(err, &se) {
		t.Fatalf("expected *StatusError, got %T: %v", err, err)
	}
	if se.StatusCode != 422 {
		t.Fatalf("expected status 422, got %d", se.StatusCode)
	}
}

func TestHTTPClient_KindInEnvelopeNotHeader(t *testing.T) {
	var gotKindHeader string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotKindHeader = r.Header.Get("X-Ingest-Kind")
		_, _ = w.Write([]byte(`{"id":"x","status":"ok"}`))
	}))
	defer srv.Close()

	client := ledger.NewHTTPClient(srv.URL)
	_, err := client.SubmitCountersign(context.Background(), ledger.CountersignInput{
		EventMeta:  ledger.EventMeta{LayerID: "repo-a", RecordedAt: "2026-01-16T00:00:00+00:00"},
		FindingRef: "FIND-001", Verdict: "true_positive",
	})
	mustNoErr(t, err)

	if gotKindHeader != "" {
		t.Fatal("kind should be in body envelope, not in headers")
	}
}

// TestHTTPClient_OpaqueLayerIDRoundTrip proves every layer-addressed call carries a
// migrated corpus ID (':' and '/') intact as the layer_id query value.
func TestHTTPClient_OpaqueLayerIDRoundTrip(t *testing.T) {
	const layerID = "corpus:layer:org/repo__main/repo__main"
	type seen struct{ method, path, layerID string }
	var got []seen
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got = append(got, seen{r.Method, r.URL.Path, r.URL.Query().Get("layer_id")})
		_, _ = w.Write([]byte(`{}`))
	}))
	defer srv.Close()

	client := ledger.NewHTTPClient(srv.URL)
	ctx := context.Background()
	calls := []struct {
		name   string
		method string
		path   string
		call   func() error
	}{
		{"GetLayer", http.MethodGet, "/v1/ledger/layer", func() error {
			_, err := client.GetLayer(ctx, layerID)
			return err
		}},
		{"GetFindings", http.MethodGet, "/v1/ledger/layer/findings", func() error {
			_, err := client.GetFindings(ctx, layerID)
			return err
		}},
		{"ListEvents", http.MethodGet, "/v1/ledger/layer/events", func() error {
			_, err := client.ListEvents(ctx, layerID, ledger.ListEventsOpts{Limit: 5})
			return err
		}},
		{"VerifyLayer", http.MethodGet, "/v1/ledger/layer/verify", func() error {
			_, err := client.VerifyLayer(ctx, layerID, ledger.VerifyOpts{CheckSignatures: true})
			return err
		}},
		{"BatchSubmit", http.MethodPost, "/v1/ledger/layer/submit", func() error {
			_, err := client.BatchSubmit(ctx, layerID, ledger.BatchSubmitInput{})
			return err
		}},
		{"ResolveReviewItem", http.MethodPost, "/v1/ledger/layer/resolve", func() error {
			_, err := client.ResolveReviewItem(ctx, layerID, ledger.ResolveInput{Key: "k"})
			return err
		}},
		{"StampEventIdentities", http.MethodPost, "/v1/ledger/layer/stamp", func() error {
			_, err := client.StampEventIdentities(ctx, layerID, ledger.StampInput{})
			return err
		}},
		{"SignLayer", http.MethodPost, "/v1/ledger/layer/sign", func() error {
			_, err := client.SignLayer(ctx, layerID, ledger.SignOpts{Rekor: true})
			return err
		}},
	}
	for _, tt := range calls {
		t.Run(tt.name, func(t *testing.T) {
			got = nil
			if err := tt.call(); err != nil {
				t.Fatalf("%s: %v", tt.name, err)
			}
			want := seen{tt.method, tt.path, layerID}
			if len(got) != 1 || got[0] != want {
				t.Fatalf("%s: got %+v, want %+v", tt.name, got, want)
			}
		})
	}
}

// TestQueryTestDouble_MatchesOpaqueLayerRoutes keeps the querytest double aligned with
// the query-addressed routes, including extra options such as limit.
func TestQueryTestDouble_MatchesOpaqueLayerRoutes(t *testing.T) {
	const layerID = "corpus:layer:org/repo__main/repo__main"
	provider := querytest.NewStaticProvider().
		WithEventsResponse(layerID, ledger.EventsResponse{LayerID: layerID, Total: 3})
	resp, err := ledger.NewClient(provider).ListEvents(
		context.Background(), layerID, ledger.ListEventsOpts{Limit: 2},
	)
	if err != nil {
		t.Fatal(err)
	}
	if resp.LayerID != layerID || resp.Total != 3 {
		t.Fatalf("unexpected response: %+v", resp)
	}
}

// TestHTTPClient_GetLayerDocumentPreservesUnknownFields proves the raw-document
// read serves the bytes verbatim: fields types.Layer does not carry survive,
// and the route and layer id match GetLayer's own.
func TestHTTPClient_GetLayerDocumentPreservesUnknownFields(t *testing.T) {
	body := `{"metadata":{},"events":[],"needs_review":[],"future":{"signature":"retained"}}`
	var gotPath, gotLayerID, gotAuth string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath, gotLayerID, gotAuth = r.URL.Path, r.URL.Query().Get("layer_id"), r.Header.Get("Authorization")
		_, _ = w.Write([]byte(body))
	}))
	defer srv.Close()

	client := ledger.NewHTTPClient(srv.URL, ledger.WithAuthFunc(func(r *http.Request) {
		r.Header.Set("Authorization", "Bearer machine")
	}))
	got, err := client.GetLayerDocument(context.Background(), "corpus:org/repo")
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != body {
		t.Fatalf("document = %s, want the verbatim bytes", got)
	}
	if gotPath != "/v1/ledger/layer" || gotLayerID != "corpus:org/repo" {
		t.Fatalf("route = %s?layer_id=%s, want GetLayer's own", gotPath, gotLayerID)
	}
	if gotAuth != "Bearer machine" {
		t.Fatal("auth function was not applied")
	}
}

// TestHTTPClient_GetLayerDocumentFailsClosed proves the structural check: a
// body that is not a JSON object holding metadata, events and needs_review
// is refused, as is every non-2xx status.
func TestHTTPClient_GetLayerDocumentFailsClosed(t *testing.T) {
	for _, body := range []string{`{}`, `null`, `{broken`, `{"metadata":null,"events":[],"needs_review":[]}`, `{"metadata":{},"events":{},"needs_review":[]}`} {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { _, _ = w.Write([]byte(body)) }))
		_, err := ledger.NewHTTPClient(srv.URL).GetLayerDocument(context.Background(), "layer")
		srv.Close()
		if !errors.Is(err, ledger.ErrIncompleteLayerDocument) {
			t.Fatalf("body %q = %v, want ErrIncompleteLayerDocument", body, err)
		}
	}
	for _, status := range []int{302, 401, 404, 503} {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(status) }))
		_, err := ledger.NewHTTPClient(srv.URL).GetLayerDocument(context.Background(), "layer")
		srv.Close()
		var statusErr *ledger.StatusError
		if !errors.As(err, &statusErr) || statusErr.StatusCode != status {
			t.Fatalf("status %d = %v, want a StatusError carrying it", status, err)
		}
	}
}

// TestHTTPClient_ResponseLimitBoundsTheRead proves WithResponseLimit stops
// the read at the limit: an oversized body fails with ErrResponseTooLarge
// and a body at the limit passes.
func TestHTTPClient_ResponseLimitBoundsTheRead(t *testing.T) {
	const limit = 16 << 10
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write(bytes.Repeat([]byte(`{"metadata":{},"events":[],"needs_review":[]},`), limit))
	}))
	defer srv.Close()

	if _, err := ledger.NewHTTPClient(srv.URL, ledger.WithResponseLimit(limit)).GetLayerDocument(context.Background(), "layer"); !errors.Is(err, ledger.ErrResponseTooLarge) {
		t.Fatalf("oversized body = %v, want ErrResponseTooLarge", err)
	}
	atLimit := `{"metadata":{},"events":[],"needs_review":[]}`
	srv2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { _, _ = w.Write([]byte(atLimit)) }))
	defer srv2.Close()
	if _, err := ledger.NewHTTPClient(srv2.URL, ledger.WithResponseLimit(int64(len(atLimit)))).GetLayerDocument(context.Background(), "layer"); err != nil {
		t.Fatalf("body at the limit = %v, want it accepted", err)
	}
}

// TestHTTPClient_GetLayerDocumentTransportFailures proves the read's
// failure edges: a cancelled context before the call, and a connection
// that drops mid-body — both errors, never a partial document.
func TestHTTPClient_GetLayerDocumentTransportFailures(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"metadata":{},"events":[],"needs_review":[]}`))
	}))
	defer srv.Close()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := ledger.NewHTTPClient(srv.URL).GetLayerDocument(ctx, "layer"); !errors.Is(err, context.Canceled) {
		t.Fatalf("cancelled context = %v, want context.Canceled", err)
	}

	dropped := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// A declared length the body never reaches: the client's read
		// fails mid-body once the server closes the short connection.
		w.Header().Set("Content-Length", "4096")
		_, _ = w.Write([]byte(`{"metadata":{`))
	}))
	defer dropped.Close()
	if _, err := ledger.NewHTTPClient(dropped.URL).GetLayerDocument(context.Background(), "layer"); err == nil {
		t.Fatal("a dropped body was served as a document")
	} else if errors.Is(err, ledger.ErrIncompleteLayerDocument) {
		t.Fatalf("a dropped body reported the structural check (%v), want the read error", err)
	}
}

// TestHTTPClient_ResponseLimitNeverMasksStatuses proves the status wins
// over the limit: an error response whose body exceeds the configured
// limit is still the StatusError carrying its code, not the read limit.
func TestHTTPClient_ResponseLimitNeverMasksStatuses(t *testing.T) {
	for _, status := range []int{http.StatusUnauthorized, http.StatusServiceUnavailable} {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(status)
			_, _ = w.Write(bytes.Repeat([]byte("core is unwell "), 4096))
		}))
		_, err := ledger.NewHTTPClient(srv.URL, ledger.WithResponseLimit(128)).GetLayerDocument(context.Background(), "layer")
		srv.Close()
		var statusErr *ledger.StatusError
		if !errors.As(err, &statusErr) || statusErr.StatusCode != status {
			t.Fatalf("status %d with an oversized body = %v, want the StatusError carrying it", status, err)
		}
		if errors.Is(err, ledger.ErrResponseTooLarge) {
			t.Fatalf("status %d masked by the response limit", status)
		}
	}
}

// TestHTTPClient_ResponseLimitAtMaximumStillReads proves the one-past
// probe cannot overflow: at the maximum int64 limit a normal body reads
// whole instead of wrapping the reader negative and serving it empty.
func TestHTTPClient_ResponseLimitAtMaximumStillReads(t *testing.T) {
	body := `{"metadata":{},"events":[],"needs_review":[]}`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { _, _ = w.Write([]byte(body)) }))
	defer srv.Close()

	got, err := ledger.NewHTTPClient(srv.URL, ledger.WithResponseLimit(math.MaxInt64)).GetLayerDocument(context.Background(), "layer")
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != body {
		t.Fatalf("document = %s, want the verbatim bytes", got)
	}
}
