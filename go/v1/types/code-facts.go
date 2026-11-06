// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

package types

import (
	"github.com/traust-security/traust-sdk/go/v1/enums"
)

type CodeFacts struct {
	Artifact   interface{}               `json:"artifact"`
	Interfaces []CodeFactsInterfacesItem `json:"interfaces,omitempty"`
	Languages  []Language                `json:"languages"`
	Packages   *CodeFactsPackages        `json:"packages,omitempty"`
	Symbols    []Symbol                  `json:"symbols"`
}

type CodeFactsInterfacesItem struct {
	ApiGroup  string              `json:"api_group"`
	Kind      interface{}         `json:"kind,omitempty"`
	Location  CodeFactsLocation   `json:"location"`
	Name      interface{}         `json:"name,omitempty"`
	Resources []string            `json:"resources"`
	Role      enums.InterfaceRole `json:"role"`
	Verbs     []string            `json:"verbs"`
}

type CodeFactsLocation struct {
	Hashes []CodeFactsLocationHashesItem `json:"hashes,omitempty"`
	Line   interface{}                   `json:"line,omitempty"`
	Path   string                        `json:"path"`
	Symbol interface{}                   `json:"symbol,omitempty"`
}

type CodeFactsLocationHashesItem struct {
	Algorithm string `json:"algorithm"`
	Digest    string `json:"digest"`
}

type CodeFactsPackages struct {
	Exports []CodeFactsPackagesExportsItem `json:"exports"`
	Gaps    []CodeFactsPackagesGapsItem    `json:"gaps"`
	Imports []CodeFactsPackagesImportsItem `json:"imports"`
	Modules []CodeFactsPackagesModulesItem `json:"modules"`
}

type CodeFactsPackagesExportsItem struct {
	Kind     interface{}       `json:"kind,omitempty"`
	Language string            `json:"language"`
	Location CodeFactsLocation `json:"location"`
	Name     string            `json:"name"`
	Package  string            `json:"package"`
}

type CodeFactsPackagesGapsItem struct {
	Count    interface{} `json:"count,omitempty"`
	Evidence interface{} `json:"evidence,omitempty"`
	Path     interface{} `json:"path,omitempty"`
	Paths    []string    `json:"paths,omitempty"`
	Reason   string      `json:"reason"`
	System   string      `json:"system"`
}

type CodeFactsPackagesImportsItem struct {
	Language string            `json:"language"`
	Location CodeFactsLocation `json:"location"`
	Package  string            `json:"package"`
}

type CodeFactsPackagesModulesItem struct {
	Ecosystem string            `json:"ecosystem"`
	Location  CodeFactsLocation `json:"location"`
	Name      string            `json:"name"`
}

type Language struct {
	Blanks   interface{} `json:"blanks"`
	Code     int         `json:"code"`
	Comments interface{} `json:"comments"`
	Files    interface{} `json:"files"`
	Name     string      `json:"name"`
}

type Symbol struct {
	Kind     interface{} `json:"kind"`
	Language interface{} `json:"language"`
	Line     int         `json:"line"`
	Name     string      `json:"name"`
	Path     string      `json:"path"`
}
