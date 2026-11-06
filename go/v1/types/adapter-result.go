// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

package types

type AdapterResult struct {
	Findings   []AdapterFinding `json:"findings"`
	FocusAreas []string         `json:"focus_areas,omitempty"`
	Metadata   Metadata         `json:"metadata"`
	ScannedAt  string           `json:"scanned_at"`
	Summary    *Summary         `json:"summary,omitempty"`
	Target     string           `json:"target"`
}

type AdapterFinding = ReportFinding

type Metadata struct {
	ScannerVersion *string `json:"scanner_version,omitempty"`
	Tool           string  `json:"tool"`
}

type Summary struct {
	BySeverity map[string]int `json:"by_severity,omitempty"`
	Total      int            `json:"total"`
}
