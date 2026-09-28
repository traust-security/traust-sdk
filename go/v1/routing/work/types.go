package work

import "fmt"

// RiskTier is a repository's live-finding risk band, not a model capability tier.
type RiskTier string

const (
	RiskP0 RiskTier = "P0"
	RiskP1 RiskTier = "P1"
	RiskP2 RiskTier = "P2"
	RiskP3 RiskTier = "P3"
)

// Exposure describes repository visibility and the impact boundary. Both external
// bands tighten audit-age ceilings; public visibility alone does not.
type Exposure string

const (
	// ExposureUnspecified preserves the Python table's absent-exposure behavior:
	// use the ordinary age ceiling. It does not infer that a repo is internal.
	ExposureUnspecified     Exposure = ""
	ExposurePublicExternal  Exposure = "public-external"
	ExposurePrivateExternal Exposure = "private-external"
	ExposurePublicInternal  Exposure = "public-internal"
	ExposurePrivateInternal Exposure = "private-internal"
)

// Lane is a change-table result. The full Traust worklist has additional lanes.
type Lane string

const (
	LaneNone              Lane = "none"
	LaneFullAudit         Lane = "full-audit"
	LaneDiffScan          Lane = "diff-scan"
	LaneDependencies      Lane = "deps-lane"
	LaneDiffScanQuarterly Lane = "diff-scan-quarterly"
)

// RiskInput describes live critical/high findings and repository lifecycle state.
// The caller obtains the live count from its authoritative findings view.
type RiskInput struct {
	LiveCriticalHigh int64
	Archived         bool
	Dormant          bool
}

// TableInput contains normalized measurements for the ordered change-decision
// table. Zero counts and false booleans mean measured zero/false, not unavailable
// data. Do not call EvaluateTable with zero-filled failed/deferred comparisons.
type TableInput struct {
	RiskTier RiskTier
	Exposure Exposure

	// ChangedLines is C: first-party lines changed since the audit baseline.
	ChangedLines int64
	// CoverageChangeRatio is R: ChangedLines / baseline lines reviewed. Nil
	// means the denominator is unavailable, so only absolute churn is used.
	// Ratios above one are valid when changes exceed the baseline size.
	CoverageChangeRatio *float64
	// SensitiveChange is S; SensitiveChangedLines is S_lines from the
	// upstream sensitive-file matcher. This package does not match file paths.
	SensitiveChange       bool
	SensitiveChangedLines int64
	DependenciesOnly      bool

	// AuditAgeDays is nil when unknown. That skips the age rule; it does not
	// represent a never-audited repository (bootstrap is a separate stage).
	AuditAgeDays *int64
	// PushChanged is the upstream push signal. CommitsAhead distinguishes
	// an unknown comparison (nil) from known zero commits (a pointer to zero).
	PushChanged  bool
	CommitsAhead *int64
}

// Decision identifies the first matching Python rule, its lane, and its original
// explanation. A zero Decision returned with an error must not be dispatched.
type Decision struct {
	Rule   string
	Lane   Lane
	Reason string
}

// InputError identifies an invalid field at the SDK boundary. Use errors.As to
// inspect Field without parsing the message. Invalid input produces no decision.
type InputError struct {
	Field  string
	Detail string
}

func (e *InputError) Error() string {
	return fmt.Sprintf("work: invalid %s: %s", e.Field, e.Detail)
}
