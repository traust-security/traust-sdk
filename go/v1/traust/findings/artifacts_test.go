package findings_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/google/uuid"

	"github.com/traust-security/traust-sdk/go/v1/traust"
	"github.com/traust-security/traust-sdk/go/v1/traust/findings"
	"github.com/traust-security/traust-sdk/go/v1/traust/traustest"
)

// content answers one artifact's bytes under the headers core documents: the kind, the digest
// as a quoted ETag, and the media type.
func content(kind, mediaType, digest string, body []byte) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		if kind != "" {
			w.Header().Set("Traust-Artifact-Kind", kind)
		}
		w.Header().Set("ETag", `"`+digest+`"`)
		w.Header().Set("Content-Type", mediaType)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(body)
	})
}

// Get reads the artifact in one call: the kind and digest from the content response's
// headers, the exact bytes — and their size — from its body.
func TestGetReadsTheArtifactsHeadersAndBody(t *testing.T) {
	id, body := uuid.New(), []byte(`{"title":"the audit"}`)
	core := traustest.New(t)
	core.Handle("GET /v1/artifacts/{id}/content", content("report", "application/json", "4c7f0b9e2a1c3d4f", body))
	obj, err := findings.New(core.Client()).Get(context.Background(), id)
	if err != nil {
		t.Fatal(err)
	}
	if obj.Kind != traust.ArtifactKindReport {
		t.Errorf("kind %q, want report", obj.Kind)
	}
	if obj.Digest != "4c7f0b9e2a1c3d4f" {
		t.Errorf("digest %q, want the ETag unquoted", obj.Digest)
	}
	if obj.MediaType != "application/json" {
		t.Errorf("media type %q, want the content's own", obj.MediaType)
	}
	if obj.Size != int64(len(body)) {
		t.Errorf("size %d, want %d", obj.Size, len(body))
	}
	if string(obj.Body) != string(body) {
		t.Errorf("body %q, want the exact bytes", obj.Body)
	}
}

// Get reads content no Content-Length header describes — a proxy serving it chunked, a
// transport decompressing it — sizing the artifact from the bytes it read.
func TestGetReadsContentWithoutALengthHeader(t *testing.T) {
	id, body := uuid.New(), []byte(`{"title":"the audit"}`)
	core := traustest.New(t)
	core.Handle("GET /v1/artifacts/{id}/content", http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Traust-Artifact-Kind", "report")
		w.Header().Set("ETag", `"4c7f0b9e2a1c3d4f"`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if flushable, ok := w.(http.Flusher); ok {
			flushable.Flush() // the headers go out before the body: no Content-Length, chunked
		}
		_, _ = w.Write(body)
	}))
	obj, err := findings.New(core.Client()).Get(context.Background(), id)
	if err != nil {
		t.Fatal(err)
	}
	if string(obj.Body) != string(body) {
		t.Errorf("body %q, want the exact bytes", obj.Body)
	}
	if obj.Size != int64(len(body)) {
		t.Errorf("size %d, want %d from the bytes read", obj.Size, len(body))
	}
}

// Get fails closed on a content response without a kind it can type: a missing header and one
// naming no kind core publishes both refuse, never bytes without a kind.
func TestGetRefusesContentWithoutAKnownKind(t *testing.T) {
	for name, kind := range map[string]string{"no Traust-Artifact-Kind header": "", "a kind core does not publish": "sofa"} {
		t.Run(name, func(t *testing.T) {
			core := traustest.New(t)
			core.Handle("GET /v1/artifacts/{id}/content", content(kind, "text/plain", "4c7f0b9e2a1c3d4f", []byte("bytes")))
			_, err := findings.New(core.Client()).Get(context.Background(), uuid.New())
			if !errors.Is(err, findings.ErrUnknownKind) {
				t.Fatalf("Get(%q): %v, want ErrUnknownKind", kind, err)
			}
		})
	}
}

// Decode reads a report artifact as the contract type core publishes, not bytes to parse.
func TestDecodeReadsAReportAsItsContractType(t *testing.T) {
	core := traustest.New(t)
	core.Handle("GET /v1/artifacts/{id}/content", content("report", "application/json", "4c7f0b9e2a1c3d4f",
		[]byte(`{"title":"the audit","metadata":{"scope":"the repository"},"findings":[]}`)))
	obj, err := findings.New(core.Client()).Get(context.Background(), uuid.New())
	if err != nil {
		t.Fatal(err)
	}
	report, err := findings.Decode[traust.ReportArtifact](obj)
	if err != nil {
		t.Fatal(err)
	}
	if report.Title != "the audit" {
		t.Errorf("title %q, want the report's own", report.Title)
	}
}

// Decode refuses a kind mismatch: an object of one contract never reads as another.
func TestDecodeRefusesAKindMismatch(t *testing.T) {
	core := traustest.New(t)
	core.Handle("GET /v1/artifacts/{id}/content", content("threat-model", "application/json", "4c7f0b9e2a1c3d4f",
		[]byte(`{"provenance":"the audit"}`)))
	obj, err := findings.New(core.Client()).Get(context.Background(), uuid.New())
	if err != nil {
		t.Fatal(err)
	}
	if _, err := findings.Decode[traust.ReportArtifact](obj); !errors.Is(err, findings.ErrKindMismatch) {
		t.Fatalf("decoding a threat model as a report: %v, want ErrKindMismatch", err)
	}
}

// Decode refuses a media type that carries no document: the kinds without a contract — a diff,
// an sbom, a feed — are read from the Object's Body, never decoded.
func TestDecodeRefusesAMediaTypeThatIsNotJSON(t *testing.T) {
	core := traustest.New(t)
	core.Handle("GET /v1/artifacts/{id}/content", content("remediation-patch.diff", "text/x-diff", "4c7f0b9e2a1c3d4f",
		[]byte("--- a/file\n+++ b/file\n")))
	obj, err := findings.New(core.Client()).Get(context.Background(), uuid.New())
	if err != nil {
		t.Fatal(err)
	}
	if string(obj.Body) != "--- a/file\n+++ b/file\n" {
		t.Fatalf("body %q, want the diff's exact bytes", obj.Body)
	}
	if _, err := findings.Decode[traust.ReportArtifact](obj); !errors.Is(err, findings.ErrKindMismatch) {
		t.Fatalf("decoding a diff as a report: %v, want ErrKindMismatch", err)
	}
	reportish := &findings.Object{Kind: traust.ArtifactKindReport, MediaType: "text/markdown", Body: []byte("# the audit")}
	if _, err := findings.Decode[traust.ReportArtifact](reportish); !errors.Is(err, findings.ErrNotJSON) {
		t.Fatalf("decoding a markdown report: %v, want ErrNotJSON", err)
	}
}

// Decode refuses a null body: a kind and media type that promise a document still carry none,
// so a null report never reads as an empty one.
func TestDecodeRefusesANullBody(t *testing.T) {
	nullish := &findings.Object{Kind: traust.ArtifactKindReport, MediaType: "application/json", Body: []byte(" null ")}
	if _, err := findings.Decode[traust.ReportArtifact](nullish); err == nil {
		t.Fatal("decoding a null report: no error, want a refusal")
	}
}
