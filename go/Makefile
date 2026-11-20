CONTRACTS_REPO := https://github.com/traust-security/traust-contracts.git
CONTRACTS_REF := 31d0d884672149e606934c2844ffcbf3bdd4b925
CONTRACTS_CACHE := $(or $(XDG_CACHE_HOME),$(HOME)/.cache)/traust-sdk/contracts
CONTRACTS_SOURCE := $(CONTRACTS_CACHE)/checkout-$(CONTRACTS_REF)
ENUMS_DIR := $(CONTRACTS_SOURCE)/enums/v1
SCHEMAS_DIR := $(CONTRACTS_SOURCE)/artifacts_schema/v1
# storage/ is frozen at the pre-db/v1 layout: the storage generator reads queries/ dirs the
# db/v1 restructure dropped. Regenerating it is the G6 layout adaptation (plan 28), and plan 05
# retires the package; until then v1/storage/ is hand-maintained and checked by drift.
CORE_REPO := https://github.com/traust-security/traust-core.git
CORE_REF := 264736f36a45a471f46dd6abf68b9dc646091ca0
CORE_SOURCE := $(or $(XDG_CACHE_HOME),$(HOME)/.cache)/traust-sdk/core/checkout-$(CORE_REF)
CORE_PYTHON := cd "$(CORE_SOURCE)" && uv run --quiet python
ENGINE_REPO := https://github.com/traust-security/traust-engine.git
ENGINE_REF := c82854b33a538e985bee7c5d932194afa1edacf2
ENGINE_SOURCE := $(or $(XDG_CACHE_HOME),$(HOME)/.cache)/traust-sdk/engine/checkout-$(ENGINE_REF)
OPENAPI_DIR := ./internal/generate/openapi
OUT_DIR := ./v1/enums
ROOT := ..
GO_CI := python3 $(ROOT)/ci/lang/go/ci.py
export PATH := $(shell go env GOPATH)/bin:$(PATH)

.PHONY: contracts core engine generate api-generate e2e webapp test webapp-generate test-integration check-drift fmt fmt-check vet lint security tools

# checkout pins $(1) to commit $(3) of $(2), fetching only that commit; $(4) are git clean's flags.
define checkout
	@set -eu; \
	if [ ! -d "$(1)/.git" ]; then rm -rf "$(1)"; mkdir -p "$(1)"; git -C "$(1)" init --quiet; git -C "$(1)" remote add origin "$(2)"; fi; \
	if ! git -C "$(1)" cat-file -e "$(3)^{commit}" 2>/dev/null; then git -C "$(1)" fetch --quiet --no-tags --depth=1 origin "$(3)"; fi; \
	git -C "$(1)" checkout --quiet --detach --force "$(3)"; \
	git -C "$(1)" clean $(4)
endef

contracts:
	$(call checkout,$(CONTRACTS_SOURCE),$(CONTRACTS_REPO),$(CONTRACTS_REF),-ffdqx)

generate: contracts
	go run ./internal/generate -enums "$(ENUMS_DIR)" -out "$(OUT_DIR)" -schemas "$(SCHEMAS_DIR)" -types ./v1/types -validate ./v1/validate/schemas_gen.go -contracts-ref "$(CONTRACTS_REF)"

test:
	go test ./...
	cd examples/webapp && go test ./...

# The example app is its own module, so the SDK's go.mod never carries its web dependencies.
webapp-generate:
	cd examples/webapp && go tool templ generate

test-integration:
	go test -count=1 -v ./v1/storage/...

check-drift: generate api-generate webapp-generate
	@git diff --exit-code ./examples/webapp/views_templ.go $(OUT_DIR) ./v1/types/ ./v1/validate/schemas_gen.go ./v1/traust/client.gen.go ./v1/traust/calls.gen.go ./v1/traust/views.gen.go ./v1/traust/insights.gen.go ./v1/ops/pack.schema.json || (echo "ERROR: generated files are out of date; run 'make generate api-generate' and commit" && exit 1)

core:
	$(call checkout,$(CORE_SOURCE),$(CORE_REPO),$(CORE_REF),-ffdq)

# The engine runs in core's environment without its own dependencies, so the core that plans is
# the core at CORE_REF that serves.
engine: core
	$(call checkout,$(ENGINE_SOURCE),$(ENGINE_REPO),$(ENGINE_REF),-ffdq)
	@cd "$(CORE_SOURCE)" && uv pip install --quiet --no-deps "$(ENGINE_SOURCE)"

# Everything the SDK knows about core comes from core at CORE_REF: the client from its OpenAPI
# (built, used once, deleted), the Go types of the Insight view contracts core validates its
# views with, and the pack schema from the contracts core validates packs with.
api-generate: core
	@set -eu; \
	tmp="$$(mktemp -d)"; trap 'rm -rf "$$tmp"' EXIT; \
	$(MAKE) -s -C "$(CORE_SOURCE)" api-check; \
	python3 $(OPENAPI_DIR)/normalize.py "$(CORE_SOURCE)/.openapi/openapi.json" "$$tmp/openapi.json"; \
	go tool oapi-codegen --config $(OPENAPI_DIR)/oapi-codegen.yaml "$$tmp/openapi.json"; \
	go run ./internal/generate/calls ./v1/traust/client.gen.go "$$tmp/openapi.json" ./v1/traust/calls.gen.go; \
	($(CORE_PYTHON) "$(CURDIR)/$(OPENAPI_DIR)/views.py" "$$tmp/views.json"); \
	go tool oapi-codegen --config $(OPENAPI_DIR)/views.yaml "$$tmp/views.json"; \
	go run ./internal/generate/insights ./v1/traust/client.gen.go "$$tmp/views.json" ./v1/traust/insights.gen.go; \
	($(CORE_PYTHON) -c "import json; from traust_core.v1.contracts import config_schema; print(json.dumps(config_schema('pack'), indent=2))") > ./v1/ops/pack.schema.json

# A local-issuer token for the e2e server: $(1) identity, $(2) machine (True/False), $(3) file.
mint = ($(CORE_PYTHON) -c "from pathlib import Path; from traust_core.v1.context import Config, Context; print(Context(Config.from_file(Path('$$home/traust.yaml'))).local_issuer.mint('$(1)', machine=$(2)).token.get_secret_value())") > "$$home/$(3)"

# The generated client and the runner against traust-server built from core at CORE_REF.
# The suite never leaves localhost: the vocabulary test onboards one repository, and
# onboarding reads its visibility at the forge, so the e2eforge mock stands in for the
# forge. Core reaches a gitlab-kind host at https://{host}/api/v4 — scheme and port are
# fixed — so e2eforge is an HTTPS_PROXY CONNECT proxy answering with its own TLS (an
# ephemeral self-signed certificate, trusted through SSL_CERT_FILE), and the e2e config
# maps the mock's host to the gitlab kind in forge_hosts.
# One package at a time: both share the server, and ops onboards what v1/traust expects empty.
# provision is what a deployment does to a fresh configured database, the sequence core runs
# itself for a fresh local store: the contracts' baseline DDL, the whole upgrade channel, then
# the data channel.
export provision := from pathlib import Path; import sys, sqlalchemy as sa; \
	from traust_core.v1 import contracts, data_migrations; from traust_core.v1.context import Config; \
	from traust_core.v1.repositories.sql import Dialect, apply_schema, on_driver_connection; \
	engine = sa.create_engine(Config.from_file(Path(sys.argv[1])).database.engine_url()); dialect = Dialect.of(engine); \
	apply_schema(engine, contracts.database_ddl(dialect)); \
	on_driver_connection(engine, lambda dbapi: contracts.run_db_upgrades(dbapi, dialect)); \
	data_migrations.run(engine)

# core_up starts traust-server at CORE_REF on $$port beside its report builder, mints a person
# token ($$home/person), an admin's ($$home/admin: admins restate decisions) and a runner token
# ($$home/runner), and stops everything on exit. With
# forge=stand-in, core reaches forges only through the e2eforge stand-in (no network); otherwise
# it reaches real ones. The config grants the example's scan and triage assets their recording and
# keeps a ledger (unsigned unless complete=yes). With
# complete=yes the config home is what the engine's preflight demands of a deployment: workspace
# locations, a cosign-signed ledger, and a model registry (core's own fixture at CORE_REF; no
# example skill uses a model). With STORE=byos core brings its own storage the way a deployment
# does: Postgres for records and an S3 store (versitygw) for artifacts, in throwaway containers.
define core_up
	home="$$(mktemp -d)"; proxy=""; pids=""; containers=""; trap 'exit 130' INT TERM; \
	printf 'home: %s\nidentity:\n  trust_local: true\n  admins: [admin@example.com]\nrecording:\n  assessment: [scan]\n  triage: [triage]\n' "$$home" > "$$home/traust.yaml"; \
	(cd "$(CORE_SOURCE)" && exec uv run --quiet python -c "import traust_core.v1.rest.server"); \
	if [ "$${complete:-}" = yes ]; then \
	  mkdir -p "$$home/workspace" "$$home/analysis-results" "$$home/progress-tracker"; \
	  (cd "$$home" && COSIGN_PASSWORD= cosign generate-key-pair --output-key-prefix ledger >/dev/null 2>&1); \
	  printf 'locations:\n  workspace: %s/workspace\n  analysis_results: %s/analysis-results\n  progress_tracker: %s/progress-tracker\nledger:\n  key_path: %s/ledger.key\n  public_key_path: %s/ledger.pub\n' "$$home" "$$home" "$$home" "$$home" "$$home" >> "$$home/traust.yaml"; \
	  cp "$(CORE_SOURCE)/tests/fixtures/operations/model-registry.yaml" "$$home/"; \
	else \
	  printf 'ledger: {}\n' >> "$$home/traust.yaml"; \
	fi; \
	if [ "$${STORE:-}" = byos ]; then \
	  podman run -d --rm --name traust-sdk-webapp-pg -e POSTGRES_USER=traust -e POSTGRES_PASSWORD=traust-test-only -e POSTGRES_DB=traust -p 127.0.0.1:15433:5432 docker.io/library/postgres:16 >/dev/null; \
	  podman run -d --rm --name traust-sdk-webapp-s3 -e ROOT_ACCESS_KEY=traust -e ROOT_SECRET_KEY=traust-test-only -p 127.0.0.1:19000:7070 docker.io/versity/versitygw:latest posix /tmp >/dev/null; \
	  containers="traust-sdk-webapp-pg traust-sdk-webapp-s3"; \
	  trap 'podman stop -t 2 $$containers >/dev/null 2>&1 || true; rm -rf "$$home"' EXIT; \
	  for _ in $$(seq 60); do podman exec traust-sdk-webapp-pg pg_isready -U traust -q 2>/dev/null && curl -s -o /dev/null http://127.0.0.1:19000/ && break; sleep 1; done; \
	  curl -sf -o /dev/null --aws-sigv4 aws:amz:us-east-1:s3 --user traust:traust-test-only -X PUT http://127.0.0.1:19000/traust; \
	  export TRAUST_DB_PASSWORD=traust-test-only TRAUST_S3_ACCESS_KEY=traust TRAUST_S3_SECRET_KEY=traust-test-only; \
	  printf 'database:\n  url: postgresql+psycopg://traust@127.0.0.1:15433/traust\n  password_env: TRAUST_DB_PASSWORD\nartifacts:\n  s3:\n    bucket: traust\n    endpoint_url: http://127.0.0.1:19000\n    region: us-east-1\n    access_key_env: TRAUST_S3_ACCESS_KEY\n    secret_key_env: TRAUST_S3_SECRET_KEY\n' >> "$$home/traust.yaml"; \
	  (cd "$(CORE_SOURCE)" && uv run --quiet --extra postgres python -c "$$provision" "$$home/traust.yaml"); \
	fi; \
	if [ "$${forge:-}" = stand-in ]; then \
	  printf 'forge_hosts:\n  e2e-forge.test: gitlab\n' >> "$$home/traust.yaml"; \
	  go build -o "$$home/e2eforge" ./internal/e2eforge; \
	  "$$home/e2eforge" -ca "$$home/forge-ca.pem" > "$$home/forge-port" & pids="$$pids $$!"; \
	  for _ in $$(seq 100); do [ -s "$$home/forge-port" ] && break; sleep 0.1; done; \
	  forge_port=$$(cat "$$home/forge-port"); \
	  proxy="SSL_CERT_FILE=$$home/forge-ca.pem HTTPS_PROXY=http://127.0.0.1:$$forge_port https_proxy=http://127.0.0.1:$$forge_port"; \
	fi; \
	(cd "$(CORE_SOURCE)" && exec env COSIGN_PASSWORD= $$proxy uv run --quiet traust-server --config "$$home/traust.yaml" --port $$port) > "$$home/server.log" 2>&1 & pids="$$pids $$!"; \
	(cd "$(CORE_SOURCE)" && exec uv run --quiet traust-server --config "$$home/traust.yaml" --role reports) > "$$home/reports.log" 2>&1 & pids="$$pids $$!"; \
	trap 'kill $$pids 2>/dev/null || true; [ -z "$$containers" ] || podman stop -t 2 $$containers >/dev/null 2>&1; rm -rf "$$home"' EXIT; \
	for _ in $$(seq 60); do curl -sf -o /dev/null http://127.0.0.1:$$port/v1/healthz && break; sleep 1; done; \
	curl -sf -o /dev/null http://127.0.0.1:$$port/v1/healthz || { cat "$$home/server.log"; exit 1; }; \
	$(call mint,e2e@example.com,False,person); \
	$(call mint,admin@example.com,False,admin); \
	$(call mint,e2e-runner,True,runner)
endef

e2e: core
	@set -eu; port=18080; forge=stand-in; \
	$(core_up); \
	export TRAUST_E2E_URL=http://127.0.0.1:$$port TRAUST_E2E_TOKEN_FILE="$$home/person" TRAUST_E2E_ADMIN_TOKEN_FILE="$$home/admin" TRAUST_E2E_RUNNER_TOKEN_FILE="$$home/runner"; \
	go test -count=1 -p 1 -tags e2e -run E2E -v ./v1/traust/ ./v1/ops/; \
	cd examples/webapp && go test -count=1 -tags e2e -run E2E -v .

# The example app against a local traust-server that reaches real forges: sign in with the
# person token it prints, then onboard any public repository. The engine plans the pack's work
# (a scan per repository, on its trigger) every minute.
webapp: engine
	@set -eu; port=18080; complete=yes; \
	command -v cosign >/dev/null || { echo "make webapp signs its ledger with cosign: install it (https://docs.sigstore.dev)" >&2; exit 1; }; \
	$(core_up); \
	(cd "$(CORE_SOURCE)" && exec uv run --quiet traust-engine run --core "$$home/traust.yaml" --pack "$(CURDIR)/examples/webapp/pack.yaml" --every 1m) > "$$home/engine.log" 2>&1 & pids="$$pids $$!"; \
	echo "sign in at http://localhost:3000 as a person: $$(cat "$$home/person")"; \
	echo "or as an admin (restates decisions): $$(cat "$$home/admin")"; \
	cd examples/webapp && TRAUST_URL=http://127.0.0.1:$$port TRAUST_TOKEN_FILE="$$home/runner" go run .

fmt:
	gofmt -w .

fmt-check:
	@test -z "$$(gofmt -l .)" || (gofmt -l . && exit 1)

vet:
	go vet ./...
	cd examples/webapp && go vet ./...

lint:
	golangci-lint run ./...

tools:
	$(GO_CI) tools

security:
	$(GO_CI) security
