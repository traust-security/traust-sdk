// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

package types

import (
	"github.com/traust-security/traust-sdk/go/v1/enums"
)

type FuzzReport struct {
	Artifact interface{}          `json:"artifact"`
	Batch    *string              `json:"batch,omitempty"`
	Coverage interface{}          `json:"coverage"`
	Crashers []Crasher            `json:"crashers"`
	Gaps     []Gap                `json:"gaps"`
	Patterns []PatternObservation `json:"patterns,omitempty"`
	Target   Target               `json:"target"`
}

type Coverage struct {
	EdgesCovered interface{} `json:"edges_covered,omitempty"`
	EdgesTotal   interface{} `json:"edges_total,omitempty"`
	LinesCovered int         `json:"lines_covered"`
	LinesTotal   int         `json:"lines_total"`
	Tool         string      `json:"tool"`
}

type Crasher struct {
	Analysis         *CrasherAnalysis   `json:"analysis,omitempty"`
	Cvss             *ReportFindingCvss `json:"cvss,omitempty"`
	Cwes             []string           `json:"cwes"`
	FindingId        *string            `json:"finding_id,omitempty"`
	Harness          string             `json:"harness"`
	Id               string             `json:"id"`
	Location         *ReportLocation    `json:"location,omitempty"`
	Reproducer       Reproducer         `json:"reproducer"`
	Sanitizer        interface{}        `json:"sanitizer"`
	SeverityProposal enums.Severity     `json:"severity_proposal"`
	Signature        Signature          `json:"signature"`
	Title            string             `json:"title"`
}

type CrasherAnalysis struct {
	Fix                       *string  `json:"fix,omitempty"`
	Narrative                 *string  `json:"narrative,omitempty"`
	NarrativeReferences       []string `json:"narrative_references,omitempty"`
	NegativeControlReferences []string `json:"negative_control_references,omitempty"`
	Priority                  *string  `json:"priority,omitempty"`
}

type Gap struct {
	Harness *string             `json:"harness,omitempty"`
	Note    *string             `json:"note,omitempty"`
	Reason  enums.FuzzGapReason `json:"reason"`
	Surface string              `json:"surface"`
}

type Harness struct {
	CorpusEntries   interface{}          `json:"corpus_entries,omitempty"`
	DurationSeconds int                  `json:"duration_seconds"`
	Engine          enums.FuzzEngine     `json:"engine"`
	Execs           interface{}          `json:"execs"`
	Name            string               `json:"name"`
	Outcome         enums.FuzzRunOutcome `json:"outcome"`
	Package         *string              `json:"package,omitempty"`
	Path            string               `json:"path"`
}

type PatternObservation struct {
	Advisory   *string     `json:"advisory,omitempty"`
	Name       string      `json:"name"`
	Reference  *string     `json:"reference,omitempty"`
	Tested     bool        `json:"tested"`
	Vulnerable interface{} `json:"vulnerable"`
}

type Reproducer struct {
	Path   *string `json:"path,omitempty"`
	Sha256 string  `json:"sha256"`
	Uri    *string `json:"uri,omitempty"`
}

type Signature struct {
	CrashClass enums.FuzzCrashClass `json:"crash_class"`
	Frames     *int                 `json:"frames,omitempty"`
	StackHash  interface{}          `json:"stack_hash"`
}

type Target struct {
	Harnesses []Harness `json:"harnesses"`
	Id        string    `json:"id"`
	Language  string    `json:"language"`
	Module    *string   `json:"module,omitempty"`
}
