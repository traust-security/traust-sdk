// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

package types

type CodeFacts struct {
	Artifact   interface{} `json:"artifact"`
	Interfaces interface{} `json:"interfaces,omitempty"`
	Languages  interface{} `json:"languages"`
	Packages   interface{} `json:"packages,omitempty"`
	Symbols    interface{} `json:"symbols"`
}

type CodeFactsPackages struct {
	Exports []CodeFactsPackagesExportsItem `json:"exports"`
	Gaps    []CodeFactsPackagesGapsItem    `json:"gaps"`
	Imports []CodeFactsPackagesImportsItem `json:"imports"`
	Modules []CodeFactsPackagesModulesItem `json:"modules"`
}

type CodeFactsPackagesExportsItem struct {
	Kind     interface{} `json:"kind,omitempty"`
	Language string      `json:"language"`
	Location Location    `json:"location"`
	Name     string      `json:"name"`
	Package  string      `json:"package"`
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
	Language string   `json:"language"`
	Location Location `json:"location"`
	Package  string   `json:"package"`
}

type CodeFactsPackagesModulesItem struct {
	Ecosystem string   `json:"ecosystem"`
	Location  Location `json:"location"`
	Name      string   `json:"name"`
}

type Language struct {
	Blanks   interface{} `json:"blanks"`
	Code     int         `json:"code"`
	Comments interface{} `json:"comments"`
	Files    interface{} `json:"files"`
	Name     string      `json:"name"`
}

type Location struct {
	Hashes []LocationHashesItem `json:"hashes,omitempty"`
	Line   interface{}          `json:"line,omitempty"`
	Path   string               `json:"path"`
	Symbol interface{}          `json:"symbol,omitempty"`
}

type LocationHashesItem struct {
	Algorithm string `json:"algorithm"`
	Digest    string `json:"digest"`
}

type Symbol struct {
	Kind     interface{} `json:"kind"`
	Language interface{} `json:"language"`
	Line     int         `json:"line"`
	Name     string      `json:"name"`
	Path     string      `json:"path"`
}
