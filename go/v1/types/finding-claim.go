// Code generated from traust-contracts c2f6e6fb12270f8b3c8703da7bf39bd01776e4b1. DO NOT EDIT.

package types

type FindingClaim struct {
	AssessmentId  string            `json:"assessment_id"`
	Base          FindingClaimBase  `json:"base"`
	Claim         FindingClaimClaim `json:"claim"`
	FindingId     string            `json:"finding_id"`
	Label         string            `json:"label"`
	RepositoryId  string            `json:"repository_id"`
	SchemaVersion interface{}       `json:"schema_version"`
}

type FindingClaimBase struct {
	ClaimDecisionId interface{} `json:"claim_decision_id"`
	Fingerprint     string      `json:"fingerprint"`
	FingerprintAlgo string      `json:"fingerprint_algo"`
	OccurrenceIds   []string    `json:"occurrence_ids"`
	Title           string      `json:"title"`
}

type FindingClaimClaim struct {
	Category        interface{}                      `json:"category"`
	Cwes            []string                         `json:"cwes"`
	Description     string                           `json:"description"`
	Fingerprint     string                           `json:"fingerprint"`
	FingerprintAlgo string                           `json:"fingerprint_algo"`
	Locations       []FindingClaimClaimLocationsItem `json:"locations"`
	Title           string                           `json:"title"`
}

type FindingClaimClaimLocationsItem struct {
	LineEnd   interface{} `json:"line_end"`
	LineStart interface{} `json:"line_start"`
	Path      string      `json:"path"`
	Symbol    interface{} `json:"symbol"`
}
