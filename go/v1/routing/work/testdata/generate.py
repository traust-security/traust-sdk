#!/usr/bin/env python3
"""Regenerate the offline parity fixture from the pinned Python source.

Usage: python3 testdata/generate.py /path/to/traust/src/traust/cli/build_rescan_worklist.py

The SHA-256 check is intentional: a policy update requires reviewing both the Go
port and this pin. Only the named pure functions/constants are evaluated; the CLI
module is never imported and its I/O/deployment dependencies are not needed.
"""

import argparse
import ast
import hashlib
import json
from pathlib import Path


REVISION = "0f95c95f48d36509b70234d18b500c0b738e23e8"
SOURCE_SHA256 = "3e4a34ead553f1ddafacb658336fdf6c579054ee9339b003ae4eab3e848a3525"
SOURCE_PATH = "src/traust/cli/build_rescan_worklist.py"
NAMES = {
    "CHURN_FULL_LINES", "CHURN_FULL_RATIO", "SENSITIVE_MIN_LINES",
    "TIER_CEILING_DAYS", "TIGHTENED_CEILING_DAYS", "P0_MIN_LIVE",
    "_rule3_lane", "_over_ceiling", "DECISION_TABLE", "decide", "risk_tier",
}


def reference(source):
    raw = source.read_bytes()
    if hashlib.sha256(raw).hexdigest() != SOURCE_SHA256:
        raise ValueError(f"source does not match pinned Traust revision {REVISION}")
    selected, found = [], set()
    for node in ast.parse(raw).body:
        names = set()
        if isinstance(node, ast.FunctionDef):
            names.add(node.name)
        elif isinstance(node, ast.Assign):
            names = {n.id for n in node.targets if isinstance(n, ast.Name)}
        elif isinstance(node, ast.AnnAssign) and isinstance(node.target, ast.Name):
            names.add(node.target.id)
        if names & NAMES:
            selected.append(node)
            found.update(names & NAMES)
    if found != NAMES:
        raise ValueError(f"missing reference definitions: {NAMES - found}")
    namespace = {}
    exec(compile(ast.Module(body=selected, type_ignores=[]), str(source), "exec"), namespace)
    return namespace


def build_fixture(ref):
    table = []

    def case(name, **overrides):
        ctx = {
            "tier": "P1", "exposure": "", "C": 0, "R": None,
            "S": False, "S_lines": 0, "deps_only": False,
            "audit_age_days": 30, "changed": False, "ahead_by": None,
        }
        ctx.update(overrides)
        rule, lane, reason = ref["decide"](ctx)
        table.append({"name": name, "input": ctx,
                      "want": {"Rule": rule, "Lane": lane, "Reason": reason}})

    for lines in (0, 1, 7999, 8000, 8001):
        case(f"churn-lines-{lines}", C=lines)
    for ratio in (None, 0, 0.099, 0.10, 0.101, 1.2):
        case(f"churn-ratio-{ratio}", C=500, R=ratio)
    for tier in ("P0", "P1", "P2", "P3"):
        for lines in (199, 200):
            case(f"sensitive-{tier}-{lines}", tier=tier, C=500, S=True, S_lines=lines)
    case("sensitive-flag-false", C=500, S=False, S_lines=500)
    case("dependencies-only", C=50, deps_only=True)
    case("churn-before-sensitive", C=9000, S=True, S_lines=300)
    case("churn-before-dependencies", C=9000, deps_only=True)
    case("sensitive-before-dependencies", C=300, S=True, S_lines=300, deps_only=True)
    case("dependencies-before-age", C=50, deps_only=True, audit_age_days=9999)
    case("age-before-quarterly", C=50, audit_age_days=9999)
    case("quarterly-before-push-noise", C=50, changed=True, ahead_by=0)
    for tier in ("P0", "P1", "P2", "P3"):
        for exposure in ("", "public-external", "private-external", "public-internal", "private-internal"):
            ceilings = (ref["TIGHTENED_CEILING_DAYS"] if exposure.endswith("-external")
                        else ref["TIER_CEILING_DAYS"])
            ceiling = ceilings.get(tier, 9999)
            for age in (None, ceiling, ceiling + 1):
                case(f"age-{tier}-{exposure or 'unspecified'}-{age}",
                     tier=tier, exposure=exposure, audit_age_days=age)
    for changed in (False, True):
        for ahead in (None, 0, 1):
            case(f"push-{changed}-ahead-{ahead}", changed=changed, ahead_by=ahead)

    risk = []
    for count in (0, 1, 4, 5, 6):
        for archived in (False, True):
            for dormant in (False, True):
                risk.append({
                    "name": f"risk-{count}-archived-{archived}-dormant-{dormant}",
                    "input": {"LiveCriticalHigh": count, "Archived": archived, "Dormant": dormant},
                    "want": ref["risk_tier"](count, archived, dormant),
                })
    return {"source": {"repo": "traust-security/traust", "revision": REVISION,
                       "path": SOURCE_PATH, "sha256": SOURCE_SHA256},
            "table": table, "risk": risk}


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("source", type=Path)
    args = parser.parse_args()
    fixture = build_fixture(reference(args.source))
    # One case per line keeps a policy update's diff easy to review.
    lines = ["{", '  "source": ' + json.dumps(fixture["source"]) + ","]
    for index, key in enumerate(("table", "risk")):
        lines.append(f'  "{key}": [')
        lines.append(",\n".join("    " + json.dumps(case) for case in fixture[key]))
        lines.append("  ]" + ("," if index == 0 else ""))
    lines.append("}")
    Path(__file__).with_name("python_decisions.json").write_text("\n".join(lines) + "\n")
    print(f"Generated {len(fixture['table'])} table and {len(fixture['risk'])} risk cases")


if __name__ == "__main__":
    main()
