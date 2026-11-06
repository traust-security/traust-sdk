CONTRACTS_REPO := https://github.com/traust-security/traust-contracts.git
CONTRACTS_REF := bf4eb27b9eff357587fe1369db3b43e4a9534eba
CONTRACTS_CACHE := $(or $(XDG_CACHE_HOME),$(HOME)/.cache)/traust-sdk/contracts
CONTRACTS_SOURCE := $(CONTRACTS_CACHE)/checkout-$(CONTRACTS_REF)
ENUMS_DIR := $(CONTRACTS_SOURCE)/enums/v1
SCHEMAS_DIR := $(CONTRACTS_SOURCE)/artifacts_schema/v1
# storage/ is frozen at the pre-db/v1 layout: the storage generator reads queries/ dirs the
# db/v1 restructure dropped. Regenerating it is the G6 layout adaptation (plan 28), and plan 05
# retires the package; until then v1/storage/ is hand-maintained and checked by drift.
CORE_REPO := https://github.com/traust-security/traust-core.git
CORE_REF := 908bfbd613b3317d16ed7bce2e8259fdb35fd786
CORE_SOURCE := $(or $(XDG_CACHE_HOME),$(HOME)/.cache)/traust-sdk/core/checkout-$(CORE_REF)
CORE_PYTHON := cd "$(CORE_SOURCE)" && uv run --quiet python
OPENAPI_DIR := ./internal/generate/openapi
OUT_DIR := ./v1/enums
ROOT := ..
GO_CI := python3 $(ROOT)/ci/lang/go/ci.py
export PATH := $(shell go env GOPATH)/bin:$(PATH)

.PHONY: contracts core generate api-generate e2e test test-integration check-drift fmt fmt-check vet lint security tools

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
	go run ./internal/generate -enums "$(ENUMS_DIR)" -out "$(OUT_DIR)" -schemas "$(SCHEMAS_DIR)" -types ./v1/types -validate ./v1/validate/schemas_gen.go -contracts-ref "$(CONTRACTS_REF)"

test:
	go test ./...

test-integration:
	go test -count=1 -v ./v1/storage/...

check-drift: generate api-generate
	@git diff --exit-code $(OUT_DIR) ./v1/types/ ./v1/validate/schemas_gen.go ./v1/traust/client.gen.go ./v1/traust/calls.gen.go ./v1/ops/pack.schema.json || (echo "ERROR: generated files are out of date; run 'make generate api-generate' and commit" && exit 1)

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

# Everything the SDK knows about core comes from core at CORE_REF: the client from its OpenAPI
# (built, used once, deleted) and the pack schema from the contracts core validates packs with.
api-generate: core
	@set -eu; \
	tmp="$$(mktemp -d)"; trap 'rm -rf "$$tmp"' EXIT; \
	$(MAKE) -s -C "$(CORE_SOURCE)" api-check; \
	python3 $(OPENAPI_DIR)/normalize.py "$(CORE_SOURCE)/.openapi/openapi.json" "$$tmp/openapi.json"; \
	go tool oapi-codegen --config $(OPENAPI_DIR)/oapi-codegen.yaml "$$tmp/openapi.json"; \
	go run ./internal/generate/calls ./v1/traust/client.gen.go "$$tmp/openapi.json" ./v1/traust/calls.gen.go; \
	($(CORE_PYTHON) -c "import json; from traust_core.v1.contracts import config_schema; print(json.dumps(config_schema('pack'), indent=2))") > ./v1/ops/pack.schema.json

# A local-issuer token for the e2e server: $(1) identity, $(2) machine (True/False), $(3) file.
mint = ($(CORE_PYTHON) -c "from pathlib import Path; from traust_core.v1.context import Config, Context; print(Context(Config.from_file(Path('$$home/config.yaml'))).local_issuer.mint('$(1)', machine=$(2)).token.get_secret_value())") > "$$home/$(3)"

# The generated client and the runner against traust-server built from core at CORE_REF.
# The suite never leaves localhost: the vocabulary test onboards one repository, and
# onboarding reads its visibility at the forge, so the e2eforge mock stands in for the
# forge. Core reaches a gitlab-kind host at https://{host}/api/v4 — scheme and port are
# fixed — so e2eforge is an HTTPS_PROXY CONNECT proxy answering with its own TLS (an
# ephemeral self-signed certificate, trusted through SSL_CERT_FILE), and the e2e config
# maps the mock's host to the gitlab kind in forge_hosts.
# One package at a time: both share the server, and ops onboards what v1/traust expects empty.
e2e: core
	@set -eu; \
	home="$$(mktemp -d)"; port=18080; \
	printf 'home: %s\nidentity:\n  trust_local: true\nforge_hosts:\n  e2e-forge.test: gitlab\n' "$$home" > "$$home/config.yaml"; \
	(cd "$(CORE_SOURCE)" && exec uv run --quiet python -c "import traust_core.v1.rest.server"); \
	go build -o "$$home/e2eforge" ./internal/e2eforge; \
	"$$home/e2eforge" -ca "$$home/forge-ca.pem" > "$$home/forge-port" & forge_pid=$$!; \
	for _ in $$(seq 100); do [ -s "$$home/forge-port" ] && break; sleep 0.1; done; \
	forge_port=$$(cat "$$home/forge-port"); \
	(cd "$(CORE_SOURCE)" && exec env SSL_CERT_FILE="$$home/forge-ca.pem" HTTPS_PROXY="http://127.0.0.1:$$forge_port" https_proxy="http://127.0.0.1:$$forge_port" uv run --quiet traust-server --config "$$home/config.yaml" --port $$port) > "$$home/server.log" 2>&1 & pid=$$!; \
	trap 'kill $$pid $$forge_pid 2>/dev/null || true; rm -rf "$$home"' EXIT; \
	for _ in $$(seq 60); do curl -sf -o /dev/null http://127.0.0.1:$$port/v1/healthz && break; sleep 1; done; \
	curl -sf -o /dev/null http://127.0.0.1:$$port/v1/healthz || { cat "$$home/server.log"; exit 1; }; \
	$(call mint,e2e@example.com,False,person); \
	$(call mint,e2e-runner,True,runner); \
	TRAUST_E2E_URL=http://127.0.0.1:$$port TRAUST_E2E_TOKEN_FILE="$$home/person" TRAUST_E2E_RUNNER_TOKEN_FILE="$$home/runner" \
	  go test -count=1 -p 1 -tags e2e -run E2E -v ./v1/traust/ ./v1/ops/

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
