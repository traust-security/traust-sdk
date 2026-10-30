# Contributing

## Setup

```bash
git clone <repo-url> && cd traust-sdk
make setup
```

Enables git hooks. One time per clone. Go SDK requires **Go 1.26+** (`toolchain` in `go/go.mod` pins the patched release for `govulncheck`).

Code generation fetches the canonical `traust-contracts` repository at the immutable commit pinned in `go/Makefile`; no sibling checkout is required. The verified checkout is reused from the user cache. See [README.md](README.md) for repository/ref/cache overrides.

## Commit messages

Conventional commits required. Format: `type(scope): subject`

Types: `feat`, `fix`, `perf`, `refactor`, `docs`, `test`, `chore`, `ci`, `build`, `style`, `revert`

Use the language id as scope when the change is SDK-specific:

```
feat(go): add harness provider interface
fix(go): correct enum code generation
ci: update shared gates              ← no language release required
```

## Hooks

| Hook | Runs | Speed |
|------|------|-------|
| `commit-msg` | subject format check | instant |
| `pre-commit` | gofmt on staged `.go` | <1s |
| `pre-push` | `go vet` + `go test ./...` | ~10s |

Bypass: `--no-verify` on commit or push. CI still enforces.

## Releasing

Each language versions independently. Tag format: `{tag_prefix}/vX.Y.Z` (see `ci/languages.json`).

If your PR has release-worthy commits (`feat`, `fix`, `perf`, or `!`) **for a language**:

```bash
make bump go minor
# add to go/CHANGELOG.md:
#   ## [X.Y.Z]
#   ### Added
#   - ...
make check-release
git add go/VERSION go/CHANGELOG.md
git commit -m "chore(go): release X.Y.Z"
```

Only bump languages touched by the PR.

CI/infra-only PRs: use `ci:` / `chore:` commits — no VERSION bump.

Registry: `ci/languages.json`.

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

`core-head.yml` runs nightly against core's `development` tip; red means bump `CORE_REF`. `review-after-merge.yml` follows traust-core's merge-by-risk flow; paths in `CODEOWNERS` wait for a code owner.

## Running tests

```bash
make test              # network-free unit tests
make test-integration  # SQLite plus optional local PostgreSQL storage coverage
make security          # govulncheck + gosec (required in CI)
make -C go test
make -C go check-drift # regenerates from the pinned contracts and core commits
make e2e               # generated client vs traust-server at CORE_REF (needs uv)
make guardrails        # AGENTS.md rules a machine can check
```

`make test-integration` uses one fixed local PostgreSQL test DSN. Start it with:

```bash
podman run --name traust-postgres --rm -d -e POSTGRES_USER=traust -e POSTGRES_PASSWORD=traust-test-only -e POSTGRES_DB=traust_test -p 127.0.0.1:5432:5432 -v traust-postgres-data:/var/lib/postgresql/data docker.io/library/postgres:16
```

These are fixed local test-only credentials, **not deployed credentials or
secrets**, and no DSN override is supported. If PostgreSQL is absent, unreachable,
or rejects authentication, its test reports a skip while SQLite still runs. Once a
connection succeeds, Open, initialization, and protocol failures fail the suite.

## Architecture

See [README.md](README.md) for module layout and contract regeneration.
