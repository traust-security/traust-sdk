package ledger

import (
	"github.com/traust-security/traust-sdk/go/v1/types"
)

// BatchSubmitInput is the body for POST /v1/ledger/layers/{layer_id}/submit.
type BatchSubmitInput struct {
	SourceRef   string                   `json:"source_ref"`
	RecordedAt  string                   `json:"recorded_at"`
	Events      []map[string]interface{} `json:"events"`
	NeedsReview []map[string]interface{} `json:"needs_review,omitempty"`
}

// ResolveInput is the body for POST /v1/ledger/layers/{layer_id}/resolve.
type ResolveInput struct {
	Key      string `json:"key"`
	Decision string `json:"decision"`
	Note     string `json:"note,omitempty"`
}

// StampInput is the body for POST /v1/ledger/layers/{layer_id}/stamp.
// Fingerprints maps finding_ref -> fingerprint; the ledger backfills each
// matching event's fingerprint (never overwriting) and re-signs the layer.
type StampInput struct {
	Fingerprints map[string]string `json:"fingerprints"`
}

// FingerprintInput is the body for POST /v1/ledger/fingerprint.
type FingerprintInput struct {
	Findings   []map[string]interface{} `json:"findings"`
	Repository string                   `json:"repository,omitempty"`
}

// EventMeta is the shared prefix for human-lane event bodies.
type EventMeta struct {
	LayerID    string `json:"layer_id"`
	RecordedAt string `json:"recorded_at"`
}

// CountersignInput carries a human countersign decision for a finding.
// Actor identity is resolved server-side from the auth context.
type CountersignInput struct {
	EventMeta
	FindingRef         string `json:"finding_ref,omitempty"`
	FindingFingerprint string `json:"finding_fingerprint,omitempty"`
	Verdict            string `json:"verdict,omitempty"`
	Decision           string `json:"decision,omitempty"`
	Rationale          string `json:"rationale,omitempty"`
	Justification      string `json:"justification,omitempty"`
}

// SeverityInput carries a human severity override for a finding.
type SeverityInput struct {
	EventMeta
	FindingRef    string `json:"finding_ref"`
	Severity      string `json:"severity"`
	Rationale     string `json:"rationale,omitempty"`
	Justification string `json:"justification,omitempty"`
}

// SignOpts configures the SignLayer request.
type SignOpts struct {
	Rekor bool
}

// InitializeInput is the body for POST /v1/ledger/layers/{layer_id}/initialize:
// a complete layer shell and the storage product_repo it belongs to. Database
// ledgers require ProductRepoID; file ledgers accept nil.
type InitializeInput struct {
	ProductRepoID *string     `json:"product_repo_id"`
	Layer         types.Layer `json:"layer"`
}
