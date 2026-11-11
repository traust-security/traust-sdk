package ledger

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/traust-security/traust-sdk/go/v1/types"
)

// Client binds a Provider for typed submission methods.
type Client struct {
	p Provider
}

// NewClient creates a Client backed by the given Provider.
func NewClient(p Provider) *Client { return &Client{p: p} }

// BatchSubmit sends pre-formed events via POST /v1/ledger/layer/submit?layer_id=.
func (c *Client) BatchSubmit(ctx context.Context, layerID string, input BatchSubmitInput) (SubmitResponse, error) {
	return BatchSubmit(ctx, c.p, layerID, input)
}

// ResolveReviewItem resolves a needs_review queue item.
func (c *Client) ResolveReviewItem(ctx context.Context, layerID string, input ResolveInput) (ResolveResponse, error) {
	return ResolveReviewItem(ctx, c.p, layerID, input)
}

// ComputeFingerprints stamps fingerprints on findings.
func (c *Client) ComputeFingerprints(ctx context.Context, input FingerprintInput) (FingerprintResponse, error) {
	return ComputeFingerprints(ctx, c.p, input)
}

// StampEventIdentities backfills event fingerprints on a layer and re-signs it.
func (c *Client) StampEventIdentities(ctx context.Context, layerID string, input StampInput) (StampResponse, error) {
	return StampEventIdentities(ctx, c.p, layerID, input)
}

// SignLayer requests the ledger service to sign a layer's merkle tree.
func (c *Client) SignLayer(ctx context.Context, layerID string, opts SignOpts) (SignResponse, error) {
	return SignLayer(ctx, c.p, layerID, opts)
}

// SubmitCountersign submits a human countersign decision for a finding.
func (c *Client) SubmitCountersign(ctx context.Context, in CountersignInput) (SubmitResponse, error) {
	return countersignOp.Submit(ctx, c.p, in)
}

// SubmitSeverity submits a human severity override for a finding.
func (c *Client) SubmitSeverity(ctx context.Context, in SeverityInput) (SubmitResponse, error) {
	return severityOp.Submit(ctx, c.p, in)
}

// Health checks ledger service availability.
func (c *Client) Health(ctx context.Context) (HealthResponse, error) {
	body, err := c.p.Query(ctx, http.MethodGet, "/healthz")
	if err != nil {
		return HealthResponse{}, fmt.Errorf("query: health: %w", err)
	}

	var out HealthResponse
	if err := json.Unmarshal(body, &out); err != nil {
		return HealthResponse{}, fmt.Errorf("query: decode health response: %w", err)
	}
	return out, nil
}

// GetLayer returns the full layer document for layerID.
func (c *Client) GetLayer(ctx context.Context, layerID string) (types.Layer, error) {
	path := LayerPath(LayerOpDocument, layerID, nil)
	body, err := c.p.Query(ctx, http.MethodGet, path)
	if err != nil {
		return types.Layer{}, fmt.Errorf("query: get layer: %w", err)
	}

	var out types.Layer
	if err := json.Unmarshal(body, &out); err != nil {
		return types.Layer{}, fmt.Errorf("query: decode layer response: %w", err)
	}
	return out, nil
}

// GetFindings returns resolved finding dispositions for layerID.
func (c *Client) GetFindings(ctx context.Context, layerID string) (FindingsResponse, error) {
	path := LayerPath(LayerOpFindings, layerID, nil)
	body, err := c.p.Query(ctx, http.MethodGet, path)
	if err != nil {
		return FindingsResponse{}, fmt.Errorf("query: get findings: %w", err)
	}

	var out FindingsResponse
	if err := json.Unmarshal(body, &out); err != nil {
		return FindingsResponse{}, fmt.Errorf("query: decode findings response: %w", err)
	}
	return out, nil
}

// ListFindings returns paginated findings across all layers.
func (c *Client) ListFindings(ctx context.Context, opts ListFindingsOpts) (BulkFindingsResponse, error) {
	path := "/v1/ledger/findings"
	values := url.Values{}
	if opts.Cursor != nil && *opts.Cursor != "" {
		values.Set("cursor", *opts.Cursor)
	}
	if opts.Limit > 0 {
		values.Set("limit", strconv.Itoa(opts.Limit))
	}
	if opts.SinceEpoch != nil {
		values.Set("since_epoch", strconv.Itoa(*opts.SinceEpoch))
	}
	if encoded := values.Encode(); encoded != "" {
		path += "?" + encoded
	}

	body, err := c.p.Query(ctx, http.MethodGet, path)
	if err != nil {
		return BulkFindingsResponse{}, fmt.Errorf("query: list findings: %w", err)
	}

	var out BulkFindingsResponse
	if err := json.Unmarshal(body, &out); err != nil {
		return BulkFindingsResponse{}, fmt.Errorf("query: decode bulk findings response: %w", err)
	}
	return out, nil
}

// ListEvents returns the raw event log for layerID with optional filtering.
func (c *Client) ListEvents(ctx context.Context, layerID string, opts ListEventsOpts) (EventsResponse, error) {
	values := url.Values{}
	if opts.FindingRef != nil && *opts.FindingRef != "" {
		values.Set("finding_ref", *opts.FindingRef)
	}
	if opts.SourceType != nil && *opts.SourceType != "" {
		values.Set("source_type", *opts.SourceType)
	}
	if opts.Limit > 0 {
		values.Set("limit", strconv.Itoa(opts.Limit))
	}
	if opts.Offset > 0 {
		values.Set("offset", strconv.Itoa(opts.Offset))
	}
	path := LayerPath(LayerOpEvents, layerID, values)

	body, err := c.p.Query(ctx, http.MethodGet, path)
	if err != nil {
		return EventsResponse{}, fmt.Errorf("query: list events: %w", err)
	}

	var out EventsResponse
	if err := json.Unmarshal(body, &out); err != nil {
		return EventsResponse{}, fmt.Errorf("query: decode events response: %w", err)
	}
	return out, nil
}

// VerifyLayer checks layer integrity and optional signatures.
func (c *Client) VerifyLayer(ctx context.Context, layerID string, opts VerifyOpts) (VerifyResponse, error) {
	values := url.Values{}
	if opts.CheckSignatures {
		values.Set("check_signatures", "true")
	}
	path := LayerPath(LayerOpVerify, layerID, values)

	body, err := c.p.Query(ctx, http.MethodGet, path)
	if err != nil {
		return VerifyResponse{}, fmt.Errorf("query: verify layer: %w", err)
	}

	var out VerifyResponse
	if err := json.Unmarshal(body, &out); err != nil {
		return VerifyResponse{}, fmt.Errorf("query: decode verify response: %w", err)
	}
	return out, nil
}

// Whoami returns the verified actor derived from the caller's auth token.
// The ledger records nothing — this resolves signer identity for callers that
// need to attribute a non-event write.
func (c *Client) Whoami(ctx context.Context) (types.Actor, error) {
	body, err := c.p.Query(ctx, http.MethodGet, "/v1/ledger/whoami")
	if err != nil {
		return types.Actor{}, fmt.Errorf("query: whoami: %w", err)
	}

	var out types.Actor
	if err := json.Unmarshal(body, &out); err != nil {
		return types.Actor{}, fmt.Errorf("query: decode whoami response: %w", err)
	}
	return out, nil
}

// InitializeLayer creates layerID from a complete shell under its product_repo.
// It never replaces an existing layer.
func (c *Client) InitializeLayer(ctx context.Context, layerID string, input InitializeInput) (InitializeResponse, error) {
	raw, err := json.Marshal(input)
	if err != nil {
		return InitializeResponse{}, fmt.Errorf("initialize: marshal: %w", err)
	}
	path := fmt.Sprintf("/v1/ledger/layers/%s/initialize", url.PathEscape(layerID))
	body, err := c.p.Post(ctx, path, raw)
	if err != nil {
		return InitializeResponse{}, fmt.Errorf("initialize: %w", err)
	}
	var out InitializeResponse
	if err := json.Unmarshal(body, &out); err != nil {
		return InitializeResponse{}, fmt.Errorf("initialize: decode response: %w", err)
	}
	return out, nil
}

// FindLayer returns the layer that belongs to productRepoID; found is false
// when the ledger has none.
func (c *Client) FindLayer(ctx context.Context, productRepoID string) (ref LayerRef, found bool, err error) {
	path := "/v1/ledger/layers?" + url.Values{"product_repo_id": {productRepoID}}.Encode()
	body, err := c.p.Query(ctx, http.MethodGet, path)
	var status *StatusError
	if errors.As(err, &status) && status.StatusCode == http.StatusNotFound {
		return LayerRef{}, false, nil
	}
	if err != nil {
		return LayerRef{}, false, fmt.Errorf("query: find layer: %w", err)
	}
	var out LayerListResponse
	if err := json.Unmarshal(body, &out); err != nil {
		return LayerRef{}, false, fmt.Errorf("query: decode layer list response: %w", err)
	}
	if len(out.Layers) != 1 {
		return LayerRef{}, false, fmt.Errorf("query: find layer: %d layers for one product_repo", len(out.Layers))
	}
	return out.Layers[0], true, nil
}

// ListLayers returns layers matching the given pluggable filters; with no
// filters, every layer known to the ledger service with its product_repo.
// Each filter renders into the ledger's allowlisted key:op:value registry;
// typed constructors (ReviewStatus, LayerIDs, UpdatedSince, ...) model the
// known keys and RawFilter reaches the rest.
func (c *Client) ListLayers(ctx context.Context, filters ...LayerListFilter) (LayerListResponse, error) {
	values := url.Values{}
	for _, f := range filters {
		if f == nil {
			continue
		}
		if err := f.apply(values); err != nil {
			return LayerListResponse{}, fmt.Errorf("query: list layers: %w", err)
		}
	}
	path := "/v1/ledger/layers"
	if len(values) > 0 {
		path += "?" + values.Encode()
	}
	body, err := c.p.Query(ctx, http.MethodGet, path)
	if err != nil {
		return LayerListResponse{}, fmt.Errorf("query: list layers: %w", err)
	}

	var out LayerListResponse
	if err := json.Unmarshal(body, &out); err != nil {
		return LayerListResponse{}, fmt.Errorf("query: decode layer list response: %w", err)
	}
	return out, nil
}
