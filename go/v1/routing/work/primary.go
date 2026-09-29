package work

// EventLane mirrors Python's event_lane. The bool distinguishes an absent lane
// (honored without queuing, or an unknown source) from a selected lane. An empty
// tier represents an event whose target is not in the audited population.
func EventLane(source EventSource, tier RiskTier) (Lane, bool) {
	if source == EventExternalReport {
		return LaneFullAuditValidate, true
	}
	if source == EventMethodology {
		if tier == RiskP0 {
			return LaneFullAudit, true
		}
		return "", false
	}
	if source == EventCVE {
		return LaneImpact, true
	}
	if source == EventRelease {
		return LaneReleasePassthrough, true
	}
	return "", false
}

// DecidePrimary composes the primary event loop, comparison-status guards,
// existing change table, and never-audited bootstrap policy from the pinned
// Python worklist builder. It preserves multiple event rows in input order;
// only queued events suppress ordinary table selection. Events for an unaudited
// inventory repository are honored as unknown audit targets, then bootstrap is
// evaluated separately. It does not resolve identities, collect facts, consume
// events, or authorize dispatch. Invalid input returns an empty result and error.
func DecidePrimary(in PrimaryInput) (PrimaryResult, error) {
	tier, table, err := primaryInputs(in)
	if err != nil {
		return PrimaryResult{}, err
	}
	result := PrimaryResult{
		Decisions:     []PrimaryDecision{},
		EventsHonored: []HonoredEvent{},
	}
	for i, event := range in.Events {
		if event.Consumed || !supportedEventSource(event.Source) {
			continue
		}
		lane, selected := EventLane(event.Source, tier)
		honored := HonoredEvent{EventIndex: i, Queued: selected && in.Audit != nil}
		if in.Audit == nil {
			honored.Disposition = "repo not in the HEAD code-audit population — dispatch manually"
		} else if !selected {
			honored.Disposition = "honored, not queued (methodology events queue P0 only; P1+ catch up via the rule-4 ceiling)"
		} else {
			reason := "event:" + string(event.Source)
			if event.Source == EventRelease {
				reason += " — passthrough: route to the branch/container/rpm audit lane, not this table"
			}
			result.Decisions = append(result.Decisions, PrimaryDecision{
				// Python event reasons deliberately omit the rule-1 prefix.
				Decision:   Decision{Rule: "rule-1", Lane: lane, Reason: reason},
				RiskTier:   tier,
				Status:     in.Audit.Status,
				EventIndex: &i,
			})
		}
		result.EventsHonored = append(result.EventsHonored, honored)
	}
	if len(result.Decisions) > 0 {
		return result, nil
	}
	if in.Audit != nil {
		selected, err := decideComparison(*in.Audit, table)
		if err != nil {
			return PrimaryResult{}, err
		}
		result.Decisions = append(result.Decisions, PrimaryDecision{
			Decision: selected, RiskTier: tier, Exposure: in.Audit.Exposure, Status: in.Audit.Status,
		})
	} else if in.Inventory != nil && !in.DisableBootstrap {
		result.Decisions = append(result.Decisions, neverAuditedRow(*in.Inventory))
	}
	return result, nil
}

func supportedEventSource(source EventSource) bool {
	switch source {
	case EventExternalReport, EventMethodology, EventCVE, EventRelease:
		return true
	default:
		return false
	}
}

// decideComparison preserves the two pre-table guards in Python's main loop.
func decideComparison(audit AuditInput, table TableInput) (Decision, error) {
	if audit.Status == StatusQuotaDeferred {
		return decision("quota-deferred", LaneNone,
			"changed, stage-2 compare deferred to the next run (API quota)"), nil
	}
	if audit.Status == StatusNoPinnedSHA {
		if audit.PushChanged {
			return decision("no-pinned-sha", LaneFullAudit,
				"changed since audit but churn unmeasurable — full audit re-establishes the anchor (self-healing)"), nil
		}
		return decision("no-pinned-sha", LaneNone,
			"no push signal; anchor is re-established at the next audit"), nil
	}
	if audit.Status == StatusOK && audit.PushChanged && audit.Change == nil {
		return Decision{}, &InputError{Field: "Audit.Change", Detail: "required for a successful changed-repository comparison"}
	}
	// For explicit non-OK observations, Python still permits the age rule to
	// fire with unavailable churn. The output retains that observation status.
	return EvaluateTable(table)
}

// neverAuditedRow ports the row policy of never_audited_rows using resolved
// population membership and designation. It performs no fleet sorting.
func neverAuditedRow(in InventoryInput) PrimaryDecision {
	exposure := ExposurePrivateInternal
	if in.Designation == DesignationExternal {
		exposure = ExposurePrivateExternal
	}
	return PrimaryDecision{
		Decision: Decision{
			Rule: "rule-1-bootstrap", Lane: LaneFullAudit,
			Reason: "rule-1: never audited — inventory repo without a baseline; full audit + threat-model bootstrap (self-heals into change detection, the graph, and dependency watch)",
		},
		RiskTier: RiskP2, Exposure: exposure, Status: StatusNeverAudited,
	}
}

func primaryInputs(in PrimaryInput) (RiskTier, TableInput, error) {
	if in.Inventory != nil {
		switch in.Inventory.Designation {
		case DesignationUnspecified, DesignationExternal, DesignationInternalTooling:
		default:
			return "", TableInput{}, &InputError{Field: "Inventory.Designation", Detail: "unknown exposure designation"}
		}
	}
	if in.Audit == nil {
		return "", TableInput{}, nil
	}
	audit := in.Audit
	switch audit.Status {
	case StatusOK, StatusQuotaDeferred, StatusNoPinnedSHA, StatusUnsupportedHost,
		StatusNoRepoURL, StatusNoCredentials, StatusUnreachable, StatusError,
		StatusBanSuspected, StatusNetworkSkipped:
	default:
		return "", TableInput{}, &InputError{Field: "Audit.Status", Detail: "must be a supported explicit observation status"}
	}
	tier, err := ClassifyRisk(audit.Risk)
	if err != nil {
		return "", TableInput{}, err
	}
	table := TableInput{
		RiskTier: tier, Exposure: audit.Exposure, AuditAgeDays: audit.AuditAgeDays, PushChanged: audit.PushChanged,
	}
	if change := audit.Change; change != nil {
		table.ChangedLines = change.ChangedLines
		table.CoverageChangeRatio = change.CoverageChangeRatio
		table.SensitiveChange = change.SensitiveChange
		table.SensitiveChangedLines = change.SensitiveChangedLines
		table.DependenciesOnly = change.DependenciesOnly
		table.CommitsAhead = change.CommitsAhead
	}
	if err := validateInput(table); err != nil {
		return "", TableInput{}, err
	}
	return tier, table, nil
}
