// Code generated from traust-contracts v0.36.0. DO NOT EDIT.

package types

import (
	"github.com/traust-security/traust-sdk/go/v1/enums"
)

type ThreatModel struct {
	Assets           []map[string]interface{} `json:"assets,omitempty"`
	AttackScenarios  []map[string]interface{} `json:"attack_scenarios,omitempty"`
	Deprioritized    []map[string]interface{} `json:"deprioritized,omitempty"`
	EntryPoints      []map[string]interface{} `json:"entry_points,omitempty"`
	Mitigations      []Mitigation             `json:"mitigations,omitempty"`
	OpenQuestions    []string                 `json:"open_questions,omitempty"`
	Provenance       Provenance               `json:"provenance"`
	SubjectId        *string                  `json:"subject_id,omitempty"`
	System           string                   `json:"system"`
	SystemContext    *string                  `json:"system_context,omitempty"`
	TenantBoundaries []TenantBoundary         `json:"tenant_boundaries,omitempty"`
	Threats          []Threat                 `json:"threats"`
	UpdateHistory    []map[string]interface{} `json:"update_history,omitempty"`
}

type Mitigation struct {
	ClosesClass string   `json:"closes_class"`
	Effort      string   `json:"effort"`
	Mitigation  string   `json:"mitigation"`
	ThreatIds   []string `json:"threat_ids"`
}

type Provenance struct {
	Date           string  `json:"date"`
	HarnessVersion *string `json:"harness_version,omitempty"`
	Inputs         *string `json:"inputs,omitempty"`
	Mode           string  `json:"mode"`
	Owner          *string `json:"owner,omitempty"`
	Target         string  `json:"target"`
}

type TenantBoundary struct {
	Authentication     *string         `json:"authentication,omitempty"`
	BoundaryId         string          `json:"boundary_id"`
	Complexity         *enums.Severity `json:"complexity,omitempty"`
	Connectivity       *string         `json:"connectivity,omitempty"`
	Encryption         *string         `json:"encryption,omitempty"`
	Exposure           string          `json:"exposure"`
	Hygiene            *string         `json:"hygiene,omitempty"`
	Interface          string          `json:"interface"`
	IsolationReviewRef *string         `json:"isolation_review_ref,omitempty"`
	Kind               string          `json:"kind"`
	Privilege          *string         `json:"privilege,omitempty"`
	ThreatIds          []string        `json:"threat_ids,omitempty"`
}

type Threat struct {
	Actor               []string               `json:"actor"`
	Asset               *string                `json:"asset,omitempty"`
	AttackRefs          []string               `json:"attack_refs,omitempty"`
	Controls            *string                `json:"controls,omitempty"`
	Evidence            []string               `json:"evidence,omitempty"`
	Id                  string                 `json:"id"`
	Impact              enums.ThreatImpact     `json:"impact"`
	IsolationDimensions []string               `json:"isolation_dimensions,omitempty"`
	Likelihood          enums.ThreatLikelihood `json:"likelihood"`
	Status              enums.ThreatStatus     `json:"status"`
	Surface             *string                `json:"surface,omitempty"`
	Threat              string                 `json:"threat"`
}
