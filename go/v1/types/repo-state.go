// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

package types

import (
	"github.com/traust-security/traust-sdk/go/v1/enums"
)

type RepoState struct {
	Comparison *Comparison            `json:"comparison,omitempty"`
	Fork       *bool                  `json:"fork,omitempty"`
	ForkOf     *string                `json:"fork_of,omitempty"`
	ObservedAt string                 `json:"observed_at"`
	PushedAt   *string                `json:"pushed_at,omitempty"`
	Status     enums.ComparisonStatus `json:"status"`
}

type Comparison struct {
	AheadBy             int  `json:"ahead_by"`
	ChangedLines        int  `json:"changed_lines"`
	DepsOnly            bool `json:"deps_only"`
	FirstPartyFiles     *int `json:"first_party_files,omitempty"`
	IacFiles            *int `json:"iac_files,omitempty"`
	SecretHits          *int `json:"secret_hits,omitempty"`
	SensitiveChanged    bool `json:"sensitive_changed"`
	SensitiveLines      int  `json:"sensitive_lines"`
	ThreatModelFiles    *int `json:"threat_model_files,omitempty"`
	TreeFirstPartyFiles *int `json:"tree_first_party_files,omitempty"`
	Truncated           bool `json:"truncated"`
}
