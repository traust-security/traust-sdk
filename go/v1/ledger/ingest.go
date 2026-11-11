package ledger

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

type submissionOp[In any] struct {
	kind IngestKind
}

func (o submissionOp[In]) Submit(ctx context.Context, p Provider, in In) (SubmitResponse, error) {
	var zero SubmitResponse

	raw, err := json.Marshal(in)
	if err != nil {
		return zero, ingestErr(o.kind, PhaseMarshal, err)
	}

	resp, err := p.Ingest(ctx, IngestMeta{Kind: o.kind}, raw)
	return decodeSubmitResponse(resp, err, o.kind)
}

func decodeSubmitResponse(resp []byte, err error, kind IngestKind) (SubmitResponse, error) {
	var zero SubmitResponse
	if err != nil {
		return zero, ingestErr(kind, PhaseIngest, err)
	}

	var out SubmitResponse
	if err := json.Unmarshal(resp, &out); err != nil {
		return zero, ingestErr(kind, PhaseDecode, err)
	}
	return out, nil
}

func decodeResolveResponse(resp []byte, err error) (ResolveResponse, error) {
	var zero ResolveResponse
	if err != nil {
		return zero, fmt.Errorf("ingest resolve: %w", err)
	}

	var out ResolveResponse
	if err := json.Unmarshal(resp, &out); err != nil {
		return zero, fmt.Errorf("ingest resolve decode: %w", err)
	}
	return out, nil
}

func decodeFingerprintResponse(resp []byte, err error) (FingerprintResponse, error) {
	var zero FingerprintResponse
	if err != nil {
		return zero, fmt.Errorf("ingest fingerprint: %w", err)
	}

	var out FingerprintResponse
	if err := json.Unmarshal(resp, &out); err != nil {
		return zero, fmt.Errorf("ingest fingerprint decode: %w", err)
	}
	return out, nil
}

func decodeStampResponse(resp []byte, err error) (StampResponse, error) {
	var zero StampResponse
	if err != nil {
		return zero, fmt.Errorf("ingest stamp: %w", err)
	}
	var out StampResponse
	if err := json.Unmarshal(resp, &out); err != nil {
		return zero, fmt.Errorf("ingest stamp decode: %w", err)
	}
	return out, nil
}

func decodeSignResponse(resp []byte, err error) (SignResponse, error) {
	var zero SignResponse
	if err != nil {
		return zero, fmt.Errorf("ingest sign: %w", err)
	}
	var out SignResponse
	if err := json.Unmarshal(resp, &out); err != nil {
		return zero, fmt.Errorf("ingest sign decode: %w", err)
	}
	return out, nil
}

// BatchSubmit sends pre-formed events to POST /v1/ledger/layer/submit?layer_id=.
func BatchSubmit(ctx context.Context, p Provider, layerID string, input BatchSubmitInput) (SubmitResponse, error) {
	raw, err := json.Marshal(input)
	if err != nil {
		return SubmitResponse{}, ingestErr(KindTriage, PhaseMarshal, err)
	}
	resp, err := p.Submit(ctx, layerID, raw)
	return decodeSubmitResponse(resp, err, KindTriage)
}

// ResolveReviewItem resolves a needs_review queue item.
func ResolveReviewItem(ctx context.Context, p Provider, layerID string, input ResolveInput) (ResolveResponse, error) {
	raw, err := json.Marshal(input)
	if err != nil {
		return ResolveResponse{}, fmt.Errorf("ingest resolve marshal: %w", err)
	}
	path := LayerPath(LayerOpResolve, layerID, nil)
	resp, err := p.Post(ctx, path, raw)
	return decodeResolveResponse(resp, err)
}

// ComputeFingerprints stamps fingerprints on findings via POST /v1/ledger/fingerprint.
func ComputeFingerprints(ctx context.Context, p Provider, input FingerprintInput) (FingerprintResponse, error) {
	raw, err := json.Marshal(input)
	if err != nil {
		return FingerprintResponse{}, fmt.Errorf("ingest fingerprint marshal: %w", err)
	}
	resp, err := p.Post(ctx, "/v1/ledger/fingerprint", raw)
	return decodeFingerprintResponse(resp, err)
}

// StampEventIdentities backfills event fingerprints on a layer via
// POST /v1/ledger/layer/stamp?layer_id=. The ledger stamps each event whose
// finding_ref appears in input.Fingerprints (never overwriting an existing
// fingerprint) and re-signs the layer, returning the new root and count stamped.
func StampEventIdentities(ctx context.Context, p Provider, layerID string, input StampInput) (StampResponse, error) {
	raw, err := json.Marshal(input)
	if err != nil {
		return StampResponse{}, fmt.Errorf("ingest stamp marshal: %w", err)
	}
	path := LayerPath(LayerOpStamp, layerID, nil)
	resp, err := p.Post(ctx, path, raw)
	return decodeStampResponse(resp, err)
}

// SignLayer requests the ledger service to sign a layer's merkle tree.
func SignLayer(ctx context.Context, p Provider, layerID string, opts SignOpts) (SignResponse, error) {
	values := url.Values{}
	if opts.Rekor {
		values.Set("rekor", "true")
	}
	path := LayerPath(LayerOpSign, layerID, values)
	resp, err := p.Post(ctx, path, nil)
	return decodeSignResponse(resp, err)
}
