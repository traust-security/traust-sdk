// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

package types

import (
	"github.com/traust-security/traust-sdk/go/v1/enums"
)

type AdrRegistry struct {
	Note      *string                    `json:"note,omitempty"`
	Registers []AdrRegistryRegistersItem `json:"registers"`
	Version   int                        `json:"version"`
}

type AdrRegistryRegistersItem struct {
	DeclaredStatus *enums.AdrDeclaredStatus `json:"declared_status,omitempty"`
	Governs        []string                 `json:"governs,omitempty"`
	Name           string                   `json:"name"`
	Note           *string                  `json:"note,omitempty"`
	Paths          []string                 `json:"paths"`
	Pin            string                   `json:"pin"`
	Repo           string                   `json:"repo"`
}
