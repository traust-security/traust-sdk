#!/usr/bin/env bash
# up.sh — the example box's startup, as named steps: write traust.yaml, start the server and
# its reports role, wait for health, mint the person/admin/runner tokens, start the planner,
# start the claim worker — and stop everything it started on exit. The Go side (make -C go e2e
# / webapp) calls this with the cached checkouts the Makefile fetched; every log lands in
# <home>/*.log, and the caller owns <home> itself (it creates and removes it).
#
# Flags:
#   --home DIR          the config home to write into (required)
#   --port N            the port traust-server serves on (required)
#   --core DIR          core's checkout; its environment runs the server, planner and worker
#   --engine DIR        the engine's checkout, installed into core's environment
#   --harness DIR       the traust harness's checkout (skill sessions, the shipped pack)
#   --worker NAME       the claim worker's name (e2e-box, webapp-box)
#   --asset NAME        run only this asset of the pack (repeatable; default: the whole pack)
#   --pack FILE         the pack to run (default: the harness's shipped security pack)
#   --claude-dir DIR    a dir whose claude stands in for the real one (the e2e stub)
#   --claude-login DIR  the Claude login dir the sessions use, symlinked into the home
#   --complete          the deployment-complete config home the engine's preflight demands
#   --gate FILE         hold the planner and claim worker until this file appears
# The environment carries the rest: GITHUB_TOKEN becomes onboarding's github.com forge token,
# and STORE=byos brings the box's own storage the way a deployment does (Postgres for records
# and an S3 store for artifacts, in throwaway containers; needs podman).
# <home>/ready appears once the server is healthy and the tokens are minted; <home>/boxed
# once the planner and the claim worker run.
set -euo pipefail

home="" port="" core="" engine="" harness="" worker=box pack="" claude_dir="" claude_login="" gate=""
assets="" complete=no
while [ $# -gt 0 ]; do
  case $1 in
    --home) home=$2; shift 2;;
    --port) port=$2; shift 2;;
    --core) core=$2; shift 2;;
    --engine) engine=$2; shift 2;;
    --harness) harness=$2; shift 2;;
    --worker) worker=$2; shift 2;;
    --pack) pack=$2; shift 2;;
    --asset) assets="$assets --asset $2"; shift 2;;
    --claude-dir) claude_dir=$2; shift 2;;
    --claude-login) claude_login=$2; shift 2;;
    --gate) gate=$2; shift 2;;
    --complete) complete=yes; shift;;
    *) echo "up.sh: unknown flag: $1" >&2; exit 2;;
  esac
done
[ -n "$home" ] && [ -n "$port" ] && [ -n "$core" ] && [ -n "$engine" ] && [ -n "$harness" ] \
  || { echo "up.sh: --home, --port, --core, --engine and --harness are required" >&2; exit 2; }
[ -n "$pack" ] || pack="$harness/packs/security/assets.yaml"
[ "$complete" = no ] || command -v cosign >/dev/null \
  || { echo "up.sh: the box signs its ledger with cosign: install it (https://docs.sigstore.dev)" >&2; exit 1; }
[ -n "$claude_dir" ] || command -v claude >/dev/null \
  || { echo "up.sh: the pack's skill sessions run with claude: install Claude Code (https://claude.com/claude-code)" >&2; exit 1; }

pids=""
containers=""
stop() { # the trap: stop what this script started (the caller removes the home)
  trap - INT TERM EXIT
  [ -z "$pids" ] || kill $pids 2>/dev/null || true
  [ -z "$containers" ] || podman stop -t 2 $containers >/dev/null 2>&1 || true
}
trap stop EXIT
trap 'exit 130' INT
trap 'exit 143' TERM

# Install the engine and the harness into core's environment: the box's planner, claim worker
# and skill sessions all run there.
(cd "$core" && uv pip install --quiet --no-deps "$engine" "$harness")

# traust.yaml: local-trusted identity, the example's recording grants, and a ledger; a
# GITHUB_TOKEN in the environment becomes onboarding's github.com forge token.
printf 'home: %s\nidentity:\n  trust_local: true\n  admins: [admin@example.com]\nrecording:\n  assessment: [scan]\n  triage: [triage]\n' "$home" > "$home/traust.yaml"
[ -z "${GITHUB_TOKEN:-}" ] || printf 'forge_tokens:\n  github.com: GITHUB_TOKEN\n' >> "$home/traust.yaml"

# Warm core's environment (uv sync) in the foreground, before any clock below starts.
(cd "$core" && exec uv run --quiet python -c "import traust_core.v1.rest.server")

# --complete: the config home the engine's preflight demands of a deployment — workspace
# locations, a cosign-signed ledger (unencrypted: the box is a throwaway) and the shipped
# model registry (no example skill uses a model).
if [ "$complete" = yes ]; then
  mkdir -p "$home/workspace" "$home/analysis-results" "$home/progress-tracker"
  (cd "$home" && COSIGN_PASSWORD= cosign generate-key-pair --output-key-prefix ledger >/dev/null 2>&1)
  printf 'locations:\n  workspace: %s/workspace\n  analysis_results: %s/analysis-results\n  progress_tracker: %s/progress-tracker\nledger:\n  key_path: %s/ledger.key\n  public_key_path: %s/ledger.pub\n' "$home" "$home" "$home" "$home" "$home" >> "$home/traust.yaml"
else
  printf 'ledger: {}\n' >> "$home/traust.yaml"
fi

# STORE=byos: bring its own storage the way a deployment does — Postgres for records and an
# S3 store (versitygw) for artifacts, in throwaway containers — then provision the database:
# the contracts' baseline DDL, the whole upgrade channel, then the data channel.
if [ "${STORE:-}" = byos ]; then
  podman run -d --rm --name traust-sdk-webapp-pg -e POSTGRES_USER=traust -e POSTGRES_PASSWORD=traust-test-only -e POSTGRES_DB=traust -p 127.0.0.1:15433:5432 docker.io/library/postgres:16 >/dev/null
  podman run -d --rm --name traust-sdk-webapp-s3 -e ROOT_ACCESS_KEY=traust -e ROOT_SECRET_KEY=traust-test-only -p 127.0.0.1:19000:7070 docker.io/versity/versitygw:latest posix /tmp >/dev/null
  containers="traust-sdk-webapp-pg traust-sdk-webapp-s3"
  for _ in $(seq 60); do podman exec traust-sdk-webapp-pg pg_isready -U traust -q 2>/dev/null && curl -s -o /dev/null http://127.0.0.1:19000/ && break; sleep 1; done
  curl -sf -o /dev/null --aws-sigv4 aws:amz:us-east-1:s3 --user traust:traust-test-only -X PUT http://127.0.0.1:19000/traust
  printf 'database:\n  url: postgresql+psycopg://traust@127.0.0.1:15433/traust\n  password_env: TRAUST_DB_PASSWORD\nartifacts:\n  s3:\n    bucket: traust\n    endpoint_url: http://127.0.0.1:19000\n    region: us-east-1\n    access_key_env: TRAUST_S3_ACCESS_KEY\n    secret_key_env: TRAUST_S3_SECRET_KEY\n' >> "$home/traust.yaml"
  export TRAUST_DB_PASSWORD=traust-test-only TRAUST_S3_ACCESS_KEY=traust TRAUST_S3_SECRET_KEY=traust-test-only
  (cd "$core" && exec uv run --quiet --extra postgres python - "$home/traust.yaml" <<'EOF'
from pathlib import Path
import sys
import sqlalchemy as sa
from traust_core.v1 import contracts, data_migrations
from traust_core.v1.context import Config
from traust_core.v1.repositories.sql import Dialect, apply_schema, on_driver_connection

engine = sa.create_engine(Config.from_file(Path(sys.argv[1])).database.engine_url())
dialect = Dialect.of(engine)
apply_schema(engine, contracts.database_ddl(dialect))
on_driver_connection(engine, lambda dbapi: contracts.run_db_upgrades(dbapi, dialect))
data_migrations.run(engine, cvss=None)
EOF
  )
fi

# The server and its reports role, on the port the caller chose.
(cd "$core" && exec env COSIGN_PASSWORD= uv run --quiet traust-server --config "$home/traust.yaml" --port "$port") > "$home/server.log" 2>&1 & pids="$pids $!"
(cd "$core" && exec uv run --quiet traust-server --config "$home/traust.yaml" --role reports) > "$home/reports.log" 2>&1 & pids="$pids $!"

# Wait for health, showing the server's log when it never comes.
for _ in $(seq 60); do curl -sf -o /dev/null "http://127.0.0.1:$port/v1/healthz" && break; sleep 1; done
curl -sf -o /dev/null "http://127.0.0.1:$port/v1/healthz" || { cat "$home/server.log"; exit 1; }

# The sign-in tokens: a person, an admin (admins restate decisions) and a runner (machine).
mint() { # $1 identity, $2 machine (True/False), $3 file
  (cd "$core" && exec uv run --quiet python -c "from pathlib import Path; from traust_core.v1.context import Config, Context; print(Context(Config.from_file(Path('$home/traust.yaml'))).local_issuer.mint('$1', machine=$2).token.get_secret_value())") > "$home/$3"
}
mint e2e@example.com False person
mint admin@example.com False admin
mint e2e-runner True runner
touch "$home/ready"

# --gate: hold the planner and the claim worker until the caller releases them (the e2e
# suites run against the server before the planner starts ticking).
if [ -n "$gate" ]; then
  while [ ! -e "$gate" ]; do sleep 1; done
fi

# The model registry the planner and the worker run with: the shipped one.
cp "$harness/config/model-registry.yaml" "$home/model-registry.yaml"

# --claude-login: the box's Claude login, kept across runs, where the sessions find it (log
# in once with CLAUDE_CONFIG_DIR=<dir> claude, or set ANTHROPIC_API_KEY).
if [ -n "$claude_login" ]; then
  mkdir -p "$claude_login"
  ln -s "$claude_login" "$home/claude-runner"
fi

# The planner: the pack's work, planned every minute (the --every flag's finest cadence), so
# a person's submitted order is admitted promptly; holds a gated one for approval, withholds
# an asset the pack does not carry.
(cd "$core" && exec env TRAUST_CONFIG_HOME="$home" HARNESS_ROOT="$harness" uv run --quiet traust-engine run --core "$home/traust.yaml" --pack "$pack" $assets --router traust.ops.router:SecurityRouter --feeds "$core/src/traust_core/v1/context/shipped/feeds.yaml" --models "$home/model-registry.yaml" --every 1m) > "$home/engine.log" 2>&1 & pids="$pids $!"

# The claim worker that runs it: the deterministic assets' materializers and the llm assets'
# skill sessions (--claude-dir's stub stands in for claude when the e2e suites run).
claim_path="$PATH"
[ -z "$claude_dir" ] || claim_path="$claude_dir:$PATH"
(cd "$core" && exec env PATH="$claim_path" TRAUST_CONFIG_HOME="$home" HARNESS_ROOT="$harness" uv run --quiet traust-engine claim --loop --core "$home/traust.yaml" --pack "$pack" $assets --worker "$worker" --label llm --feeds "$core/src/traust_core/v1/context/shipped/feeds.yaml" --models "$home/model-registry.yaml") > "$home/claim.log" 2>&1 & pids="$pids $!"

touch "$home/boxed"
wait
