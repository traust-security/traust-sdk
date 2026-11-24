#!/usr/bin/env bash
# box-init — everything a deployment does before the box's services start: the config home
# (traust.yaml with its ledger keys, workspace locations and database), the database's schema,
# and the sign-in tokens the app prints. Compose runs this one-shot, to completion, before
# traust-server, the engine and the app start; the rest of the box reads what it wrote from the
# shared /box volume. Re-running it is safe: the keys are regenerated, provisioning is the
# upgrade channel core itself follows, and the tokens are minted again.
set -euo pipefail

home=/box
seed=/opt/box-seed

mkdir -p "$home"/workspace "$home"/analysis-results "$home"/progress-tracker
rm -f "$home"/ledger.key "$home"/ledger.pub
# The ledger's cosign key pair, generated the way `make webapp` does: unencrypted, because the
# box is a throwaway (COSIGN_PASSWORD empty — cosign refuses a truly empty prompt otherwise).
COSIGN_PASSWORD= cosign generate-key-pair --output-key-prefix "$home/ledger" >/dev/null 2>&1
cp "$seed/model-registry.yaml" "$home/model-registry.yaml"

# The config `make webapp` writes for the same box: local-trusted identity, the example's
# recording grants, the deployment's workspace locations, a signed ledger and the database.
cat > "$home/traust.yaml" <<EOF
home: $home
identity:
  trust_local: true
  admins: [admin@example.com]
recording:
  assessment: [scan]
  triage: [triage]
locations:
  workspace: $home/workspace
  analysis_results: $home/analysis-results
  progress_tracker: $home/progress-tracker
ledger:
  key_path: $home/ledger.key
  public_key_path: $home/ledger.pub
database:
  url: postgresql+psycopg://traust@postgres:5432/traust
  password_env: TRAUST_DB_PASSWORD
EOF

# Provision: the contracts' baseline DDL, the whole upgrade channel, then the data channel —
# what a deployment does to a fresh configured database (TRAUST_DB_PASSWORD comes from compose).
python - <<'EOF'
from pathlib import Path

import sqlalchemy as sa

from traust_core.v1 import contracts, data_migrations
from traust_core.v1.context import Config
from traust_core.v1.repositories.sql import Dialect, apply_schema, on_driver_connection

engine = sa.create_engine(Config.from_file(Path("/box/traust.yaml")).database.engine_url())
dialect = Dialect.of(engine)
apply_schema(engine, contracts.database_ddl(dialect))
on_driver_connection(engine, lambda dbapi: contracts.run_db_upgrades(dbapi, dialect))
data_migrations.run(engine, cvss=None)
EOF

# The server's compose healthcheck: its image is distroless (no shell), and podman's compose
# provider word-splits healthcheck arguments, so the check is this file — a command with no
# spaces — against core's liveness endpoint.
cat > "$home/healthz.py" <<'EOF'
"""The server container's healthcheck: core's liveness endpoint."""
import urllib.request

urllib.request.urlopen("http://localhost:8080/v1/healthz")
EOF

# The sign-in tokens: a person, an admin (admins restate decisions) and the app's runner
# (machine) token — the same identities `make webapp` mints.
python - <<'EOF'
from pathlib import Path

from traust_core.v1.context import Config, Context

issuer = Context(Config.from_file(Path("/box/traust.yaml"))).local_issuer
for name, identity, machine in (
    ("person", "e2e@example.com", False),
    ("admin", "admin@example.com", False),
    ("runner", "e2e-runner", True),
):
    Path(f"/box/{name}").write_text(
        issuer.mint(identity, machine=machine).token.get_secret_value()
    )
EOF

# The whole box runs as one user (1001, the box image's), so core's private logging
# directory (0700) is the engine's too; nothing here needs loosening.

echo "sign in at http://localhost:3000 as a person: $(cat "$home/person")"
echo "or as an admin (restates decisions): $(cat "$home/admin")"
