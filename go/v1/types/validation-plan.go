// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

package types

import (
	"github.com/traust-security/traust-sdk/go/v1/enums"
)

type ValidationPlan struct {
	Chains           []Chain                `json:"chains"`
	Engagement       interface{}            `json:"engagement,omitempty"`
	ScopeBindingMode enums.ScopeBindingMode `json:"scope_binding_mode"`
	Steps            []Step                 `json:"steps"`
	Target           string                 `json:"target"`
}

type Chain struct {
	ChainId         string   `json:"chain_id"`
	EntryPoint      string   `json:"entry_point"`
	FindingIds      []string `json:"finding_ids"`
	MitreAttackRefs []string `json:"mitre_attack_refs"`
	Name            string   `json:"name"`
	Narrative       string   `json:"narrative"`
	Path            []string `json:"path"`
	TerminalAsset   string   `json:"terminal_asset"`
}

type Step struct {
	Adapter              enums.ValidationAdapter             `json:"adapter"`
	ChainRef             *string                             `json:"chain_ref,omitempty"`
	Classification       enums.ValidationProbeClassification `json:"classification"`
	Cmd                  *string                             `json:"cmd,omitempty"`
	Expected             *string                             `json:"expected,omitempty"`
	FindingRef           *string                             `json:"finding_ref,omitempty"`
	Id                   string                              `json:"id"`
	NovelRef             *string                             `json:"novel_ref,omitempty"`
	Payload              *string                             `json:"payload,omitempty"`
	Preconditions        []string                            `json:"preconditions,omitempty"`
	Rollback             *string                             `json:"rollback,omitempty"`
	Skip                 *string                             `json:"skip,omitempty"`
	Summary              *string                             `json:"summary,omitempty"`
	Target               StepTarget                          `json:"target"`
	Technique            enums.ValidationTechnique           `json:"technique"`
	TriageWasNeedsReview *bool                               `json:"triage_was_needs_review,omitempty"`
	Verb                 string                              `json:"verb"`
}

type StepTarget struct {
	Artifact  interface{} `json:"artifact,omitempty"`
	Context   interface{} `json:"context,omitempty"`
	Image     interface{} `json:"image,omitempty"`
	Method    *string     `json:"method,omitempty"`
	Name      interface{} `json:"name,omitempty"`
	Namespace interface{} `json:"namespace,omitempty"`
	Path      *string     `json:"path,omitempty"`
	Port      *int        `json:"port,omitempty"`
	Pqc       *bool       `json:"pqc,omitempty"`
	Resource  interface{} `json:"resource,omitempty"`
	Runtime   *string     `json:"runtime,omitempty"`
}
