// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

package types

type RefutedRegister struct {
	Entries     []Entry  `json:"entries"`
	GeneratedAt string   `json:"generated_at"`
	Source      string   `json:"source"`
	Sources     []string `json:"sources,omitempty"`
}

type Entry struct {
	AssertedAt      string      `json:"asserted_at"`
	AssertedBy      string      `json:"asserted_by"`
	Category        interface{} `json:"category,omitempty"`
	ClaimedSeverity interface{} `json:"claimed_severity,omitempty"`
	EvidenceRefs    []string    `json:"evidence_refs"`
	ExclusionRule   interface{} `json:"exclusion_rule,omitempty"`
	File            interface{} `json:"file,omitempty"`
	FindingRef      string      `json:"finding_ref"`
	Line            interface{} `json:"line,omitempty"`
	Note            string      `json:"note"`
	RefuteReasons   []string    `json:"refute_reasons"`
	Source          *string     `json:"source,omitempty"`
	Tier            string      `json:"tier"`
	Title           string      `json:"title"`
	TriageId        *string     `json:"triage_id,omitempty"`
}
