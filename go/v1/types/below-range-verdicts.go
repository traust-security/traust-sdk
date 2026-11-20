// Code generated from traust-contracts 15bded2e3e4b6c70bd1e20ef8df50c7b66b5ebc1. DO NOT EDIT.

package types

import (
	"github.com/traust-security/traust-sdk/go/v1/enums"
)

type BelowRangeVerdicts struct {
	Artifact            interface{}                   `json:"artifact"`
	Cases               []BelowRangeVerdictsCasesItem `json:"cases"`
	CoverageGaps        []CoverageGap                 `json:"coverage_gaps"`
	PacketsAssessmentId string                        `json:"packets_assessment_id"`
}

type BelowRangeVerdictsCasesItem struct {
	Case    string                                   `json:"case"`
	Judge   *BelowRangeVerdictsCasesItemJudge        `json:"judge,omitempty"`
	Reviews []BelowRangeVerdictsCasesItemReviewsItem `json:"reviews"`
	Verdict enums.BelowRangeVerdict                  `json:"verdict"`
}

type BelowRangeVerdictsCasesItemJudge struct {
	Model   string                  `json:"model"`
	Ruling  string                  `json:"ruling"`
	Settle  *string                 `json:"settle,omitempty"`
	Verdict enums.BelowRangeVerdict `json:"verdict"`
}

type BelowRangeVerdictsCasesItemReviewsItem struct {
	Evidence string                  `json:"evidence"`
	Model    string                  `json:"model"`
	Verdict  enums.BelowRangeVerdict `json:"verdict"`
}
