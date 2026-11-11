#!/usr/bin/env bash
# AGENTS.md rules a machine can check. A failure names the rule; fix the code, not this file.
set -euo pipefail
cd "$(git rev-parse --show-toplevel)"

# Framework packages (AGENTS.md "How it works") and the frozen ones. A new package is a design
# change: add it here in the same PR, which routes the PR to a code owner.
framework="traust ops feed"
frozen="enums ledger skills storage types validate"

fail=0
rule() { echo "guardrail: $1" >&2; fail=1; }

if spec=$(git ls-files '*openapi*.json') && [ -n "$spec" ]; then
  rule "the OpenAPI spec is a build input, never committed: $spec"
fi

for dir in go/v1/*/; do
  pkg=$(basename "$dir")
  case " $framework $frozen " in *" $pkg "*) ;; *) rule "go/v1/$pkg is not a package in AGENTS.md" ;; esac
done

for pkg in $framework; do
  [ -d "go/v1/$pkg" ] || continue
  src=$(find "go/v1/$pkg" -name '*.go' ! -name '*_test.go' ! -name '*.gen.go')
  [ -n "$src" ] || continue
  # shellcheck disable=SC2086
  if hits=$(grep -nE 'os\.(Getenv|LookupEnv)\(' $src); then rule "library code never reads env vars:"$'\n'"$hits"; fi
  # shellcheck disable=SC2086
  if hits=$(grep -n '"database/sql"' $src); then rule "core's REST API is the only way in:"$'\n'"$hits"; fi
  if [ "$pkg" != traust ]; then
    # shellcheck disable=SC2086
    if hits=$(grep -nE 'http\.NewRequest|url\.JoinPath|"/v1/' $src); then
      rule "call the generated client, never build URLs:"$'\n'"$hits"
    fi
  fi
done

python3 ci/check_estate_data.py --quiet . || rule "deployment-specific data in a public repo"

if [ -n "${GUARDRAILS_BASE:-}" ]; then
  python3 ci/gates.py range "$GUARDRAILS_BASE" "${GUARDRAILS_HEAD:-HEAD}" || rule "commit subjects are conventional"
fi

exit "$fail"
