#!/usr/bin/env python3
"""Normalize core's OpenAPI 3.1 document to 3.0 for oapi-codegen.

kin-openapi (behind oapi-codegen v2) reads the 3.0 shapes. Two 3.1-isms
appear in core's document: exclusiveMinimum/exclusiveMaximum as numbers
instead of the 3.0 boolean flag, and pydantic's Optional as
anyOf: [schema, {type: "null"}] instead of 3.0's nullable: true. This
rewrites those forms, namespaces the enums in NAMESPACED, and gives Core's error
responses one JSON Problem schema so generated parsers accept problem+json.
Wire values and validation constraints are unchanged.

Usage: normalize.py <openapi.json> <openapi-3.0.json>
"""

from __future__ import annotations

import json
import sys


def _strip_null_branches(node: dict) -> None:
    """Drop {type: "null"} branches from anyOf/oneOf (pydantic's Optional);
    a single remaining branch collapses with nullable: true, several remain
    as a list marked nullable."""
    for key in ("anyOf", "oneOf"):
        branches = node.get(key)
        if not isinstance(branches, list):
            continue
        kept = [b for b in branches if not (isinstance(b, dict) and b.get("type") == "null")]
        if len(kept) == len(branches):
            continue
        if len(kept) == 1 and isinstance(kept[0], dict):
            del node[key]
            node.update(kept[0])
        else:
            node[key] = kept
        node["nullable"] = True


def normalize(node: object) -> None:
    if isinstance(node, dict):
        for key in ("exclusiveMinimum", "exclusiveMaximum"):
            value = node.get(key)
            if isinstance(value, (int, float)):
                bound = "minimum" if key == "exclusiveMinimum" else "maximum"
                node[bound] = value if bound in node else value
                node[key] = True
        _strip_null_branches(node)
        for value in node.values():
            normalize(value)
    elif isinstance(node, list):
        for item in node:
            normalize(item)


def normalize_problems(spec: dict) -> None:
    """Use one JSON schema so oapi-codegen accepts problem+json, including charset parameters."""
    problem = {"schema": {"$ref": "#/components/schemas/Problem"}}
    for path in spec.get("paths", {}).values():
        for operation in path.values():
            if not isinstance(operation, dict):
                continue
            for status, response in operation.get("responses", {}).items():
                if str(status).isdigit() and 400 <= int(status) < 600:
                    response["content"] = {"application/json": problem}


# Enums whose constants carry their type name, so a value shared with another enum never renames
# existing constants (SourceType beside SourceChannel; RunOrderReason's not_found, bad_request;
# Validity's confirmed beside ReviewItemStatus and ValidationVerdict, whose naming oapi-codegen
# flipped when the impact schemas joined the spec).
NAMESPACED = ("SourceType", "RunOrderReason", "Validity")


def _camel(value: str) -> str:
    return "".join(part.title() for part in value.replace("-", "_").split("_"))


def main(argv: list[str]) -> int:
    if len(argv) != 3:
        print(__doc__, file=sys.stderr)
        return 2
    with open(argv[1], encoding="utf-8") as handle:
        spec = json.load(handle)
    normalize(spec)
    normalize_problems(spec)
    for name in NAMESPACED:
        schema = spec["components"]["schemas"].get(name)
        if schema is not None:
            schema["x-enum-varnames"] = [name + _camel(value) for value in schema["enum"]]
    spec["openapi"] = "3.0.3"
    with open(argv[2], "w", encoding="utf-8") as handle:
        json.dump(spec, handle, indent=2)
        handle.write("\n")
    return 0


if __name__ == "__main__":
    raise SystemExit(main(sys.argv))
