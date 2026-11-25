# go/Makefile — the SDK's targets (`make help` lists them); the pins live in ../pins.yaml
# and the box's startup in ../box/up.sh (both shared by every language SDK).
# ../pins.yaml is the one place pins live (one why per pin); plain awk reads it, failing loudly
# unless every ref is a full 40-char sha. A CORE_REF=... on the command line still overrides.
pin = $(shell awk -v key='$(1):' -v field='$(2)' '$$1 == key {hit = 1; next} /^[^ \t]/ {hit = 0} hit && $$1 == field ":" {if (field == "ref" && ($$2 !~ /^[0-9a-f]+$$/ || length($$2) != 40)) {printf "pins.yaml: %s %s is not a full 40-char sha: %s\n", key, field, $$2 > "/dev/stderr"; bad = 1; exit 1} print $$2; ok = 1; exit} END {if (!ok && !bad) {printf "pins.yaml: no %s under %s\n", field, key > "/dev/stderr"; exit 1}}' ../pins.yaml)
CORE_REF := $(call pin,core,ref)
CORE_SOURCE := $(or $(XDG_CACHE_HOME),$(HOME)/.cache)/traust-sdk/core/checkout-$(CORE_REF)
ENGINE_REF := $(call pin,engine,ref)
ENGINE_SOURCE := $(or $(XDG_CACHE_HOME),$(HOME)/.cache)/traust-sdk/engine/checkout-$(ENGINE_REF)
TRAUST_REF := $(call pin,traust,ref)
TRAUST_SOURCE := $(or $(XDG_CACHE_HOME),$(HOME)/.cache)/traust-sdk/harness/checkout-$(TRAUST_REF)
$(if $(filter 120,$(shell printf '%s' '$(CORE_REF)$(ENGINE_REF)$(TRAUST_REF)' | wc -c)),,$(error pins.yaml: every ref must be a full 40-char sha))
# The box's Claude login, kept across runs: log in once with CLAUDE_CONFIG_DIR=$(CLAUDE_RUNNER) claude, or set ANTHROPIC_API_KEY.
CLAUDE_RUNNER ?= $(or $(XDG_CONFIG_HOME),$(HOME)/.config)/traust-sdk/claude-runner
OPENAPI_DIR := ./internal/generate/openapi
export PATH := $(shell go env GOPATH)/bin:$(PATH)

.PHONY: help checkouts api-generate check-drift generate test fmt fmt-check vet lint security tools e2e webapp webapp-compose

# help — this list, read from the one-line comments below.
help:
	@grep -E '^# [a-z0-9-]+ —' $(MAKEFILE_LIST) | sed 's/^# //'
# checkout pins $(1) to commit $(3) of $(2), fetching only that commit; $(4) are git clean's flags.
define checkout
	@set -eu; \
	if [ ! -d "$(1)/.git" ]; then rm -rf "$(1)"; mkdir -p "$(1)"; git -C "$(1)" init --quiet; git -C "$(1)" remote add origin "$(2)"; fi; \
	if ! git -C "$(1)" cat-file -e "$(3)^{commit}" 2>/dev/null; then git -C "$(1)" fetch --quiet --no-tags --depth=1 origin "$(3)"; fi; git -C "$(1)" checkout --quiet --detach --force "$(3)"; git -C "$(1)" clean $(4)
endef
# checkouts — the box's three sources at their pins, into the cache every target shares (up.sh installs the engine and the harness into core's environment).
checkouts:
	$(call checkout,$(CORE_SOURCE),$(call pin,core,repo),$(CORE_REF),-ffdq)
	$(call checkout,$(ENGINE_SOURCE),$(call pin,engine,repo),$(ENGINE_REF),-ffdq)
	$(call checkout,$(TRAUST_SOURCE),$(call pin,traust,repo),$(TRAUST_REF),-ffdq)

# api-generate — everything the SDK knows about core, from core at its pin: the client from its OpenAPI (built, used once, deleted) (the SDK's only source of types).
api-generate: checkouts
	@set -eu; tmp="$$(mktemp -d)"; trap 'rm -rf "$$tmp"' EXIT; $(MAKE) -s -C "$(CORE_SOURCE)" api-check; \
	python3 $(OPENAPI_DIR)/normalize.py "$(CORE_SOURCE)/.openapi/openapi.json" "$$tmp/openapi.json"; go tool oapi-codegen --config $(OPENAPI_DIR)/oapi-codegen.yaml "$$tmp/openapi.json"; go run ./internal/generate/calls ./v1/traust/client.gen.go "$$tmp/openapi.json" ./v1/traust;

# check-drift — regenerate everything, fail when committed output is stale.
check-drift: api-generate generate
	@git diff --exit-code -- './examples/webapp/*_templ.go' ./v1/traust/client.gen.go ./v1/traust/calls.gen.go './v1/traust/*/*.gen.go' || (echo "ERROR: generated files are out of date; run 'make api-generate generate' and commit" && exit 1)

# generate — the example app's templ templates (its own module: no web deps in the SDK's go.mod); test — the network-free unit suites, SDK and example app.
generate: ; cd examples/webapp && go tool templ generate
test: ; go test ./... && cd examples/webapp && go test ./...
# fmt — gofmt everything; fmt-check — fail when not gofmt-clean; vet — go vet; lint — golangci-lint.
fmt: ; gofmt -w .
fmt-check: ; @test -z "$$(gofmt -l .)" || (gofmt -l . && exit 1)
vet: ; go vet ./... && cd examples/webapp && go vet ./...
lint: ; golangci-lint run ./...
# security — govulncheck and gosec; tools — install the pinned tools when missing.
security: ; python3 ../ci/lang/go/ci.py security
tools: ; python3 ../ci/lang/go/ci.py tools

# box_start — the shared start of the e2e and webapp targets: a free port, a throwaway home,
# traps that stop the box, then up.sh in the background with $(1) as its extra flags.
define box_start
	@set -eu; port=$$(python3 -c 'import socket; s = socket.socket(); s.bind(("127.0.0.1", 0)); print(s.getsockname()[1])'); home=$$(mktemp -d); \
	trap 'kill "$$box" 2>/dev/null || true; wait "$$box" 2>/dev/null || true; rm -rf "$$home"' EXIT; trap 'exit 130' INT; trap 'exit 143' TERM; \
	../box/up.sh --home "$$home" --port "$$port" --core "$(CORE_SOURCE)" --engine "$(ENGINE_SOURCE)" --harness "$(TRAUST_SOURCE)" $(1) >"$$home/box.log" 2>&1 & box=$$!; \
	for _ in $$(seq 900); do [ -f "$$home/ready" ] && break; kill -0 "$$box" 2>/dev/null || break; sleep 1; done; \
	[ -f "$$home/ready" ] || { echo "--- box.log"; cat "$$home/box.log"; exit 1; }
endef

# e2e — the SDK's and the app's suites end to end through the box: onboarding reads a real
# public GitHub repository (a GITHUB_TOKEN, if set, lifts its rate limit) and the e2e stub
# stands in for claude. The SDK suites run before the planner ticks (the gate).
e2e: checkouts
	$(call box_start,--worker e2e-box --asset full_audit --asset traust_triage --claude-dir ../box/e2eclaude --complete --gate "$$home/go"); \
	export TRAUST_E2E_URL=http://127.0.0.1:$$port TRAUST_E2E_TOKEN_FILE="$$home/person" TRAUST_E2E_ADMIN_TOKEN_FILE="$$home/admin" TRAUST_E2E_RUNNER_TOKEN_FILE="$$home/runner"; go test -count=1 -p 1 -tags e2e -run E2E -v ./v1/traust/...; touch "$$home/go"; \
	for _ in $$(seq 120); do [ -f "$$home/boxed" ] && break; kill -0 "$$box" 2>/dev/null || break; sleep 1; done; [ -f "$$home/boxed" ] || { echo "--- box.log"; cat "$$home/box.log"; exit 1; }; \
	(cd examples/webapp && go test -count=1 -tags e2e -run E2E -v .) || { echo "--- engine.log"; tail -40 "$$home/engine.log"; echo "--- claim.log"; tail -40 "$$home/claim.log"; echo "--- e2eclaude.log"; cat "$$home/e2eclaude.log" 2>/dev/null; exit 1; }

# webapp — the example app as the box: the server, the planner and the claim worker beside the
# app's pages. Needs uv, cosign, gitleaks and a Claude login (CLAUDE_RUNNER or
# ANTHROPIC_API_KEY); STORE=byos brings its own storage (podman). Sign in with the tokens.
webapp: checkouts
	$(call box_start,--worker webapp-box --claude-login "$(CLAUDE_RUNNER)" --complete); \
	echo "sign in at http://localhost:3000 as a person: $$(cat "$$home/person")"; echo "or as an admin (restates decisions): $$(cat "$$home/admin")"; \
	cd examples/webapp && go run . -core-url http://127.0.0.1:$$port

# webapp-compose — the same box as containers, every image built here from the pinned
# checkouts (podman's compose provider builds with the classic builder, which takes no extra
# build contexts). Needs only podman with a compose provider; `podman compose -f
# examples/webapp/compose.yaml down -v` tears it down.
webapp-compose: checkouts
	@command -v podman >/dev/null && podman compose version >/dev/null 2>&1 || { echo "webapp-compose needs podman (with a compose provider)" >&2; exit 1; }
	@$(MAKE) -C "$(ENGINE_SOURCE)" gitleaks && mkdir -p "$(CLAUDE_RUNNER)"
	@podman build -t localhost/traust-webapp-core:local "$(CORE_SOURCE)" && podman build -t localhost/traust-webapp:local -f examples/webapp/Containerfile ..
	@podman build --build-context core="$(CORE_SOURCE)" --build-context engine="$(ENGINE_SOURCE)" --build-context harness="$(TRAUST_SOURCE)" -t localhost/traust-webapp-box:local -f ../box/Containerfile ..
	@cd examples/webapp && CLAUDE_RUNNER="$(CLAUDE_RUNNER)" podman compose up -d --wait && podman compose logs --no-log-prefix box-init | sed -n '/^sign in at/p;/^or as an admin/p'
