#!/usr/bin/env python3
"""Lean gate: the SDK gets smaller or stays the same size.

Compares HEAD against a base and fails on three kinds of growth. The base is
LEAN_BASE when set, else origin/development; the diff starts at that ref's
merge-base with HEAD, so a base tip that moved on doesn't spill its own
commits into the range. ci/lean.py <base> <head> overrides both for probing.

  replace-is-delete  an added `Deprecated:` line or an alias marker in a
                      non-test file means the old thing stayed behind: delete
                      it, don't park it beside the new one
  justified-deps      a require gained by go/go.mod or the example webapp's
                      go.mod needs a `deps: <module> — <why>` line in a commit
                      in the range
  generated-first     a hand-written type under go/v1/traust that
                      client.gen.go already declares is a hand-edit of
                      generated ground

A go.mod that first appears in the range is measured from the commit that
introduced it: a new module's initial require set ships with the module in
one review, and the gate stops it growing afterwards. `// indirect` requires
are the tree a chosen dep brings; only the dep a person chose carries a line.

Exit 0 when lean, 1 on any finding.
"""

from __future__ import annotations

import os
import re
import sys
from collections.abc import Iterator
from dataclasses import dataclass
from pathlib import Path

_CI_DIR = Path(__file__).resolve().parent
if str(_CI_DIR) not in sys.path:
    sys.path.insert(0, str(_CI_DIR))

from commits import GitError, git, log  # noqa: E402

ROOT = _CI_DIR.parent
DEFAULT_BASE_REF = "origin/development"

#: The gate's own vocabulary lives in this file; scanning it would flag the
#: gate, not the code.
SELF = Path(__file__).resolve()

GO_MODS = ("go/go.mod", "go/examples/webapp/go.mod")
GEN_CLIENT = "go/v1/traust/client.gen.go"
TRAUST_DIR = "go/v1/traust"

DEPRECATED_RE = re.compile(r"deprecated:", re.IGNORECASE)
ALIAS_RE = re.compile(r"\b(?:compat|shim|legacy alias)\b", re.IGNORECASE)
DEPS_LINE_RE = re.compile(r"^deps:[ \t]*(?P<module>\S+)[ \t]*[—-][ \t]*(?P<why>\S.*)$")
HUNK_RE = re.compile(r"^@@ -\d+(?:,\d+)? \+(\d+)(?:,\d+)? @@")
IDENT_RE = re.compile(r"[A-Za-z_][A-Za-z0-9_]*")
TOKEN_RE = re.compile(r"[A-Za-z_][A-Za-z0-9_]*|[{}()\[\];\n]|\S")


@dataclass(frozen=True)
class AddedLine:
    path: str
    line: int
    text: str


def _resolve_base(base: str, head: str) -> str:
    ref = base or os.environ.get("LEAN_BASE", "").strip() or DEFAULT_BASE_REF
    merge = git("merge-base", ref, head).strip()
    if not merge:
        raise GitError(f"no merge-base between {ref} and {head}")
    return merge


def _is_test(path: str) -> bool:
    name = Path(path).name
    return name.endswith(("_test.go", "_test.py")) or name.startswith("test_") or "/testdata/" in path


def _added_lines(base: str, head: str) -> Iterator[AddedLine]:
    """Added lines with their new-file line numbers, from a zero-context diff."""
    diff = git("diff", "--no-color", "--unified=0", f"{base}..{head}")
    path: str | None = None
    line = 0
    for raw in diff.splitlines():
        if raw.startswith("+++"):
            if raw.startswith("+++ b/"):
                path = raw[6:]
            continue
        hunk = HUNK_RE.match(raw)
        if hunk:
            line = int(hunk.group(1))
            continue
        if raw.startswith("+"):
            yield AddedLine(path or "?", line, raw[1:])
            line += 1


def _replace_is_delete(base: str, head: str) -> list[str]:
    findings = []
    for added in _added_lines(base, head):
        if _is_test(added.path) or (ROOT / added.path).resolve() == SELF:
            continue
        if DEPRECATED_RE.search(added.text) or ALIAS_RE.search(added.text):
            findings.append(
                f"replace-is-delete {added.path}:{added.line}: added {added.text.strip()!r}"
                " — the old thing stayed; delete it, don't park it beside the new one"
            )
    return findings


def _git_show(rev: str) -> str:
    try:
        return git("show", rev)
    except GitError:
        return ""


def _requires(text: str) -> set[str]:
    """The module paths a go.mod requires directly."""
    mods: set[str] = set()
    block = False
    for raw in text.splitlines():
        line = raw.strip()
        if block:
            if line == ")":
                block = False
            elif line and not line.startswith("//") and "// indirect" not in line:
                mods.add(line.split()[0])
        elif line.startswith("require ("):
            block = True
        elif line.startswith("require ") and "// indirect" not in line:
            rest = line[len("require ") :].strip()
            if rest and not rest.startswith("//"):
                mods.add(rest.split()[0])
    return mods


def _mod_baseline(path: str, base: str, head: str) -> set[str]:
    """The requires the range starts from: the go.mod at base, or — when the
    module is itself new in the range — at the commit that introduced it."""
    if text := _git_show(f"{base}:{path}"):
        return _requires(text)
    for sha in git("log", "--reverse", "--format=%H", f"{base}..{head}").split():
        if text := _git_show(f"{sha}:{path}"):
            return _requires(text)
    return set()


def _justified(base: str, head: str) -> dict[str, str]:
    """Module paths a `deps: <module> — <why>` line in the range names."""
    justified: dict[str, str] = {}
    for commit in log(f"{base}..{head}"):
        for line in f"{commit.subject}\n{commit.body}".splitlines():
            match = DEPS_LINE_RE.match(line.strip())
            if match:
                justified.setdefault(
                    match.group("module"), f"{commit.sha[:8]}: {line.strip()}"
                )
    return justified


def _justified_deps(base: str, head: str) -> list[str]:
    justified = _justified(base, head)
    findings = []
    for path in GO_MODS:
        baseline = _mod_baseline(path, base, head)
        for module in sorted(_requires(_git_show(f"{head}:{path}")) - baseline):
            if note := justified.get(module):
                print(f"{path}: new require {module} — justified by {note}")
            else:
                findings.append(
                    f"justified-deps {path}: new require {module} — no"
                    f" `deps: {module} — <why>` line in any commit in"
                    f" {base[:8]}..{head}"
                )
    return findings


def _blanked(source: str) -> str:
    """Go source with comments and string/rune contents spaced out, so the
    token scan sees structure only; offsets and line numbers are kept."""
    chars = list(source)
    i, n = 0, len(source)
    while i < n:
        char = source[i]
        nxt = source[i + 1] if i + 1 < n else ""
        if char == "/" and nxt == "/":
            while i < n and chars[i] != "\n":
                chars[i] = " "
                i += 1
        elif char == "/" and nxt == "*":
            chars[i] = chars[i + 1] = " "
            i += 2
            while i < n and not (source[i] == "*" and source[i + 1 : i + 2] == "/"):
                if chars[i] != "\n":
                    chars[i] = " "
                i += 1
            if i < n:
                chars[i] = chars[i + 1] = " "
                i += 2
        elif char in "\"'`":
            quote = char
            chars[i] = " "
            i += 1
            while i < n and chars[i] != quote:
                if chars[i] == "\\" and quote != "`" and i + 1 < n:
                    chars[i] = chars[i + 1] = " "
                    i += 2
                    continue
                if chars[i] != "\n":
                    chars[i] = " "
                i += 1
            if i < n:
                chars[i] = " "
                i += 1
        else:
            i += 1
    return "".join(chars)


def _declared_types(source: str) -> dict[str, int]:
    """The type names a Go file declares, with line numbers — a token scan of
    the blanked source tracking nesting, not a grep for `^type `."""
    names: dict[str, int] = {}
    line = 1
    paren = brace = bracket = 0
    in_group = False  # inside a `type (` block
    entry_pending = False  # the next ident at the group's top level is a name
    after_type = False  # a top-level `type` keyword, awaiting its name or `(`
    for match in TOKEN_RE.finditer(_blanked(source)):
        tok = match.group(0)
        if tok == "\n":
            line += 1
            if in_group and paren == 1 and brace == bracket == 0:
                entry_pending = True
            continue
        if IDENT_RE.fullmatch(tok):
            if tok == "type" and not in_group and paren == brace == bracket == 0:
                after_type = True
            elif after_type:
                after_type = False
                names.setdefault(tok, line)
            elif in_group and entry_pending and paren == 1 and brace == bracket == 0:
                if tok != "type":
                    names.setdefault(tok, line)
                    entry_pending = False
            continue
        if tok == "(":
            if after_type:
                in_group = True
                entry_pending = True
                after_type = False
            paren += 1
        elif tok == ")":
            if in_group and paren == 1:
                in_group = False
                entry_pending = False
            paren -= 1
        elif tok == "{":
            brace += 1
        elif tok == "}":
            brace -= 1
        elif tok == "[":
            bracket += 1
        elif tok == "]":
            bracket -= 1
        elif tok == ";" and in_group and paren == 1 and brace == bracket == 0:
            entry_pending = True
    return names


def _generated_first() -> list[str]:
    gen_names = _declared_types((ROOT / GEN_CLIENT).read_text(encoding="utf-8"))
    findings = []
    for path in sorted((ROOT / TRAUST_DIR).rglob("*.go")):
        rel = path.relative_to(ROOT).as_posix()
        if rel.endswith((".gen.go", "_test.go")):
            continue
        for name, line in _declared_types(path.read_text(encoding="utf-8")).items():
            if gen_line := gen_names.get(name):
                findings.append(
                    f"generated-first {rel}:{line}: type {name} — {GEN_CLIENT}:"
                    f"{gen_line} already declares it; the generated client owns"
                    " the name, a better shape is a generator change"
                )
    return findings


def main(argv: list[str] | None = None) -> int:
    args = argv if argv is not None else sys.argv[1:]
    if len(args) > 2:
        print("usage: lean.py [base] [head]", file=sys.stderr)
        return 2
    base = args[0] if args else ""
    head = args[1] if len(args) > 1 else "HEAD"
    try:
        base = _resolve_base(base, head)
        findings = _replace_is_delete(base, head)
        findings += _justified_deps(base, head)
    except GitError as exc:
        print(f"lean: {exc}", file=sys.stderr)
        return 1
    findings += _generated_first()
    if findings:
        print("lean check failed — the SDK grew:", file=sys.stderr)
        for finding in findings:
            print(f"  {finding}", file=sys.stderr)
        return 1
    print(f"ok: lean {base[:8]}..{head}")
    return 0


if __name__ == "__main__":
    sys.exit(main())
