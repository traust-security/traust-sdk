// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

package types

import (
	"github.com/traust-security/traust-sdk/go/v1/enums"
)

type CorpusRegistry struct {
	Note     *string   `json:"note,omitempty"`
	Subjects []Subject `json:"subjects"`
	Updated  *string   `json:"updated,omitempty"`
	Version  int       `json:"version"`
}

type Subject struct {
	BusinessUnit  string          `json:"business_unit"`
	IsBranchAudit *bool           `json:"is_branch_audit,omitempty"`
	Label         *string         `json:"label,omitempty"`
	Ownership     enums.Ownership `json:"ownership"`
	Product       *string         `json:"product,omitempty"`
	Ref           *string         `json:"ref,omitempty"`
	RefKind       *enums.RefKind  `json:"ref_kind,omitempty"`
	RepoUrl       *string         `json:"repo_url,omitempty"`
	SubjectId     string          `json:"subject_id"`
	Tree          string          `json:"tree"`
}
