// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

package types

import (
	"github.com/traust-security/traust-sdk/go/v1/enums"
)

type RunManifest struct {
	Approval *RunManifestApproval    `json:"approval,omitempty"`
	Artifact interface{}             `json:"artifact"`
	Budget   *RunManifestBudget      `json:"budget,omitempty"`
	Config   []RunManifestConfigItem `json:"config,omitempty"`
	Degraded map[string]int          `json:"degraded,omitempty"`
	Progress *RunManifestProgress    `json:"progress,omitempty"`
	Sessions []Session               `json:"sessions,omitempty"`
}

type RunManifestApproval struct {
	ApprovedAt string `json:"approved_at"`
	Approver   string `json:"approver"`
}

type RunManifestBudget struct {
	BudgetUsd   interface{}              `json:"budget_usd,omitempty"`
	Enforcement *enums.BudgetEnforcement `json:"enforcement,omitempty"`
	EstimateUsd interface{}              `json:"estimate_usd"`
	Lane        string                   `json:"lane"`
	SpentUsd    string                   `json:"spent_usd"`
	Status      enums.BudgetStatus       `json:"status"`
}

type RunManifestConfigItem struct {
	ArtifactId *string `json:"artifact_id,omitempty"`
	Digest     string  `json:"digest"`
	Name       string  `json:"name"`
}

type RunManifestProgress struct {
	ShardsDone []string    `json:"shards_done,omitempty"`
	Step       interface{} `json:"step,omitempty"`
	UpdatedAt  string      `json:"updated_at"`
}

type Session struct {
	ExecutorSession interface{}         `json:"executor_session"`
	Inputs          []SessionInputsItem `json:"inputs"`
	Mode            string              `json:"mode"`
	Models          map[string]string   `json:"models"`
	Output          SessionOutput       `json:"output"`
	Skill           string              `json:"skill"`
	TaskSha256      string              `json:"task_sha256"`
}

type SessionInputsItem struct {
	Bytes  int    `json:"bytes"`
	Path   string `json:"path"`
	Sha256 string `json:"sha256"`
}

type SessionOutput struct {
	ArtifactId interface{} `json:"artifact_id"`
	Path       interface{} `json:"path"`
	Sha256     interface{} `json:"sha256"`
}
