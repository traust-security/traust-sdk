// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

package types

import (
	"github.com/traust-security/traust-sdk/go/v1/enums"
)

type IsolationReview struct {
	Gaps       []IsolationReviewGap    `json:"gaps"`
	Interfaces []Interface             `json:"interfaces"`
	Metadata   IsolationReviewMetadata `json:"metadata"`
	Notes      *string                 `json:"notes,omitempty"`
	Posture    IsolationReviewPosture  `json:"posture"`
	Title      string                  `json:"title"`
}

type Dimension struct {
	Evidence  []string                   `json:"evidence,omitempty"`
	Rationale *string                    `json:"rationale,omitempty"`
	Result    enums.IsolationCheckResult `json:"result"`
}

type Interface struct {
	Complexity     enums.InterfaceComplexity `json:"complexity"`
	Dimensions     InterfaceDimensions       `json:"dimensions"`
	Exposure       enums.InterfaceExposure   `json:"exposure"`
	Id             string                    `json:"id"`
	Kind           enums.InterfaceKind       `json:"kind"`
	Name           string                    `json:"name"`
	Repos          []string                  `json:"repos,omitempty"`
	SharedInstance interface{}               `json:"shared_instance,omitempty"`
}

type InterfaceDimensions struct {
	Authentication Dimension `json:"authentication"`
	Connectivity   Dimension `json:"connectivity"`
	Encryption     Dimension `json:"encryption"`
	Hygiene        Dimension `json:"hygiene"`
	Privilege      Dimension `json:"privilege"`
}

type IsolationReviewGap struct {
	Description string         `json:"description"`
	InterfaceId string         `json:"interface_id"`
	Remediation string         `json:"remediation"`
	Severity    enums.Severity `json:"severity"`
}

type IsolationReviewMetadata struct {
	Framework      *string  `json:"framework,omitempty"`
	GraphRef       string   `json:"graph_ref"`
	HarnessVersion string   `json:"harness_version"`
	Repos          []string `json:"repos"`
	ReviewedAt     string   `json:"reviewed_at"`
	Service        string   `json:"service"`
}

type IsolationReviewPosture struct {
	DimensionRollup    map[string]int         `json:"dimension_rollup,omitempty"`
	InterfacesReviewed *int                   `json:"interfaces_reviewed,omitempty"`
	Overall            enums.IsolationPosture `json:"overall"`
	Summary            string                 `json:"summary"`
}
