// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

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
