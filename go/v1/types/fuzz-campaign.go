// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

package types

type FuzzCampaign struct {
	AdvisoryFiles       []string              `json:"advisory_files,omitempty"`
	Artifact            interface{}           `json:"artifact"`
	Batches             []Batch               `json:"batches,omitempty"`
	Bugs                []Bug                 `json:"bugs,omitempty"`
	CampaignId          string                `json:"campaign_id"`
	Duration            *FuzzCampaignDuration `json:"duration,omitempty"`
	Generated           string                `json:"generated"`
	HarnessVersion      *string               `json:"harness_version,omitempty"`
	ManifestGenerated   *string               `json:"manifest_generated,omitempty"`
	Narrative           *string               `json:"narrative,omitempty"`
	NegativeControls    *string               `json:"negative_controls,omitempty"`
	Patterns            []Pattern             `json:"patterns,omitempty"`
	RemediationPriority *string               `json:"remediation_priority,omitempty"`
	Repositories        []string              `json:"repositories"`
	SourceReferences    []string              `json:"source_references"`
	Totals              *FuzzCampaignTotals   `json:"totals,omitempty"`
}

type Batch struct {
	Batch         string      `json:"batch"`
	Bugs          interface{} `json:"bugs"`
	FuzzFunctions interface{} `json:"fuzz_functions"`
	Targets       interface{} `json:"targets"`
}

type Bug struct {
	Advisory   interface{} `json:"advisory,omitempty"`
	Cvss       *string     `json:"cvss,omitempty"`
	Cwe        *string     `json:"cwe,omitempty"`
	Method     *string     `json:"method,omitempty"`
	Narrative  *string     `json:"narrative,omitempty"`
	Num        int         `json:"num"`
	References []string    `json:"references,omitempty"`
	Target     string      `json:"target"`
	Title      string      `json:"title"`
}

type FuzzCampaignDuration struct {
	End   string `json:"end"`
	Start string `json:"start"`
}

type FuzzCampaignTotals struct {
	Advisories      interface{} `json:"advisories,omitempty"`
	Bugs            interface{} `json:"bugs,omitempty"`
	BugsByFuzz      interface{} `json:"bugs_by_fuzz,omitempty"`
	BugsMediumPlus  interface{} `json:"bugs_medium_plus,omitempty"`
	ExecutedTargets interface{} `json:"executed_targets,omitempty"`
	FuzzFunctions   interface{} `json:"fuzz_functions,omitempty"`
	FuzzedTargets   interface{} `json:"fuzzed_targets,omitempty"`
	HitRate         *string     `json:"hit_rate,omitempty"`
	Repos           interface{} `json:"repos,omitempty"`
	RunsLogged      interface{} `json:"runs_logged,omitempty"`
	Targets         interface{} `json:"targets,omitempty"`
}

type Pattern struct {
	Advisory   interface{} `json:"advisory,omitempty"`
	Name       string      `json:"name"`
	Rate       *string     `json:"rate,omitempty"`
	References []string    `json:"references,omitempty"`
	Tested     int         `json:"tested"`
	Vulnerable int         `json:"vulnerable"`
}
