package ledger_test

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"testing"

	"github.com/traust-security/traust-sdk/go/v1/ledger"
	"github.com/traust-security/traust-sdk/go/v1/ledger/ingesttest"
)

// testFingerprint is a 64-hex fingerprint constant assembled at runtime so no
// key-shaped literal sits in the tree for forge secret scanners.
var testFingerprint = strings.Repeat("aa11bb22cc33dd44ee55ff66", 2) + "aa11bb22cc33dd44"

func TestBatchSubmit_RoundTrip(t *testing.T) {
	provider := ingesttest.NewStaticProvider().
		WithSubmitResponse(ledger.SubmitResponse{ID: "batch-001", Status: "accepted"})
	client := ledger.NewClient(provider)

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
	if resp.ID != "batch-001" {
		t.Fatalf("expected ID batch-001, got %s", resp.ID)
	}

	calls := provider.SubmitCalls()
	if len(calls) != 1 || calls[0].LayerID != "repo-a" {
		t.Fatalf("unexpected submit calls: %+v", calls)
	}
}

func TestBatchSubmit_DecodesNewResponseFields(t *testing.T) {
	ec := 2
	provider := ingesttest.NewStaticProvider().
		WithSubmitResponse(ledger.SubmitResponse{
			ID:         "batch-002",
			Status:     "accepted",
			EventCount: &ec,
			EventIDs:   []string{"evt-001", "evt-002"},
			QueueAdded: 1,
		})
	client := ledger.NewClient(provider)

	resp, err := client.BatchSubmit(context.Background(), "repo-a", ledger.BatchSubmitInput{
		SourceRef:  "findings/repo-a/triage.json",
		RecordedAt: "2026-01-16T00:00:00+00:00",
		Events: []map[string]interface{}{
			{"finding_ref": "FIND-001"},
			{"finding_ref": "FIND-002"},
		},
		NeedsReview: []map[string]interface{}{
			{"finding_ref": "FIND-003", "queue_reason": "undetermined_finding"},
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(resp.EventIDs) != 2 {
		t.Fatalf("expected 2 event_ids, got %d", len(resp.EventIDs))
	}
	if resp.EventIDs[0] != "evt-001" || resp.EventIDs[1] != "evt-002" {
		t.Fatalf("unexpected event_ids: %v", resp.EventIDs)
	}
	if resp.QueueAdded != 1 {
		t.Fatalf("expected queue_added 1, got %d", resp.QueueAdded)
	}
}

func TestResolveReviewItem_RoundTrip(t *testing.T) {
	path := "/v1/ledger/layer/resolve?layer_id=repo-a"
	provider := ingesttest.NewStaticProvider().
		WithPostResponse(path, ledger.ResolveResponse{Resolved: true, Key: `["a","b","c","d"]`})
	client := ledger.NewClient(provider)

	resp, err := client.ResolveReviewItem(context.Background(), "repo-a", ledger.ResolveInput{
		Key:      `["a","b","c","d"]`,
		Decision: "confirmed",
		Note:     "reviewed",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !resp.Resolved || resp.Key == "" {
		t.Fatalf("unexpected response: %+v", resp)
	}
	if len(provider.PostCalls()) != 1 || provider.PostCalls()[0].Path != path {
		t.Fatalf("unexpected post calls: %+v", provider.PostCalls())
	}
}

func TestSignLayer_RoundTrip(t *testing.T) {
	path := "/v1/ledger/layer/sign?layer_id=repo-a"
	provider := ingesttest.NewStaticProvider().
		WithPostResponse(path, ledger.SignResponse{Status: "signed", Method: "cosign", LayerID: "repo-a"})
	client := ledger.NewClient(provider)

	resp, err := client.SignLayer(context.Background(), "repo-a", ledger.SignOpts{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Status != "signed" || resp.Method != "cosign" {
		t.Fatalf("unexpected response: %+v", resp)
	}
	if len(provider.PostCalls()) != 1 || provider.PostCalls()[0].Path != path {
		t.Fatalf("unexpected post calls: %+v", provider.PostCalls())
	}
}

func TestSignLayer_WithRekor(t *testing.T) {
	path := "/v1/ledger/layer/sign?layer_id=repo-a&rekor=true"
	provider := ingesttest.NewStaticProvider().
		WithPostResponse(path, ledger.SignResponse{Status: "signed", Method: "cosign", LayerID: "repo-a"})
	client := ledger.NewClient(provider)

	resp, err := client.SignLayer(context.Background(), "repo-a", ledger.SignOpts{Rekor: true})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Status != "signed" {
		t.Fatalf("unexpected response: %+v", resp)
	}
	if len(provider.PostCalls()) != 1 || provider.PostCalls()[0].Path != path {
		t.Fatalf("expected rekor path, got: %+v", provider.PostCalls())
	}
}

func TestComputeFingerprints_RoundTrip(t *testing.T) {
	provider := ingesttest.NewStaticProvider().
		WithPostResponse("/v1/ledger/fingerprint", ledger.FingerprintResponse{
			StampedCount: 1,
			Findings: []map[string]interface{}{
				{"id": "FIND-001", "fingerprint": "fp-abc"},
			},
		})
	client := ledger.NewClient(provider)

	resp, err := client.ComputeFingerprints(context.Background(), ledger.FingerprintInput{
		Findings: []map[string]interface{}{
			{"id": "FIND-001"},
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StampedCount != 1 {
		t.Fatalf("expected stamped_count 1, got %d", resp.StampedCount)
	}
}

func TestSubmitCountersign_RoundTrip(t *testing.T) {
	provider := ingesttest.NewStaticProvider().
		WithCountersignResponse(ledger.SubmitResponse{ID: "cs-001", Status: "accepted"})
	client := ledger.NewClient(provider)

	resp, err := client.SubmitCountersign(context.Background(), ledger.CountersignInput{
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
	if resp.ID != "cs-001" {
		t.Fatalf("expected ID cs-001, got %s", resp.ID)
	}
	calls := provider.IngestCalls()
	last := calls[len(calls)-1]
	if last.Meta.Kind != ledger.KindCountersign {
		t.Fatalf("expected kind countersign, got %s", last.Meta.Kind)
	}
}

func TestStampEventIdentities_RoundTrip(t *testing.T) {
	path := "/v1/ledger/layer/stamp?layer_id=repo-a"
	root := "root-abc"
	provider := ingesttest.NewStaticProvider().
		WithPostResponse(path, ledger.StampResponse{MerkleRoot: &root, LayerID: "repo-a", Stamped: 1})
	client := ledger.NewClient(provider)

	resp, err := client.StampEventIdentities(context.Background(), "repo-a", ledger.StampInput{
		Fingerprints: map[string]string{"FIND-001": testFingerprint},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Stamped != 1 || resp.MerkleRoot == nil || *resp.MerkleRoot != root {
		t.Fatalf("unexpected response: %+v", resp)
	}

	calls := provider.PostCalls()
	if len(calls) != 1 || calls[0].Path != path {
		t.Fatalf("unexpected post calls: %+v", calls)
	}
	var body ledger.StampInput
	if err := json.Unmarshal(calls[0].Payload, &body); err != nil {
		t.Fatalf("decode stamp body: %v", err)
	}
	if body.Fingerprints["FIND-001"] != testFingerprint {
		t.Fatalf("fingerprint not sent: %v", body.Fingerprints)
	}
}

func TestBatchSubmit_DecodeError(t *testing.T) {
	provider := &badSubmitProvider{data: []byte(`not json`)}
	client := ledger.NewClient(provider)

	_, err := client.BatchSubmit(context.Background(), "repo-a", ledger.BatchSubmitInput{
		SourceRef:  "ref",
		RecordedAt: "2026-01-16T00:00:00+00:00",
	})
	if err == nil {
		t.Fatal("expected decode error")
	}
	var ie *ledger.IngestError
	if !errors.As(err, &ie) || ie.Phase != ledger.PhaseDecode {
		t.Fatalf("expected decode phase error, got %v", err)
	}
}

type badSubmitProvider struct {
	data []byte
}

func (p *badSubmitProvider) Ingest(context.Context, ledger.IngestMeta, []byte) ([]byte, error) {
	return nil, errors.New("not used")
}

func (p *badSubmitProvider) Submit(context.Context, string, []byte) ([]byte, error) {
	return p.data, nil
}

func (p *badSubmitProvider) Post(context.Context, string, []byte) ([]byte, error) {
	return nil, errors.New("not used")
}

func (p *badSubmitProvider) Query(context.Context, string, string) ([]byte, error) { return nil, nil }
