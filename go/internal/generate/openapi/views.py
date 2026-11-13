#!/usr/bin/env python3
"""Turn core's Insight view contracts into OpenAPI 3.0 components for oapi-codegen.

Runs in core's environment at CORE_REF (`$(CORE_PYTHON)`), so the view schemas are the
`views_schema/v1` contracts core validates its views with: no second contracts pin. Each
view becomes `<View>Insight` and each of its `$defs` `<View>Insight<Def>`; references into
artifact schemas are inlined. A view composed from others (the executive summary) references
their contracts: a reference to another view, or to one of its `$defs`, becomes that view's
component, so a shape is generated once; a reference deeper into another view is inlined. Validation-only keywords (if/then/else, contains, not,
propertyNames, const) are dropped: the structs carry the shape, core enforces the rest.
Numbers are doubles so a measured value round-trips exactly.

Usage: views.py <views-openapi.json>
"""

from __future__ import annotations

import copy
import json
import sys
from typing import TypeAlias

# a JSON Schema tree as json.load returns it
Json: TypeAlias = dict[str, "Json"] | list["Json"] | str | int | float | bool | None

VALIDATION_ONLY = (
    "if",
    "then",
    "else",
    "contains",
    "minContains",
    "maxContains",
    "not",
    "propertyNames",
    "const",
    "$schema",
    "$id",
    "$comment",
)


def pascal(value: str) -> str:
    return "".join(part[:1].upper() + part[1:] for part in value.replace("-", "_").split("_"))


def pointer(document: Json, path: str) -> Json:
    node = document
    for part in filter(None, path.split("/")):
        node = node[part.replace("~1", "/").replace("~0", "~")]
    return node


VIEWS = "views_schema/v1/"


def view_reference(ref: str) -> tuple[str, str] | None:
    """The view and its pointer when `ref` names another view contract, else None."""
    target, _, path = ref.partition("#")
    if VIEWS not in target:
        return None
    return target.rsplit("/", 1)[-1].removesuffix(".schema.json"), path


def view_component(ref: str) -> str | None:
    """The component a reference to another view (or one of its `$defs`) names, else None."""
    found = view_reference(ref)
    if found is None:
        return None
    view, path = found
    parts = [p for p in path.split("/") if p]
    if not parts:
        return pascal(view) + "Insight"
    if len(parts) == 2 and parts[0] == "$defs":
        return pascal(view) + "Insight" + pascal(parts[1])
    return None


def inline_external(node: Json, document: Json) -> Json:
    """Every reference outside the view inlined from the schema it names, except a reference
    to another view or one of its `$defs`: that stays, to become the view's component."""
    from traust_core.v1 import contracts  # core's environment, at CORE_REF

    if isinstance(node, list):
        return [inline_external(item, document) for item in node]
    if not isinstance(node, dict):
        return node
    ref = node.get("$ref")
    if isinstance(ref, str) and not ref.startswith("#") and view_component(ref) is not None:
        return {key: inline_external(value, document) for key, value in node.items()}
    if isinstance(ref, str) and not ref.startswith("#"):
        target, _, path = ref.partition("#")
        view = view_reference(ref)
        name = target.rsplit("/", 1)[-1].removesuffix(".schema.json")
        external = contracts.view_schema(view[0]) if view else contracts.schema(name)
        resolved = inline_local(copy.deepcopy(pointer(external, path)), external)
        rest = {k: v for k, v in node.items() if k != "$ref"}
        return inline_external({**resolved, **rest}, document)
    return {key: inline_external(value, document) for key, value in node.items()}


def inline_local(node: Json, document: Json) -> Json:
    """An artifact-schema fragment's own `#/...` references inlined against that schema."""
    if isinstance(node, list):
        return [inline_local(item, document) for item in node]
    if not isinstance(node, dict):
        return node
    ref = node.get("$ref")
    if isinstance(ref, str) and ref.startswith("#"):
        resolved = inline_local(copy.deepcopy(pointer(document, ref[1:])), document)
        return {**resolved, **{k: v for k, v in node.items() if k != "$ref"}}
    return {key: inline_local(value, document) for key, value in node.items()}


def openapi(node: Json, prefix: str) -> Json:
    """One schema node in OpenAPI 3.0 form, its `#/$defs/x` references renamed."""
    if isinstance(node, list):
        return [openapi(item, prefix) for item in node]
    if not isinstance(node, dict):
        return node
    out: dict[str, Json] = {}
    for key, value in node.items():
        if key in VALIDATION_ONLY or key == "$defs":
            continue
        if key == "allOf":
            structural = [b for b in value if isinstance(b, dict) and set(b) - set(VALIDATION_ONLY)]
            if len(structural) == 1 and set(structural[0]) == {"$ref"}:
                out.update(openapi(structural[0], prefix))  # a def composed by reference
            elif structural:
                raise SystemExit(f"{prefix}: allOf carries structure; extend views.py")
            continue
        component = view_component(value) if key == "$ref" else None
        if component is not None:
            out["$ref"] = "#/components/schemas/" + component  # another view's, generated there
            continue
        if key == "$ref":
            if not value.startswith("#/$defs/") or value.count("/") != 2:
                raise SystemExit(f"{prefix}: unsupported reference {value}")
            out["$ref"] = "#/components/schemas/" + prefix + pascal(value.split("/")[-1])
            continue
        out[key] = openapi(value, prefix)
    kinds = out.get("type")
    if isinstance(kinds, list):
        concrete = [k for k in kinds if k != "null"]
        if len(concrete) != 1:
            raise SystemExit(f"{prefix}: a type list needs one non-null type: {kinds}")
        out["type"] = concrete[0]
        if "null" in kinds:
            out["nullable"] = True
    if "enum" in out and "type" not in out and all(isinstance(v, str) for v in out["enum"]):
        out["type"] = "string"  # an artifact schema's bare enum, inlined
    if isinstance(out.get("enum"), list) and None in out["enum"]:
        out["enum"] = [v for v in out["enum"] if v is not None]
        out["nullable"] = True
    branches = out.get("oneOf") or out.get("anyOf")
    if isinstance(branches, list):
        key = "oneOf" if "oneOf" in out else "anyOf"
        kept = [b for b in branches if b != {"type": "null"}]
        if len(kept) != len(branches):
            out["nullable"] = True
            if len(kept) == 1:
                del out[key]
                single = kept[0]
                if "$ref" in single:
                    # 3.0 ignores siblings of $ref, so nullability rides an allOf wrapper
                    out["allOf"] = [single]
                else:
                    out.update(single)
            else:
                out[key] = kept
    if out.get("type") == "number":
        out.setdefault("format", "double")
    return out


def refs(node: Json) -> list[str]:
    if isinstance(node, list):
        return [r for item in node for r in refs(item)]
    if not isinstance(node, dict):
        return []
    own = [node["$ref"]] if isinstance(node.get("$ref"), str) else []
    return own + [r for value in node.values() for r in refs(value)]


def components() -> tuple[dict[str, Json], dict[str, str]]:
    from traust_core.v1 import contracts  # core's environment, at CORE_REF
    from traust_core.v1.errors import ConfigError
    from traust_core.v1.models.insight import ViewName

    found: dict[str, Json] = {}
    views: dict[str, str] = {}
    for view in ViewName:
        try:
            schema = contracts.view_schema(view.value)
        except ConfigError:
            continue  # no contract yet: GetInsight still returns it untyped
        schema = inline_external(schema, schema)
        prefix = pascal(view.value) + "Insight"
        found[prefix] = openapi(schema, prefix)
        found.update(
            (prefix + pascal(name), openapi(definition, prefix))
            for name, definition in schema.get("$defs", {}).items()
        )
        views[prefix] = view.value
    # only definitions a view's shape reaches, across views (a composed view reaches its
    # components'): validation-only helpers (a contains target) go
    reached, pending = set(), list(views)
    while pending:
        name = pending.pop()
        if name not in reached:
            if name not in found:
                raise SystemExit(f"{name}: referenced, but no view contract defines it")
            reached.add(name)
            pending += [r.rsplit("/", 1)[-1] for r in refs(found[name])]
    return {name: found[name] for name in sorted(reached)}, views


def main(argv: list[str]) -> int:
    if len(argv) != 2:
        print(__doc__, file=sys.stderr)
        return 2
    from traust_core.v1 import contracts

    schemas, views = components()
    document = {
        "openapi": "3.0.3",
        "info": {"title": "traust-core Insight views", "version": contracts.version()},
        "paths": {},
        "components": {"schemas": schemas},
        "x-traust-insights": views,
    }
    with open(argv[1], "w", encoding="utf-8") as handle:
        json.dump(document, handle, indent=2, sort_keys=True)
        handle.write("\n")
    return 0


if __name__ == "__main__":
    raise SystemExit(main(sys.argv))
