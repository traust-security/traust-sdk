package work

// EventSource is an upstream rescan event source.
type EventSource string

const (
	EventExternalReport EventSource = "external-report"
	EventMethodology    EventSource = "methodology"
	EventCVE            EventSource = "cve"
	EventRelease        EventSource = "release"
)

const (
	LaneFullAuditValidate  Lane = "full-audit+validate"
	LaneImpact             Lane = "impact-lane"
	LaneReleasePassthrough Lane = "release-passthrough"
)

// RescanEvent contains the fields used by primary scan selection. The caller
// resolves events to this repository by canonical URL or audit/sibling key and
// retains their payloads. Result indexes refer to this original input slice.
// Consumed events and unknown sources are ignored, as in Python's load_events.
type RescanEvent struct {
	Source   EventSource
	Consumed bool
}

// ComparisonStatus preserves the upstream observation status. Only
// quota-deferred and no-pinned-sha bypass the ordinary decision table. Other
// failure statuses can still select an age-ceiling audit; they do not disappear
// from results or become evidence of a successful comparison.
type ComparisonStatus string

const (
	StatusOK              ComparisonStatus = "ok"
	StatusQuotaDeferred   ComparisonStatus = "quota-deferred"
	StatusNoPinnedSHA     ComparisonStatus = "no-pinned-sha"
	StatusUnsupportedHost ComparisonStatus = "unsupported-host"
	StatusNoRepoURL       ComparisonStatus = "no-repo-url"
	StatusNoCredentials   ComparisonStatus = "no-credentials"
	StatusUnreachable     ComparisonStatus = "unreachable"
	StatusError           ComparisonStatus = "error"
	StatusBanSuspected    ComparisonStatus = "ban-suspected"
	StatusNetworkSkipped  ComparisonStatus = "network-skipped"
	// StatusNeverAudited appears only on bootstrap output, not AuditInput.
	StatusNeverAudited ComparisonStatus = "never-audited"
)

// ChangeInput contains collected comparison measurements. Nil Change in
// AuditInput means no comparison is available; a pointer to an empty ChangeInput
// means a measured zero change. The caller computes first-party/sensitive file
// counts and the coverage ratio using the upstream measurement policy.
type ChangeInput struct {
	ChangedLines          int64
	CoverageChangeRatio   *float64
	SensitiveChange       bool
	SensitiveChangedLines int64
	DependenciesOnly      bool
	CommitsAhead          *int64
}

// AuditInput describes membership in the deduplicated HEAD code-audit
// population, even when the audit has no recoverable pinned SHA. The caller
// derives PushChanged and Dormant using the upstream date policy and supplies
// the current observation status explicitly.
type AuditInput struct {
	Risk         RiskInput
	Exposure     Exposure
	Status       ComparisonStatus
	PushChanged  bool
	AuditAgeDays *int64
	Change       *ChangeInput
}

// ExposureDesignation is the resolved operator designation for an inventory
// repository. It is not the repository's organizational ownership label.
type ExposureDesignation string

const (
	DesignationUnspecified     ExposureDesignation = ""
	DesignationExternal        ExposureDesignation = "external"
	DesignationInternalTooling ExposureDesignation = "internal-tooling"
)

// InventoryInput confirms that this repository is in the supplied inventory.
// The caller resolves designation URL/prefix rules before constructing it.
type InventoryInput struct {
	Designation ExposureDesignation
}

// PrimaryInput is one repository's already-correlated facts. Audit nil means
// absent from the HEAD code-audit population; Inventory nil means absent from
// the supplied inventory. These are known membership facts, not failed reads.
// Do not interpret an unavailable inventory/audit view as an empty population.
type PrimaryInput struct {
	Audit            *AuditInput
	Inventory        *InventoryInput
	Events           []RescanEvent
	DisableBootstrap bool
}

// PrimaryDecision retains the status that led to a decision. EventIndex is set
// only for an event-generated decision. Exposure on event rows is unspecified,
// matching Python's primary event rows; bootstrap rows carry the inferred band.
type PrimaryDecision struct {
	Decision
	RiskTier   RiskTier
	Exposure   Exposure
	Status     ComparisonStatus
	EventIndex *int
}

// HonoredEvent records an accepted event even when it queues no primary work.
// Queued means a routing row was selected, not that a job was dispatched.
// Neither this result nor DecidePrimary marks the source event consumed.
type HonoredEvent struct {
	EventIndex  int
	Queued      bool
	Disposition string
}

// PrimaryResult is primary scan selection before companion lanes, tripwire,
// refusal pre-routing, fleet ordering, budget/model gates, or dispatch. Preserve
// the original release events for the separate threat-model companion stage.
// A successful empty result can mean no audited or bootstrap-eligible target.
// Decisions may include LaneNone or LaneReleasePassthrough; do not dispatch every
// returned row as an immediate code audit.
type PrimaryResult struct {
	Decisions     []PrimaryDecision
	EventsHonored []HonoredEvent
}
