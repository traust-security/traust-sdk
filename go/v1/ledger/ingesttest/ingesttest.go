// Package ingesttest provides test doubles for the ingest package.
package ingesttest

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/traust-security/traust-sdk/go/v1/ledger"
)

// StaticProvider is a test double that returns pre-configured responses.
// It implements ledger.Provider.
type StaticProvider struct {
	ingestResponses map[ledger.IngestKind][]byte
	submitResponse  []byte
	postResponses   map[string][]byte
	err             error
	ingestCalls     []IngestCall
	submitCalls     []SubmitCall
	postCalls       []PostCall
}

// IngestCall records one invocation of Ingest.
type IngestCall struct {
	Meta    ledger.IngestMeta
	Payload []byte
}

// SubmitCall records one invocation of Submit.
type SubmitCall struct {
	LayerID string
	Payload []byte
}

// PostCall records one invocation of Post.
type PostCall struct {
	Path    string
	Payload []byte
}

// NewStaticProvider creates an empty StaticProvider.
func NewStaticProvider() *StaticProvider {
	return &StaticProvider{
		ingestResponses: make(map[ledger.IngestKind][]byte),
		postResponses:   make(map[string][]byte),
	}
}

// WithResponse sets the response for a specific human-lane kind.
func (p *StaticProvider) WithResponse(kind ledger.IngestKind, resp interface{}) *StaticProvider {
	data, _ := json.Marshal(resp)
	p.ingestResponses[kind] = data
	return p
}

// WithSubmitResponse sets the response for batch submit calls.
func (p *StaticProvider) WithSubmitResponse(resp ledger.SubmitResponse) *StaticProvider {
	data, _ := json.Marshal(resp)
	p.submitResponse = data
	return p
}

// WithCountersignResponse is a shorthand for WithResponse(KindCountersign, resp).
func (p *StaticProvider) WithCountersignResponse(resp ledger.SubmitResponse) *StaticProvider {
	return p.WithResponse(ledger.KindCountersign, resp)
}

// WithPostResponse sets the response for a specific Post path.
func (p *StaticProvider) WithPostResponse(path string, resp interface{}) *StaticProvider {
	data, _ := json.Marshal(resp)
	p.postResponses[path] = data
	return p
}

// WithSignResponse sets the response for POST /v1/ledger/layer/sign?layer_id={layerID}.
func (p *StaticProvider) WithSignResponse(layerID string, resp ledger.SignResponse) *StaticProvider {
	return p.WithPostResponse(ledger.LayerPath(ledger.LayerOpSign, layerID, nil), resp)
}

// WithError makes all Provider calls return the given error.
func (p *StaticProvider) WithError(err error) *StaticProvider {
	p.err = err
	return p
}

// Ingest implements ledger.Provider.
func (p *StaticProvider) Ingest(_ context.Context, meta ledger.IngestMeta, payload []byte) ([]byte, error) {
	p.ingestCalls = append(p.ingestCalls, IngestCall{Meta: meta, Payload: append([]byte(nil), payload...)})
	if p.err != nil {
		return nil, p.err
	}
	if resp, ok := p.ingestResponses[meta.Kind]; ok {
		return resp, nil
	}
	return json.Marshal(ledger.SubmitResponse{ID: "test-id", Status: "accepted"})
}

// Submit implements ledger.Provider.
func (p *StaticProvider) Submit(_ context.Context, layerID string, payload []byte) ([]byte, error) {
	p.submitCalls = append(p.submitCalls, SubmitCall{LayerID: layerID, Payload: append([]byte(nil), payload...)})
	if p.err != nil {
		return nil, p.err
	}
	if len(p.submitResponse) > 0 {
		return p.submitResponse, nil
	}
	return json.Marshal(ledger.SubmitResponse{ID: "batch-test-id", Status: "accepted"})
}

// Post implements ledger.Provider.
func (p *StaticProvider) Post(_ context.Context, path string, payload []byte) ([]byte, error) {
	p.postCalls = append(p.postCalls, PostCall{Path: path, Payload: append([]byte(nil), payload...)})
	if p.err != nil {
		return nil, p.err
	}
	if resp, ok := p.postResponses[path]; ok {
		return resp, nil
	}
	return []byte("{}"), nil
}

// IngestCalls returns all recorded Ingest invocations.
func (p *StaticProvider) IngestCalls() []IngestCall { return p.ingestCalls }

// SubmitCalls returns all recorded Submit invocations.
func (p *StaticProvider) SubmitCalls() []SubmitCall { return p.submitCalls }

// PostCalls returns all recorded Post invocations.
func (p *StaticProvider) PostCalls() []PostCall { return p.postCalls }

// Calls is an alias for IngestCalls for backward compatibility in tests.
func (p *StaticProvider) Calls() []IngestCall { return p.ingestCalls }

// FixtureSignResponse returns a typical layer signing response.
func FixtureSignResponse() ledger.SignResponse {
	return ledger.SignResponse{Status: "signed", Method: "cosign", LayerID: "repo-a"}
}

// Query implements ledger.Provider for read calls not exercised by this double.
func (p *StaticProvider) Query(context.Context, string, string) ([]byte, error) {
	return nil, fmt.Errorf("query response not configured")
}
