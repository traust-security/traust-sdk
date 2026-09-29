package work

import "math"

// Policy constants mirroring Traust's build_rescan_worklist at the pinned
// commit. Named to match the Python originals for parity review.
const (
	// P0MinLive is the minimum live critical/high finding count for P0 risk.
	P0MinLive = 5

	// ChurnFullLines is rule-2's absolute line-change threshold (C).
	ChurnFullLines int64 = 8000
	// ChurnFullRatio is rule-2's coverage-change ratio threshold (R).
	ChurnFullRatio = 0.10
	// SensitiveMinLines is rule-3's minimum sensitive-file changed lines.
	SensitiveMinLines int64 = 200

	// TierCeilingP0 / P1 / P2 are rule-4's ordinary audit-age ceilings (days).
	TierCeilingP0 int64 = 180
	TierCeilingP1 int64 = 270
	TierCeilingP2 int64 = 365

	// TightenedCeilingP1 / P2 apply to external-exposure repos.
	TightenedCeilingP1 int64 = 180
	TightenedCeilingP2 int64 = 270
)

// ClassifyRisk reproduces Traust's risk_tier rule. Live critical/high findings
// take precedence over archival/dormancy: P0MinLive or more means P0; a positive
// count below P0MinLive means P1.
// With no live critical/high findings, archived/dormant repos are P3, others P2.
func ClassifyRisk(in RiskInput) (RiskTier, error) {
	if in.LiveCriticalHigh < 0 {
		return "", &InputError{Field: "LiveCriticalHigh", Detail: "must be non-negative"}
	}
	switch {
	case in.LiveCriticalHigh >= P0MinLive:
		return RiskP0, nil
	case in.LiveCriticalHigh >= 1:
		return RiskP1, nil
	case in.Archived || in.Dormant:
		return RiskP3, nil
	default:
		return RiskP2, nil
	}
}

// EvaluateTable applies Traust's ordered rules 2, 3, 3b, 4, 5, 6, and 7 to
// normalized inputs. Event routing (rule 1) and other worklist stages are not
// evaluated here. The first matching rule wins; see the package documentation
// for the caller's responsibilities before and after this stage.
func EvaluateTable(in TableInput) (Decision, error) {
	if err := validateInput(in); err != nil {
		return Decision{}, err
	}

	switch {
	case in.ChangedLines >= ChurnFullLines || (in.CoverageChangeRatio != nil && *in.CoverageChangeRatio >= ChurnFullRatio):
		return decision("rule-2", LaneFullAudit,
			"coverage-map churn (R>=10% or C>=8000 first-party lines)"), nil
	case in.SensitiveChange && in.SensitiveChangedLines >= SensitiveMinLines:
		lane := LaneDiffScan
		if in.RiskTier == RiskP0 || in.RiskTier == RiskP1 {
			lane = LaneFullAudit
		}
		return decision("rule-3", lane, "sensitive-identifier files changed, >=200 lines"), nil
	case in.DependenciesOnly:
		return decision("rule-3b", LaneDependencies, "dependency manifests only (deterministic osv/impact lane)"), nil
	case overCeiling(in):
		// Keep the Python reason verbatim, including its ordinary-ceiling
		// summary; external exposure uses the tightened ceilings below.
		return decision("rule-4", LaneFullAudit,
			"audit age over tier ceiling (P0 180d / P1 270d / P2 365d)"), nil
	case in.ChangedLines > 0:
		return decision("rule-5", LaneDiffScanQuarterly, "below-threshold change (accumulates for the quarterly batch)"), nil
	case in.PushChanged && in.CommitsAhead != nil && *in.CommitsAhead == 0:
		return decision("rule-6", LaneNone, "pushed_at moved but 0 commits ahead (tag/branch noise)"), nil
	default:
		return decision("rule-7", LaneNone, "no qualifying change (dormant repos re-enter on push)"), nil
	}
}

func decision(rule string, lane Lane, reason string) Decision {
	return Decision{Rule: rule, Lane: lane, Reason: rule + ": " + reason}
}

func overCeiling(in TableInput) bool {
	if in.AuditAgeDays == nil || in.RiskTier == RiskP3 {
		return false
	}
	var ceiling int64
	switch in.RiskTier {
	case RiskP0:
		ceiling = TierCeilingP0
	case RiskP1:
		ceiling = TierCeilingP1
	case RiskP2:
		ceiling = TierCeilingP2
	}
	if in.Exposure == ExposurePublicExternal || in.Exposure == ExposurePrivateExternal {
		switch in.RiskTier {
		case RiskP1:
			ceiling = TightenedCeilingP1
		case RiskP2:
			ceiling = TightenedCeilingP2
		}
	}
	return *in.AuditAgeDays > ceiling
}

func validateInput(in TableInput) error {
	switch in.RiskTier {
	case RiskP0, RiskP1, RiskP2, RiskP3:
	default:
		return &InputError{Field: "RiskTier", Detail: "must be P0, P1, P2, or P3"}
	}
	switch in.Exposure {
	case ExposureUnspecified, ExposurePublicExternal, ExposurePrivateExternal, ExposurePublicInternal, ExposurePrivateInternal:
	default:
		return &InputError{Field: "Exposure", Detail: "unknown exposure band"}
	}
	for _, field := range []struct {
		name  string
		value int64
	}{
		{"ChangedLines", in.ChangedLines},
		{"SensitiveChangedLines", in.SensitiveChangedLines},
	} {
		if field.value < 0 {
			return &InputError{Field: field.name, Detail: "must be non-negative"}
		}
	}
	for _, field := range []struct {
		name  string
		value *int64
	}{
		{"AuditAgeDays", in.AuditAgeDays},
		{"CommitsAhead", in.CommitsAhead},
	} {
		if field.value != nil && *field.value < 0 {
			return &InputError{Field: field.name, Detail: "must be non-negative when supplied"}
		}
	}
	if ratio := in.CoverageChangeRatio; ratio != nil && (math.IsNaN(*ratio) || math.IsInf(*ratio, 0) || *ratio < 0) {
		return &InputError{Field: "CoverageChangeRatio", Detail: "must be finite and non-negative when supplied"}
	}
	return nil
}
