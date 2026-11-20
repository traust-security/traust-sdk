// Code generated from traust-contracts 15bded2e3e4b6c70bd1e20ef8df50c7b66b5ebc1. DO NOT EDIT.

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
