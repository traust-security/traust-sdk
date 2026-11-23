PYTHON ?= python3
GO_DIR := go

.PHONY: help setup hooks lint lint-fix test security guardrails lean e2e

help:
	@echo "Targets ($(notdir $(CURDIR))):"
	@echo "  make setup          — enable .githooks (run once per clone)"
	@echo "  make hooks          — git config core.hooksPath .githooks"
	@echo "  make lint           — go vet + gofmt check (go/)"
	@echo "  make lint-fix       — gofmt -w (go/)"
	@echo "  make test           — go test ./... (go/)"
	@echo "  make security       — govulncheck + gosec (go/)"
	@echo "  make guardrails     — AGENTS.md rules a machine can check"
	@echo "  make lean           — bloat gates vs LEAN_BASE: replace-is-delete, justified deps, generated-first"
	@echo "  make e2e            — generated client vs traust-server at CORE_REF"

setup: hooks
	@echo "ready — local hooks enabled (.githooks). Bypass: git commit --no-verify"

hooks:
	git config core.hooksPath .githooks
	@chmod +x .githooks/* 2>/dev/null || true

lint:
	$(MAKE) -C $(GO_DIR) vet fmt-check

lint-fix:
	$(MAKE) -C $(GO_DIR) fmt

test:
	$(MAKE) -C $(GO_DIR) test

security:
	$(MAKE) -C $(GO_DIR) security

guardrails:
	ci/guardrails.sh

lean:
	$(PYTHON) ci/lean.py

e2e:
	$(MAKE) -C $(GO_DIR) e2e
