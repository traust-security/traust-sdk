#!/usr/bin/env bash
# AGENTS.md rules a machine can check. A failure names the rule; fix the code, not this file.
set -euo pipefail
cd "$(git rev-parse --show-toplevel)"

# The SDK's root package (go/README.md). A new package under go/v1 is a design change: add it
# here in the same PR, which routes the PR to a code owner.
framework="traust"
# The domain clients under go/v1/traust (go/README.md): one package per OpenAPI tag group.
domains="inventory runs findings remediation exposure ledger insight"

fail=0
rule() { echo "guardrail: $1" >&2; fail=1; }

if spec=$(git ls-files '*openapi*.json') && [ -n "$spec" ]; then
  rule "the OpenAPI spec is a build input, never committed: $spec"
fi

for dir in go/v1/*/; do
  pkg=$(basename "$dir")
  case " $framework " in *" $pkg "*) ;; *) rule "go/v1/$pkg is not a package in AGENTS.md" ;; esac
done

for dir in go/v1/traust/*/; do
  pkg=$(basename "$dir")
  [ "$pkg" = traustest ] && continue # a fake core for tests, not a domain client
  case " $domains " in *" $pkg "*) ;; *) rule "go/v1/traust/$pkg is not a domain package in AGENTS.md" ;; esac
done

# Every hand-written source under go/v1/traust and its domain packages.
src=$(find go/v1/traust -name '*.go' ! -name '*_test.go' ! -name '*.gen.go')
[ -n "$src" ] || rule "go/v1/traust has no hand-written sources"
# shellcheck disable=SC2086
if hits=$(grep -nE 'os\.(Getenv|LookupEnv)\(' $src); then rule "library code never reads env vars:"$'\n'"$hits"; fi
# shellcheck disable=SC2086
if hits=$(grep -n '"database/sql"' $src); then rule "core's REST API is the only way in:"$'\n'"$hits"; fi

# The domain clients call the generated client, never build URLs themselves; the root keeps
# oidc_tokens.go, the one hand-written request (the issuer's token endpoint, not core's API).
domain_src=$(find go/v1/traust -mindepth 2 -name '*.go' ! -path '*/traustest/*' ! -name '*_test.go' ! -name '*.gen.go')
# shellcheck disable=SC2086
if hits=$(grep -nE 'http\.NewRequest|url\.JoinPath|"/v1/' $domain_src); then
  rule "call the generated client, never build URLs:"$'\n'"$hits"
fi

python3 ci/check_estate_data.py --quiet . || rule "deployment-specific data in a public repo"

if [ -n "${GUARDRAILS_BASE:-}" ]; then
  python3 ci/gates.py range "$GUARDRAILS_BASE" "${GUARDRAILS_HEAD:-HEAD}" || rule "commit subjects are conventional"
fi

exit "$fail"
