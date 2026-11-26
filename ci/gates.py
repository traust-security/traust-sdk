#!/usr/bin/env python3
"""CI gates: conventional-commit lint.

Used by:
  - .githooks/commit-msg  (subject validation)
  - ci/guardrails.sh      (range validation)
"""

from __future__ import annotations

import sys
from pathlib import Path

_CI_DIR = Path(__file__).resolve().parent
if str(_CI_DIR) not in sys.path:
    sys.path.insert(0, str(_CI_DIR))

from commits import HINT, GitError, validate_range, validate_subject  # noqa: E402


def _fail_commit(message: str) -> int:
    print("conventional-commit check failed:", file=sys.stderr)
    print(f"  {message}", file=sys.stderr)
    print(f"  {HINT}", file=sys.stderr)
    return 1


def main(argv: list[str] | None = None) -> int:
    args = argv if argv is not None else sys.argv[1:]
    if not args:
        print("usage: gates.py {subject|range} ...", file=sys.stderr)
        return 2

    match args[0]:
        case "subject":
            if len(args) < 2:
                print("usage: gates.py subject <message-file>", file=sys.stderr)
                return 2
            text = Path(args[1]).read_text(encoding="utf-8")
            subject = text.splitlines()[0].strip() if text else ""
            if err := validate_subject(subject):
                return _fail_commit(err)
            return 0

        case "range":
            if len(args) < 3:
                print("usage: gates.py range <base> <head>", file=sys.stderr)
                return 2
            try:
                if err := validate_range(args[1], args[2]):
                    return _fail_commit(err)
            except GitError as exc:
                print(f"gates: {exc}", file=sys.stderr)
                return 1
            return 0

        case _:
            print(f"unknown command: {args[0]}", file=sys.stderr)
            print("usage: gates.py {subject|range} ...", file=sys.stderr)
            return 2


if __name__ == "__main__":
    sys.exit(main())
