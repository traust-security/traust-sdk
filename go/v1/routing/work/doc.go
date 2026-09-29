// Package work implements primary scan selection, deterministic risk
// classification, and the change table used by Traust's rescan worklist builder.
//
// DecidePrimary evaluates one repository's already-correlated inventory, audit,
// comparison, and event facts. It composes event precedence, comparison-status
// guards, EvaluateTable, and never-audited bootstrap. Callers supply authoritative
// population membership and resolve event identities before calling it. Failed
// population reads must not be presented as known absence. EvaluateTable remains
// the lower-level stage for already-measured audited repositories.
//
// Additive IaC/threat-model lanes, tripwire and refusal pre-routing, fleet ordering,
// quarterly drain selection, and budget/model decisions remain outside this
// package. A returned lane is a classification, not permission to dispatch a job;
// diff-scan-quarterly enters a deferred pool and release-passthrough is a handoff.
// Events remain unconsumed; their original payloads must remain available for
// companion processing, including release-driven threat-model review.
//
// The rules and their order match traust.cli.build_rescan_worklist at Traust commit
// 0f95c95f48d36509b70234d18b500c0b738e23e8. Policy thresholds are fixed to that
// implementation. This package performs no I/O, reads no deployment configuration,
// and neither schedules nor dispatches work.
package work
