// Code generated from traust-contracts 31d0d884672149e606934c2844ffcbf3bdd4b925. DO NOT EDIT.

package types

import (
	"github.com/traust-security/traust-sdk/go/v1/enums"
)

type RepoLiveness struct {
	Generated    string              `json:"generated"`
	Repositories map[string]Liveness `json:"repositories"`
}

type Liveness struct {
	Archived     *bool                   `json:"archived,omitempty"`
	CanonicalUrl *string                 `json:"canonical_url,omitempty"`
	ObservedAt   string                  `json:"observed_at"`
	PushedAt     *string                 `json:"pushed_at,omitempty"`
	Reason       *string                 `json:"reason,omitempty"`
	RepositoryId string                  `json:"repository_id"`
	Status       enums.LivenessStatus    `json:"status"`
	StatusSince  string                  `json:"status_since"`
	Unobserved   *enums.ComparisonStatus `json:"unobserved,omitempty"`
	Url          string                  `json:"url"`
}
