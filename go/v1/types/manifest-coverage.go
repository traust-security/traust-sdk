// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

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
