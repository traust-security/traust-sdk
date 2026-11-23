# Contributing

## Setup

```bash
git clone <repo-url> && cd traust-sdk
make setup
```

Enables git hooks. One time per clone. Go SDK requires **Go 1.26+** (`toolchain` in `go/go.mod` pins the patched release for `govulncheck`).

## Regenerating

```bash
make -C go api-generate  # client, calls and pack schema from core at CORE_REF
make -C go generate      # the legacy contracts-generated packages at CONTRACTS_REF
make -C go check-drift   # regenerates everything, fails when committed output is stale
```

`api-generate` builds core's OpenAPI at `CORE_REF` (`go/Makefile`), generates
`v1/traust/client.gen.go` and the `(T, error)` calls and `…All` iterators in
`v1/traust/calls.gen.go`, writes `v1/ops/pack.schema.json`, and deletes the spec.

`generate` fetches the canonical `traust-contracts` repository at the immutable `CONTRACTS_REF`
pinned in `go/Makefile` — never latest `main` — and validates that the checkout matches it
exactly; no sibling checkout is required. The verified checkout is reused from `CONTRACTS_CACHE`
(by default the user cache) at `CONTRACTS_SOURCE`. `CONTRACTS_REF`, `CONTRACTS_REPO`,
`CONTRACTS_CACHE` and `CONTRACTS_SOURCE` are all caller-overridable for fork, CI-cache and
pinned-ref workflows.

## Commit messages

Conventional commits required. Format: `type(scope): subject`

Types: `feat`, `fix`, `perf`, `refactor`, `docs`, `test`, `chore`, `ci`, `build`, `style`, `revert`

Use the language id as scope when the change is SDK-specific:

```
feat(go): add harness provider interface
fix(go): correct enum code generation
ci: update shared gates
```

## Hooks

| Hook | Runs | Speed |
|------|------|-------|
| `commit-msg` | subject format check | instant |
| `pre-commit` | gofmt on staged `.go` | <1s |
| `pre-push` | `go vet` + `go test ./...` | ~10s |

Bypass: `--no-verify` on commit or push. CI still enforces.

## CI pipeline

GitHub Actions (`.github/workflows/ci.yml`) on every PR, in the merge queue and after each merge to `main` / `development`. Each job is a Makefile target, so CI and local behave the same:

| Job | Runs |
|-----|------|
| `lint` | `make lint` |
| `test` | `go test -race ./...` |
| `integration` | `make test-integration` on PostgreSQL 16 |
| `security` | `make security` |
| `drift` | `make -C go check-drift`: generated code matches the pinned contracts and core commits |
| `e2e` | `make e2e`: the generated client against `traust-server` built from core at `CORE_REF` |
| `guardrails` | `make guardrails`: the AGENTS.md rules a machine can check, plus commit subjects |

`core-head.yml` runs nightly against core's `development` tip; red means bump `CORE_REF`. Paths in `CODEOWNERS` wait for a code owner.

## Running tests

```bash
make test              # network-free unit tests
make test-integration  # SQLite plus optional local PostgreSQL storage coverage
make security          # govulncheck + gosec (required in CI)
make e2e               # generated client vs traust-server at CORE_REF (needs uv)
make guardrails        # AGENTS.md rules a machine can check
```

`make test-integration` uses one fixed local PostgreSQL test DSN, disables the Go test cache so
availability is checked on every run, recreates `traust_storage` for each test, and verifies that
same-named application tables remain untouched. Start the database with:

```bash
podman run --name traust-postgres --rm -d -e POSTGRES_USER=traust -e POSTGRES_PASSWORD=traust-test-only -e POSTGRES_DB=traust_test -p 127.0.0.1:5432:5432 -v traust-postgres-data:/var/lib/postgresql/data docker.io/library/postgres:16
```

These are fixed local test-only credentials, **not deployed credentials or
secrets**, and no DSN override is supported. If PostgreSQL is absent, unreachable,
or rejects authentication, its test reports a skip while SQLite still runs. Once a
connection succeeds, Open, initialization, and protocol failures fail the suite.
