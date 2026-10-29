CONTRACTS_REPO := https://github.com/traust-security/traust-contracts.git
CONTRACTS_REF := ac517047c6566b019e63b911f6b7522000826b48
PACK_SCHEMA_REF := 4b7532bf044571a3a53077b03cf2955066a83374
CONTRACTS_CACHE := $(or $(XDG_CACHE_HOME),$(HOME)/.cache)/traust-sdk/contracts
CONTRACTS_SOURCE := $(CONTRACTS_CACHE)/checkout-$(CONTRACTS_REF)
PACK_SCHEMA_SOURCE := $(CONTRACTS_CACHE)/checkout-$(PACK_SCHEMA_REF)
ENUMS_DIR := $(CONTRACTS_SOURCE)/enums/v1
SCHEMAS_DIR := $(CONTRACTS_SOURCE)/schemas/v1
STORAGE_DIR := $(CONTRACTS_SOURCE)/storage/v1
FIXTURES_DIR := $(CONTRACTS_SOURCE)/tests/fixtures/storage
STORAGE_SAMPLES := $(CONTRACTS_SOURCE)/tests/storage_samples.py
CORE_REPO := https://github.com/traust-security/traust-core.git
CORE_REF := 977e8bf22117c2d58ba3c7bdfd38518274bb4f5b
CORE_SOURCE := $(or $(XDG_CACHE_HOME),$(HOME)/.cache)/traust-sdk/core/checkout-$(CORE_REF)
CORE_PYTHON := uv run --no-dev --quiet --extra rest python
OPENAPI_DIR := ./internal/generate/openapi
OUT_DIR := ./v1/enums
ROOT := ..
GO_CI := python3 $(ROOT)/ci/lang/go/ci.py
export PATH := $(shell go env GOPATH)/bin:$(PATH)

.PHONY: contracts pack-schema core generate api-generate e2e e2e-runner test test-integration check-drift fmt fmt-check vet lint security tools

contracts:
	@set -eu; \
	if [ ! -d "$(CONTRACTS_SOURCE)/.git" ]; then \
	  rm -rf "$(CONTRACTS_SOURCE)"; \
	  mkdir -p "$(CONTRACTS_SOURCE)"; \
	  git -C "$(CONTRACTS_SOURCE)" init --quiet; \
	  git -C "$(CONTRACTS_SOURCE)" remote add origin "$(CONTRACTS_REPO)"; \
	fi; \
	if ! git -C "$(CONTRACTS_SOURCE)" cat-file -e "$(CONTRACTS_REF)^{commit}" 2>/dev/null; then \
	  git -C "$(CONTRACTS_SOURCE)" fetch --quiet --no-tags --depth=1 origin "$(CONTRACTS_REF)"; \
	fi; \
	git -C "$(CONTRACTS_SOURCE)" checkout --quiet --detach --force "$(CONTRACTS_REF)"; \
	git -C "$(CONTRACTS_SOURCE)" clean -ffdqx

generate: contracts
	go run ./internal/generate -enums "$(ENUMS_DIR)" -out "$(OUT_DIR)" -schemas "$(SCHEMAS_DIR)" -types ./v1/types -validate ./v1/validate/schemas_gen.go -storage "$(STORAGE_DIR)" -storage-out ./v1/storage -storage-fixtures "$(FIXTURES_DIR)" -storage-samples "$(STORAGE_SAMPLES)" -contracts-ref "$(CONTRACTS_REF)"

pack-schema:
	@set -eu; \
	if [ ! -d "$(PACK_SCHEMA_SOURCE)/.git" ]; then \
	  rm -rf "$(PACK_SCHEMA_SOURCE)"; \
	  mkdir -p "$(PACK_SCHEMA_SOURCE)"; \
	  git -C "$(PACK_SCHEMA_SOURCE)" init --quiet; \
	  git -C "$(PACK_SCHEMA_SOURCE)" remote add origin "$(CONTRACTS_REPO)"; \
	fi; \
	if ! git -C "$(PACK_SCHEMA_SOURCE)" cat-file -e "$(PACK_SCHEMA_REF)^{commit}" 2>/dev/null; then \
	  git -C "$(PACK_SCHEMA_SOURCE)" fetch --quiet --no-tags --depth=1 origin "$(PACK_SCHEMA_REF)"; \
	fi; \
	git -C "$(PACK_SCHEMA_SOURCE)" checkout --quiet --detach --force "$(PACK_SCHEMA_REF)"; \
	cp "$(PACK_SCHEMA_SOURCE)/config/v1/pack.schema.json" ./v1/ops/pack.schema.json

test:
	go test ./...

test-integration:
	go test -count=1 -v ./v1/storage/...

check-drift: generate api-generate pack-schema
	@git diff --exit-code $(OUT_DIR) ./v1/types/ ./v1/validate/schemas_gen.go ./v1/storage/operations.go ./v1/storage/projection_identifiers.go ./v1/storage/schema.go ./v1/storage/queries.go ./v1/storage/storagetest/testdata/ ./v1/api/client.gen.go ./v1/ops/pack.schema.json || (echo "ERROR: generated files are out of date; run 'make generate api-generate pack-schema' and commit" && exit 1)

core:
	@set -eu; \
	if [ ! -d "$(CORE_SOURCE)/.git" ]; then \
	  rm -rf "$(CORE_SOURCE)"; \
	  mkdir -p "$(CORE_SOURCE)"; \
	  git -C "$(CORE_SOURCE)" init --quiet; \
	  git -C "$(CORE_SOURCE)" remote add origin "$(CORE_REPO)"; \
	fi; \
	if ! git -C "$(CORE_SOURCE)" cat-file -e "$(CORE_REF)^{commit}" 2>/dev/null; then \
	  git -C "$(CORE_SOURCE)" fetch --quiet --no-tags --depth=1 origin "$(CORE_REF)"; \
	fi; \
	git -C "$(CORE_SOURCE)" checkout --quiet --detach --force "$(CORE_REF)"; \
	git -C "$(CORE_SOURCE)" clean -ffdq

# The spec is built from core at CORE_REF, used once and deleted; only client.gen.go is committed.
api-generate: core
	@set -eu; \
	tmp="$$(mktemp -d)"; trap 'rm -rf "$$tmp"' EXIT; \
	$(MAKE) -s -C "$(CORE_SOURCE)" api-check; \
	python3 $(OPENAPI_DIR)/normalize.py "$(CORE_SOURCE)/.openapi/openapi.json" "$$tmp/openapi.json"; \
	go tool oapi-codegen --config $(OPENAPI_DIR)/oapi-codegen.yaml "$$tmp/openapi.json"

# The generated client against traust-server built from core at CORE_REF (SQLite, local issuer).
e2e: core
	@set -eu; \
	home="$$(mktemp -d)"; port=18080; \
	printf 'home: %s\nidentity:\n  trust_local: true\n' "$$home" > "$$home/config.yaml"; \
	(cd "$(CORE_SOURCE)" && exec uv run --quiet traust-server --config "$$home/config.yaml" --port $$port) > "$$home/server.log" 2>&1 & pid=$$!; \
	trap 'kill $$pid 2>/dev/null || true; rm -rf "$$home"' EXIT; \
	for _ in $$(seq 60); do curl -sf -o /dev/null http://127.0.0.1:$$port/v1/healthz && break; sleep 1; done; \
	curl -sf -o /dev/null http://127.0.0.1:$$port/v1/healthz || { cat "$$home/server.log"; exit 1; }; \
	(cd "$(CORE_SOURCE)" && uv run --quiet python -c "from pathlib import Path; from traust_core.v1.context import Config, Context; print(Context(Config.from_file(Path('$$home/config.yaml'))).local_issuer.mint('e2e@example.com').token.get_secret_value())") > "$$home/token"; \
	TRAUST_E2E_URL=http://127.0.0.1:$$port TRAUST_E2E_TOKEN_FILE="$$home/token" go test -count=1 -tags e2e -run E2E -v ./v1/api/

e2e-runner: core
	@set -eu; \
	home="$$(mktemp -d)"; port=18081; \
	pack="$(CURDIR)/v1/ops/testdata/local-command.yaml"; \
	printf 'home: %s\nidentity:\n  trust_local: true\n' "$$home" > "$$home/config.yaml"; \
	(cd "$(CORE_SOURCE)" && exec $(CORE_PYTHON) -c 'from traust_core.v1.rest.server import main; main()' --config "$$home/config.yaml" --port $$port) > "$$home/server.log" 2>&1 & pid=$$!; \
	trap 'kill $$pid 2>/dev/null || true; rm -rf "$$home"' EXIT; \
	for _ in $$(seq 60); do curl -sf -o /dev/null http://127.0.0.1:$$port/v1/healthz && break; sleep 1; done; \
	curl -sf -o /dev/null http://127.0.0.1:$$port/v1/healthz || { cat "$$home/server.log"; exit 1; }; \
	(cd "$(CORE_SOURCE)" && $(CORE_PYTHON) "$(CURDIR)/v1/ops/testdata/seed_order.py" "$$home/config.yaml" "$$pack" "$$home/token"); \
	TRAUST_E2E_URL=http://127.0.0.1:$$port TRAUST_E2E_TOKEN_FILE="$$home/token" go test -count=1 -tags e2e -run TestE2ERunnerLocalCommand -v ./v1/ops/

fmt:
	gofmt -w .

fmt-check:
	@test -z "$$(gofmt -l .)" || (gofmt -l . && exit 1)

vet:
	go vet ./...

lint:
	golangci-lint run ./...

tools:
	$(GO_CI) tools

security:
	$(GO_CI) security
