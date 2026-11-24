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
            if set(kept[0]) == {"$ref"}:
                # a bare $ref: 3.0 ignores nullable as its sibling, so the null
                # would silently vanish (a required nullable enum decoded as its
                # zero value) — wrap the ref so the flag survives
                node["allOf"] = [kept[0]]
            else:
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


def normalize_file_downloads(spec: dict) -> None:
    """Treat multi-format file downloads (e.g. GET /v1/reports/{order_id}/file) as raw bytes
    rather than JSON strings: core declares application/json alongside text/* with schema
    {type: string}, which would make oapi-codegen try to unmarshal JSON objects into *string."""
    for path in spec.get("paths", {}).values():
        for operation in path.values():
            if not isinstance(operation, dict):
                continue
            ok = operation.get("responses", {}).get("200")
            if not isinstance(ok, dict):
                continue
            content = ok.get("content")
            if (
                isinstance(content, dict)
                and len(content) > 1
                and all(
                    isinstance(media, dict) and media.get("schema") == {"type": "string"}
                    for media in content.values()
                )
            ):
                ok["content"] = {
                    "application/octet-stream": {"schema": {"type": "string", "format": "binary"}}
                }


# Enums whose constants carry their type name, so a value shared with another enum never renames
# existing constants (SourceType beside SourceChannel; RunOrderReason's not_found, bad_request;
# FindingSort's severity beside the Severity type itself; ExecutivePanel's pqc beside the pqc
# view name; ReportFormat's bare html and csv, before any other enum takes them; Ecosystem's,
# which lost their prefix when core's enums stopped sharing a value with it).
NAMESPACED = (
    "Ecosystem",
    "SourceType",
    "RunOrderReason",
    "FindingSort",
    "FindingExposure",
    "SortDirection",
    "ExecutivePanel",
    "ReportFormat",
)


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
    normalize_file_downloads(spec)
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
