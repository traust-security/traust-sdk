// Code generated from traust-contracts 369a2ab2c9b65225c6878eca75bada6a65df548f. DO NOT EDIT.

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
