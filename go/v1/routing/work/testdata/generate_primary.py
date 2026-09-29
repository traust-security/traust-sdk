#!/usr/bin/env python3
"""Generate primary-selection fixtures from the pinned Python implementation.

Usage: python3 testdata/generate_primary.py /path/to/build_rescan_worklist.py

Runs the original event/table loops and never_audited_rows, with temporary event
and graph files. Companion stages are disabled/excluded. Identity resolution is
outside the per-repository Go API: all reference URLs here are already canonical.
No production database, network, shell command, or CLI module import is used.
"""

import argparse
import ast
import datetime
import hashlib
import json
from pathlib import Path
import sqlite3
import tempfile
from types import SimpleNamespace

from generate import REVISION, SOURCE_PATH, SOURCE_SHA256, reference


URL = "https://github.com/example/repository"
EXTRA_NAMES = {
    "EVENT_SOURCES", "event_lane", "load_events", "_parse_date",
    "RELEASE_REMODEL_CHANGES", "EXPOSURE_ORDER", "lookup_designation",
    "never_audited_rows",
}


def assigned(node, name):
    return isinstance(node, ast.Assign) and any(
        isinstance(part, ast.Name) and part.id == name
        for target in node.targets for part in ast.walk(target)
    )


def primary_reference(source):
    raw = source.read_bytes()
    if hashlib.sha256(raw).hexdigest() != SOURCE_SHA256:
        raise ValueError(f"source does not match pinned Traust revision {REVISION}")
    tree = ast.parse(raw)
    ns = reference(source)
    ns.update(Path=Path, sqlite3=sqlite3, json=json, _dt=datetime,
              normalize_repo_url=lambda value: value)
    selected, found = [], set()
    for node in tree.body:
        names = {node.name} if isinstance(node, ast.FunctionDef) else set()
        if isinstance(node, ast.Assign):
            names = {n.id for n in node.targets if isinstance(n, ast.Name)}
        if names & EXTRA_NAMES:
            selected.append(node)
            found.update(names & EXTRA_NAMES)
    if found != EXTRA_NAMES:
        raise ValueError(f"missing reference definitions: {EXTRA_NAMES - found}")
    exec(compile(ast.Module(body=selected, type_ignores=[]), str(source), "exec"), ns)
    main = next(n for n in tree.body if isinstance(n, ast.FunctionDef) and n.name == "main")
    start = next(i for i, n in enumerate(main.body) if assigned(n, "event_rows"))
    stop = next(i for i, n in enumerate(main.body) if assigned(n, "cc_urls"))
    # The unchanged source hash plus these anchors fix the exact original loops.
    loops = compile(ast.Module(body=main.body[start:stop], type_ignores=[]), str(source), "exec")
    return ns, loops


def run_reference(reference_ns, loops, supplied):
    ns = reference_ns.copy()
    audit = supplied.get("Audit")
    entries = []
    if audit is not None:
        risk = audit.get("Risk", {})
        entry = {
            "repo_key": "example/repository", "repo_url": URL,
            "live_crit_high": risk.get("LiveCriticalHigh", 0),
            "archived": risk.get("Archived", False), "dormant": risk.get("Dormant", False),
            "status": audit["Status"], "exposure": audit.get("Exposure", ""),
            "audit_age_days": audit.get("AuditAgeDays"),
            "changed": audit.get("PushChanged", False),
        }
        change = audit.get("Change")
        if change is not None:
            for go, py in (("ChangedLines", "C"), ("SensitiveChange", "S"),
                           ("SensitiveChangedLines", "S_lines"), ("DependenciesOnly", "deps_only")):
                entry[py] = change.get(go, False if go in ("SensitiveChange", "DependenciesOnly") else 0)
            entry["ahead_by"] = change.get("CommitsAhead")
            ratio = change.get("CoverageChangeRatio")
            if ratio:
                if not entry["C"]:
                    raise ValueError("positive reference ratio needs nonzero changed lines")
                entry["lines_reviewed"] = entry["C"] / ratio
        entries.append(entry)
    with tempfile.TemporaryDirectory() as directory:
        directory = Path(directory)
        event_path = directory / "events.jsonl"
        events = supplied.get("Events", [])
        event_path.write_text("".join(json.dumps({
            "repo": URL, "source": e["Source"], "consumed": e.get("Consumed", False),
        }) + "\n" for e in events))
        ns.update(entries=entries, events=ns["load_events"](event_path),
                  by_url={e["repo_url"]: e for e in entries},
                  by_key={e["repo_key"]: e for e in entries},
                  args=SimpleNamespace(no_threat_model=True), today=datetime.date(2026, 9, 28))
        exec(loops, ns)
        graph = directory / "graph.db"
        con = sqlite3.connect(graph)
        con.execute("CREATE TABLE nodes (id TEXT)")
        inventory = supplied.get("Inventory")
        if inventory is not None:
            con.execute("INSERT INTO nodes VALUES (?)", ("repo:" + URL.removeprefix("https://"),))
        con.commit()
        con.close()
        designation = inventory.get("Designation", "") if inventory is not None else ""
        bootstrap = [] if supplied.get("DisableBootstrap", False) else ns["never_audited_rows"](
            graph, {e["repo_url"] for e in entries}, [(URL, designation)] if designation else [],
        )
    kept_indexes = [i for i, e in enumerate(events)
                    if not e.get("Consumed", False) and e["Source"] in ns["EVENT_SOURCES"]]
    honored = [{"EventIndex": i, "Queued": e["queued"], "Disposition": e.get("disposition", "")}
               for i, e in zip(kept_indexes, ns["events_honored"], strict=True)]
    queued_indexes = iter(e["EventIndex"] for e in honored if e["Queued"])
    rows = ns["event_rows"] + ns["table_rows"] + bootstrap
    return {
        "Decisions": [{"Rule": r["rule"], "Lane": r["lane"], "Reason": r["reason"],
                       "RiskTier": r["tier"], "Exposure": r.get("exposure", ""),
                       "Status": r["status"],
                       "EventIndex": next(queued_indexes) if "event_source" in r else None}
                      for r in rows],
        "EventsHonored": honored,
    }


def build_fixture(ns, loops):
    cases = []

    def case(name, **supplied):
        cases.append({"name": name, "input": supplied, "want": run_reference(ns, loops, supplied)})

    sources = (*ns["EVENT_SOURCES"], "unknown", "")
    risks = {"P0": {"LiveCriticalHigh": 5}, "P1": {"LiveCriticalHigh": 1},
             "P2": {}, "P3": {"Dormant": True}}
    for tier, risk in risks.items():
        for source in sources:
            for consumed in (False, True):
                case(f"event-{source or 'empty'}-{tier}-consumed-{consumed}",
                     Audit={"Risk": risk, "Status": "ok", "AuditAgeDays": 9999},
                     Inventory={}, Events=[{"Source": source, "Consumed": consumed}])
    for source in sources:
        for inventory in (None, {}, {"Designation": "external"}):
            case(f"unaudited-{source or 'empty'}-inventory-{inventory}",
                 Inventory=inventory, Events=[{"Source": source}])
    for designation in ("", "external", "internal-tooling"):
        for disabled in (False, True):
            case(f"bootstrap-{designation or 'unspecified'}-disabled-{disabled}",
                 Inventory={"Designation": designation}, DisableBootstrap=disabled,
                 Events=[{"Source": "release"}])
    case("empty")
    for source in ns["EVENT_SOURCES"]:
        case(f"event-before-deferred-{source}",
             Audit={"Risk": risks["P0"], "Status": "quota-deferred", "PushChanged": True},
             Events=[{"Source": source}])
    case("multiple-events-keep-order-and-original-indexes", Audit={"Status": "ok"}, Events=[
        {"Source": "cve", "Consumed": True}, {"Source": "unknown"},
        {"Source": "external-report"}, {"Source": "methodology"},
        {"Source": "cve"}, {"Source": "cve"}, {"Source": "release"},
    ])
    case("methodology-P1-falls-through-to-churn", Audit={
        "Risk": risks["P1"], "Status": "ok", "PushChanged": True,
        "Change": {"ChangedLines": 8000},
    }, Events=[{"Source": "methodology"}])
    statuses = ("ok", "quota-deferred", "no-pinned-sha", "unsupported-host", "no-repo-url",
                "no-credentials", "unreachable", "error", "ban-suspected", "network-skipped")
    for status in statuses:
        for pushed in (False, True):
            for age in (None, 270, 271):
                audit = {"Status": status, "PushChanged": pushed, "AuditAgeDays": age,
                         "Risk": risks["P1"]}
                if status == "ok" and pushed:
                    audit["Change"] = {"CommitsAhead": 0}
                case(f"status-{status}-pushed-{pushed}-age-{age}", Audit=audit, Inventory={})
    for status in ("ok", "quota-deferred", "no-pinned-sha"):
        for change in ({"ChangedLines": 8000}, {"ChangedLines": 1000, "CoverageChangeRatio": .1},
                       {"ChangedLines": 500, "SensitiveChange": True, "SensitiveChangedLines": 200},
                       {"ChangedLines": 10, "DependenciesOnly": True}, {"ChangedLines": 1}):
            case(f"status-{status}-metrics-{change}", Audit={
                "Status": status, "PushChanged": True, "Change": change, "Risk": risks["P2"],
            })
    for exposure in ("", "public-external", "private-external", "public-internal", "private-internal"):
        case(f"exposure-{exposure or 'unspecified'}-unreachable-age", Audit={
            "Status": "unreachable", "Exposure": exposure, "Risk": risks["P1"], "AuditAgeDays": 181,
        })
    event_lanes = [{"source": source, "tier": tier,
                    "lane": ns["event_lane"](source, tier or None)}
                   for source in sources for tier in ("", *risks)]
    return {"source": {"repo": "traust-security/traust", "revision": REVISION,
                       "path": SOURCE_PATH, "sha256": SOURCE_SHA256},
            "event_lanes": event_lanes, "primary": cases}


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("source", type=Path)
    args = parser.parse_args()
    fixture = build_fixture(*primary_reference(args.source))
    lines = ["{", '  "source": ' + json.dumps(fixture["source"]) + ","]
    for index, key in enumerate(("event_lanes", "primary")):
        lines.append(f'  "{key}": [')
        lines.append(",\n".join("    " + json.dumps(case) for case in fixture[key]))
        lines.append("  ]" + ("," if index == 0 else ""))
    lines.append("}")
    Path(__file__).with_name("python_primary.json").write_text("\n".join(lines) + "\n")
    print(f"Generated {len(fixture['event_lanes'])} event-lane and {len(fixture['primary'])} primary cases")


if __name__ == "__main__":
    main()
