// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

package types

import (
	"github.com/traust-security/traust-sdk/go/v1/enums"
)

type PqcVendorClaims struct {
	Artifact interface{} `json:"artifact"`
	Note     *string     `json:"note,omitempty"`
	Updated  string      `json:"updated"`
	Vendors  []Vendor    `json:"vendors"`
	Version  int         `json:"version"`
}

type Package struct {
	Ecosystem *enums.Ecosystem `json:"ecosystem,omitempty"`
	Name      string           `json:"name"`
}

type Source struct {
	Retrieved string  `json:"retrieved"`
	Title     *string `json:"title,omitempty"`
	Url       string  `json:"url"`
}

type Vendor struct {
	Claim        string                `json:"claim"`
	Class        enums.PqcVendorClass  `json:"class"`
	Id           string                `json:"id"`
	LastReviewed string                `json:"last_reviewed"`
	Name         string                `json:"name"`
	Packages     []Package             `json:"packages"`
	Reviewer     string                `json:"reviewer"`
	Sources      []Source              `json:"sources"`
	Status       enums.PqcVendorStatus `json:"status"`
	Watch        *string               `json:"watch,omitempty"`
}
