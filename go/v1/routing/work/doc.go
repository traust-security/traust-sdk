// Package work implements the deterministic risk classification and
// change-decision table used by Traust's rescan worklist builder.
//
// EvaluateTable is one stage of work routing, not a complete worklist builder.
// Callers supply already-collected change measurements for an audited repository.
// They must handle bootstrap, event precedence, unavailable or deferred comparisons,
// and missing baseline commits before using this table. Additive IaC/threat-model
// lanes, fleet ordering, quarterly drain selection, and budget/model decisions
// remain outside this package. A returned lane is a classification, not permission
// to dispatch a job; in particular, diff-scan-quarterly enters a deferred pool.
//
// The rules and their order match traust.cli.build_rescan_worklist at Traust commit
// 0f95c95f48d36509b70234d18b500c0b738e23e8. Policy thresholds are fixed to that
// implementation. This package performs no I/O, reads no deployment configuration,
// and neither schedules nor dispatches work.
package work
