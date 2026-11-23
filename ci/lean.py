#!/usr/bin/env python3
"""Lean gate: the SDK gets smaller or stays the same size.

Compares HEAD against a base and fails on four kinds of growth. The base is
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
  used-export         an exported func, type, var, const or method on an
                      exported type, added to a hand-written file under
                      go/v1/, that nothing in go/examples/ or a _test.go
                      references: an export exists for a caller

A go.mod that first appears in the range is measured from the commit that
introduced it: a new module's initial require set ships with the module in
one review, and the gate stops it growing afterwards. `// indirect` requires
are the tree a chosen dep brings; only the dep a person chose carries a line.
Declarations are read by a token scan of the source with comments and strings
blanked, not a grep. A new export is used when its name — `.Name(` for a
method — appears in go/examples/ or a _test.go; a type is used the same way
when one of its methods is called, since a type returned and used only
through inference never appears by name. `*.gen.go` files declare, never
count, and stay out of scope.

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
GO_V1 = "go/v1"
EXAMPLES = "go/examples"

DEPRECATED_RE = re.compile(r"deprecated:", re.IGNORECASE)
ALIAS_RE = re.compile(r"\b(?:compat|shim|legacy alias)\b", re.IGNORECASE)
DEPS_LINE_RE = re.compile(r"^deps:[ \t]*(?P<module>\S+)[ \t]*[—-][ \t]*(?P<why>\S.*)$")
HUNK_RE = re.compile(r"^@@ -\d+(?:,\d+)? \+(\d+)(?:,\d+)? @@")
IDENT_RE = re.compile(r"[A-Za-z_][A-Za-z0-9_]*")
TOKEN_RE = re.compile(r"[A-Za-z_][A-Za-z0-9_]*|[{}()\[\];\n]|\S")

DECL_KEYWORDS = frozenset(("type", "var", "const", "func"))
#: The idents a var or const entry may sit beside: `PhaseMarshal Phase =
#: "marshal"` names one const, not a second one called Phase.
BUILTIN_TYPES = frozenset((
    "any", "bool", "byte", "comparable", "complex64", "complex128", "error",
    "float32", "float64", "int", "int8", "int16", "int32", "int64", "rune",
    "string", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr",
))


@dataclass(frozen=True)
class AddedLine:
    path: str
    line: int
    text: str


@dataclass(frozen=True)
class Decl:
    """A top-level declaration in a Go file."""

    kind: str  # func, method, type, var or const
    name: str
    line: int
    recv: str = ""  # the receiver's type name, for methods


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


def _recv_type(tokens: list[str]) -> str:
    """The receiver's type name from the tokens inside its parens."""
    stop = tokens.index("[") if "[" in tokens else len(tokens)
    for tok in reversed(tokens[:stop]):
        if IDENT_RE.fullmatch(tok):
            return tok
    return ""


def _declarations(
    source: str, known_types: frozenset[str] = frozenset()
) -> list[Decl]:
    """Top-level declarations — funcs, methods, types, vars and consts — with
    line numbers: a token scan of the blanked source tracking nesting, not a
    grep for `^func `. known_types carries the type names the package
    declares elsewhere, so a var or const entry like `PhaseMarshal Phase =
    "marshal"` yields the name only, not the type beside it."""
    decls: list[Decl] = []
    line = 1
    paren = brace = bracket = 0
    keyword = ""  # a top-level decl keyword, awaiting its name or `(`
    kw_line = 0
    in_decl = False  # a named decl, done at its line's or body's end
    group = ""  # the decl keyword whose `(...)` block is open
    entry: list[str] = []  # tokens of the current group entry
    entry_line = 0
    entry_closed = False  # past the entry's `=`, until the next line or `;`
    recv: list[str] = []  # tokens inside a method receiver's parens
    in_recv = False
    after_recv = ""  # the receiver's type, awaiting the method's name

    def flush() -> None:
        nonlocal entry, entry_closed
        if group == "type" and entry and IDENT_RE.fullmatch(entry[0]):
            decls.append(Decl("type", entry[0], entry_line))
        elif group in ("var", "const") and entry:
            stop = entry.index("=") if "=" in entry else len(entry)
            for tok in entry[:stop]:
                if (
                    IDENT_RE.fullmatch(tok)
                    and tok not in BUILTIN_TYPES
                    and tok not in known_types
                ):
                    decls.append(Decl(group, tok, entry_line))
        entry = []
        entry_closed = False

    for match in TOKEN_RE.finditer(_blanked(source)):
        tok = match.group(0)
        if tok == "\n":
            line += 1
            if group and paren == 1 and brace == bracket == 0:
                flush()
            elif in_decl and paren == brace == bracket == 0:
                in_decl = False
            continue
        top = paren == brace == bracket == 0
        if IDENT_RE.fullmatch(tok):
            if in_recv:
                recv.append(tok)
            elif after_recv and top:
                decls.append(Decl("method", tok, kw_line, after_recv))
                after_recv = ""
                in_decl = True
            elif group and paren == 1 and brace == bracket == 0 and not entry_closed:
                if not entry:
                    entry_line = line
                entry.append(tok)
            elif keyword and top:
                decls.append(Decl(keyword, tok, kw_line))
                keyword = ""
                in_decl = True
            elif top and not in_decl and tok in DECL_KEYWORDS:
                keyword = tok
                kw_line = line
            continue
        if tok == "(":
            if keyword and top:
                if keyword == "func":
                    in_recv = True
                    recv = []
                else:
                    group = keyword
                    entry = []
                    entry_closed = False
                keyword = ""
            paren += 1
        elif tok == ")":
            if in_recv and paren == 1:
                in_recv = False
                after_recv = _recv_type(recv)
            elif group and paren == 1:
                flush()
                group = ""
            paren -= 1
        elif tok == "=" and group and paren == 1 and brace == bracket == 0:
            flush()
            entry_closed = True
        elif tok == ";":
            if group and paren == 1 and brace == bracket == 0:
                flush()
            elif top:
                in_decl = False
        elif tok == "{":
            brace += 1
        elif tok == "}":
            brace -= 1
            if in_decl and brace == bracket == paren == 0:
                in_decl = False
        elif tok == "[":
            bracket += 1
            if in_recv:
                recv.append(tok)
        elif tok == "]":
            bracket -= 1
            if in_recv:
                recv.append(tok)
        elif in_recv and tok not in "()":
            recv.append(tok)
    return decls


def _declared_types(source: str) -> dict[str, int]:
    """The type names a Go file declares, with line numbers."""
    names: dict[str, int] = {}
    for decl in _declarations(source):
        if decl.kind == "type":
            names.setdefault(decl.name, decl.line)
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


def _go_files(rev: str, prefix: str) -> list[str]:
    """The .go files under a directory at a revision."""
    return [
        path
        for path in git("ls-tree", "-r", "--name-only", rev, "--", prefix).split()
        if path.endswith(".go")
    ]


def _handwritten(rev: str) -> dict[str, str]:
    """Hand-written (not `*.gen.go`, not `_test.go`) sources under go/v1/, by
    path, at a revision."""
    return {
        path: _git_show(f"{rev}:{path}")
        for path in _go_files(rev, GO_V1)
        if not path.endswith((".gen.go", "_test.go"))
    }


def _usage_corpus(head: str) -> str:
    """go/examples/ and every _test.go under go/, comments and strings
    blanked: where a new export has to show up to earn its name."""
    paths = [
        path
        for path in _go_files(head, "go")
        if path.startswith(EXAMPLES) or path.endswith("_test.go")
    ]
    return "\n".join(_blanked(_git_show(f"{head}:{path}")) for path in paths)


def _used_exports(base: str, head: str) -> list[str]:
    """Exports added to hand-written go/v1/ files in the range that nothing
    in go/examples/ or a _test.go references."""
    known_types = frozenset(
        decl.name
        for rev in (base, head)
        for path in _go_files(rev, GO_V1)
        for decl in _declarations(_git_show(f"{rev}:{path}"))
        if decl.kind == "type"
    )
    base_files = _handwritten(base)
    head_files = _handwritten(head)
    before = {
        (decl.kind, decl.recv, decl.name)
        for src in base_files.values()
        for decl in _declarations(src, known_types)
    }
    #: A type reached only through inference — `poller, err := client.F()` —
    #: is used when its methods are; methods live in the same package dir.
    methods: dict[tuple[str, str], list[str]] = {}
    for path, src in head_files.items():
        for decl in _declarations(src, known_types):
            if decl.kind == "method":
                methods.setdefault((path.rsplit("/", 1)[0], decl.recv), []).append(
                    decl.name
                )
    corpus = _usage_corpus(head)
    findings = []
    for path, src in sorted(head_files.items()):
        for decl in _declarations(src, known_types):
            if not decl.name[:1].isupper():
                continue
            if decl.kind == "method" and not decl.recv[:1].isupper():
                continue  # a method on an unexported type is package-private
            if (decl.kind, decl.recv, decl.name) in before:
                continue
            used = re.search(
                rf"\.{decl.name}\s*\(" if decl.kind == "method" else rf"\b{decl.name}\b",
                corpus,
            )
            if not used and decl.kind == "type":
                used = any(
                    re.search(rf"\.{name}\s*\(", corpus)
                    for name in methods.get((path.rsplit("/", 1)[0], decl.name), ())
                )
            if used:
                continue
            findings.append(
                f"used-export {path}:{decl.line}: {decl.name} — exported and"
                f" added in {base[:8]}..{head}, but nothing in go/examples/ or"
                " any _test.go references it: an export someone must call, or"
                " keep it unexported"
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
        findings += _used_exports(base, head)
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
