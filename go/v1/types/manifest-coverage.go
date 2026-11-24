// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

package types

import (
	"github.com/traust-security/traust-sdk/go/v1/enums"
)

type ManifestCoverage struct {
	Artifact     interface{}                                        `json:"artifact"`
	AssessmentId string                                             `json:"assessment_id"`
	SbomId       string                                             `json:"sbom_id"`
	Tool         ManifestCoverageTool                               `json:"tool"`
	Unresolved   map[string]map[string]enums.ManifestCoverageReason `json:"unresolved"`
}

type ManifestCoverageTool struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
