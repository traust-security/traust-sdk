// Code generated from traust-contracts 15bded2e3e4b6c70bd1e20ef8df50c7b66b5ebc1. DO NOT EDIT.

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
