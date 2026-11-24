// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

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
