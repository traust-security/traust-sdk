#!/usr/bin/env python3
"""The example's pack against the shipped security pack: `make check-drift`'s pack half.

Every asset the webapp's pack declares must be the shipped asset of the same name in traust at
TRAUST_REF, and the routing header (route_on, approvers) must match, so the box `make webapp`
starts plans and runs exactly the work a deployment's security pack does. The pack is then
loaded the way the harness loads it (the engine's own `Pack.load`: the contracts schema and
the engine's pack models), so a copy that would not validate fails here too.
"""

from __future__ import annotations

import argparse
import sys
from pathlib import Path
from typing import Any

import yaml


def _load(path: Path) -> dict[str, Any]:
    try:
        document = yaml.safe_load(path.read_text())
    except OSError as e:
        raise SystemExit(f"ERROR: cannot read {path}: {e}")
    except yaml.YAMLError as e:
        raise SystemExit(f"ERROR: {path}: YAML syntax error: {e}")
    if not isinstance(document, dict):
        raise SystemExit(f"ERROR: {path}: not a mapping")
    return document


def _problems(webapp: dict[str, Any], shipped: dict[str, Any]) -> list[str]:
    problems = []
    for key in ("route_on", "approvers"):
        if webapp.get(key) != shipped.get(key):
            problems.append(f"{key}: {webapp.get(key)!r} here, {shipped.get(key)!r} shipped")
    assets = webapp.get("assets")
    if not isinstance(assets, dict):
        return [*problems, "assets: not a mapping"]
    for name, asset in sorted(assets.items()):
        other = (shipped.get("assets") or {}).get(name)
        if other is None:
            problems.append(f"asset {name}: the shipped pack has no asset of this name")
        elif asset != other:
            fields = sorted(k for k in set(asset) | set(other) if asset.get(k) != other.get(k))
            problems.append(
                f"asset {name}: differs from the shipped asset in {', '.join(fields) or 'value'}"
            )
    return problems


def main() -> int:
    ap = argparse.ArgumentParser(description=__doc__)
    ap.add_argument("--webapp", type=Path, required=True, help="the example's pack.yaml")
    ap.add_argument(
        "--shipped", type=Path, required=True, help="traust's packs/security/assets.yaml"
    )
    args = ap.parse_args()
    problems = _problems(_load(args.webapp), _load(args.shipped))
    if problems:
        print(
            f"ERROR: {args.webapp} drifted from the shipped security pack at TRAUST_REF;"
            " copy the asset verbatim from it or drop the asset",
            file=sys.stderr,
        )
        for problem in problems:
            print(f"  {problem}", file=sys.stderr)
        return 1
    from traust_engine.ops.pack import Pack

    try:
        Pack.load(args.webapp)
    except Exception as e:  # the engine's ConfigError wording names the file and the rule
        print(f"ERROR: the engine cannot load {args.webapp}: {e}", file=sys.stderr)
        return 1
    return 0


if __name__ == "__main__":
    sys.exit(main())
