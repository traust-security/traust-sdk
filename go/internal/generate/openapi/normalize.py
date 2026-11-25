#!/usr/bin/env python3
# Workaround: rewrites core's OpenAPI into a client-ready 3.0 shape. Delete when core publishes
# a client-ready spec.
"""Normalize core's OpenAPI 3.1 document to 3.0 for oapi-codegen.

kin-openapi (behind oapi-codegen v2) reads the 3.0 shapes. The 3.1-isms in
core's document: exclusiveMinimum/exclusiveMaximum as numbers instead of the
3.0 boolean flag; pydantic's Optional as anyOf: [schema, {type: "null"}] instead
of 3.0's nullable: true; type as a list (3.1's union, e.g. ["integer",
"null"]) instead of 3.0's single type, rewritten to the one non-null type or a
oneOf of {type: t} when several, with nullable: true when "null" was a member;
and const as a single-value enum. The if/then/else conditionals (3.1-only,
assertions with no 3.0 shape) are dropped: oapi-codegen silently renders a
schema that carries one as interface{}, losing its properties. This rewrites
those forms, namespaces the enums in NAMESPACED, and gives Core's error
responses one JSON Problem schema so generated parsers accept problem+json.
Wire values are unchanged; the dropped conditionals were validation assertions
no Go client enforced anyway. Enums under the artifact contracts get explicit
path-prefixed constant names (mark_artifact_enums), so the values they share
with core's API enums never rename the API's constants.

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


def _rewrite_type_list(node: dict) -> None:
    """3.1's type list (its union) as 3.0's single type, or a oneOf of one-member
    schemas when several remain; nullable: true when "null" was a member."""
    types = node.get("type")
    if not isinstance(types, list) or not types:
        return
    non_null = [t for t in types if t != "null"]
    if len(non_null) == 1:
        node["type"] = non_null[0]
    elif non_null:
        del node["type"]
        node["oneOf"] = [{"type": t} for t in non_null]
    else:
        del node["type"]
    if "null" in types:
        node["nullable"] = True


def _rewrite_const(node: dict) -> None:
    """3.1's const as the single-value enum 3.0 reads."""
    if "const" in node:
        node["enum"] = [node.pop("const")]


def _drop_conditionals(node: dict) -> None:
    """Drop if/then/else: 3.1-only assertions no 3.0 tool reads, whose presence
    makes oapi-codegen render the whole schema as interface{}."""
    for key in ("if", "then", "else"):
        node.pop(key, None)
    branches = node.get("allOf")
    if not isinstance(branches, list):
        return
    kept = [b for b in branches if not (isinstance(b, dict) and not set(b) - {"$comment", "if", "then", "else"})]
    if not kept:
        del node["allOf"]
    elif len(kept) < len(branches):
        node["allOf"] = kept


def normalize(node: object) -> None:
    if isinstance(node, dict):
        for key in ("exclusiveMinimum", "exclusiveMaximum"):
            value = node.get(key)
            if isinstance(value, (int, float)):
                bound = "minimum" if key == "exclusiveMinimum" else "maximum"
                node[bound] = value if bound in node else value
                node[key] = True
        _strip_null_branches(node)
        _rewrite_type_list(node)
        _rewrite_const(node)
        _drop_conditionals(node)
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


# The keys a schema's vocabulary hangs under: a path segment through them names no word of its own.
STRUCTURAL_KEYS = frozenset((
    "additionalProperties", "allOf", "anyOf", "contains", "items", "not",
    "oneOf", "patternProperties", "prefixItems", "properties", "propertyNames",
))


def _pascal(segment: str) -> str:
    return "".join(part[:1].upper() + part[1:] for part in segment.split("_") if part)


def _identifier(name: str) -> str:
    return "".join(rune for rune in name if rune.isalnum())


def _kind_schema(spec: dict) -> str:
    """The schema the Traust-Artifact-Kind header carries: the enum of the kinds themselves,
    whose bare value names (Report, Scanner, Triage…) would otherwise collide with the API's."""
    for path in spec.get("paths", {}).values():
        for operation in path.values():
            if not isinstance(operation, dict):
                continue
            for response in operation.get("responses", {}).values():
                if not isinstance(response, dict):
                    continue
                header = response.get("headers", {}).get("Traust-Artifact-Kind")
                if isinstance(header, dict):
                    ref = header.get("schema", {}).get("$ref", "")
                    if ref.startswith("#/components/schemas/"):
                        return ref.rsplit("/", 1)[-1]
    return ""


def mark_artifact_enums(spec: dict) -> None:
    """Name every enum under an artifact contract after its schema path, and the kind enum the
    Traust-Artifact-Kind header carries after itself. The contracts hoist enums that share
    values with core's API enums, and oapi-codegen breaks that collision by prefixing both
    sides' constants — which renames the API's (RefKind's Branch became RefKindBranch when the
    artifact metadata enums took "branch"). Explicit x-enum-varnames keep the artifact
    constants path-qualified, so adding schemas never renames an identifier a consumer
    already holds."""
    schemas = spec["components"]["schemas"]
    contracts = [name for name, schema in schemas.items()
                 if isinstance(schema, dict) and "x-traust-artifact-kind" in schema]

    def mark(prefix: list[str], node: object) -> None:
        if isinstance(node, dict):
            enum = node.get("enum")
            if isinstance(enum, list) and enum and "x-enum-varnames" not in node:
                node["x-enum-varnames"] = [_identifier("".join(prefix) + _camel(str(value))) for value in enum]
            for key, value in node.items():
                if key == "x-enum-varnames":
                    continue
                mark(prefix if key in STRUCTURAL_KEYS else prefix + [_pascal(str(key))], value)
        elif isinstance(node, list):
            for index, item in enumerate(node):
                mark(prefix + [str(index)], item)

    family = [name for name in schemas if any(name.startswith(contract) for contract in contracts)]
    if kind := _kind_schema(spec):
        family.append(kind)
    for name in family:
        mark([name], schemas[name])


def main(argv: list[str]) -> int:
    if len(argv) != 3:
        print(__doc__, file=sys.stderr)
        return 2
    with open(argv[1], encoding="utf-8") as handle:
        spec = json.load(handle)
    normalize(spec)
    normalize_problems(spec)
    mark_artifact_enums(spec)
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
