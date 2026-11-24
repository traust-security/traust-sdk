// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

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
