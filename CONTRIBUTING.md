# Contributing

## Setup

```bash
git clone <repo-url> && cd traust-sdk
make setup
```

Enables git hooks. One time per clone. Go SDK requires **Go 1.26+** (`toolchain` in `go/go.mod` pins the patched release for `govulncheck`).

## Regenerating

```bash
make -C go api-generate  # the client, its calls and the insight views from core at CORE_REF
make -C go check-drift   # regenerates everything, fails when committed output is stale
```

`api-generate` builds core's OpenAPI at `CORE_REF` (`go/Makefile`), generates
`v1/traust/client.gen.go`, the `(T, error)` calls and `…All` iterators in
`v1/traust/calls.gen.go`, and the insight views in `v1/traust/views.gen.go` and
`v1/traust/insights.gen.go`, and deletes the spec.

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
| `security` | `make security` |
| `drift` | `make -C go check-drift`: generated code matches the pinned core commit |
| `e2e` | `make e2e`: the generated client against `traust-server` built from core at `CORE_REF` |
| `guardrails` | `make guardrails`: the AGENTS.md rules a machine can check, plus commit subjects |

`core-head.yml` runs nightly against core's `development` tip; red means bump `CORE_REF`. Paths in `CODEOWNERS` wait for a code owner.

## Running tests

```bash
make test              # network-free unit tests
make security          # govulncheck + gosec (required in CI)
make e2e               # generated client vs traust-server at CORE_REF (needs uv)
make guardrails        # AGENTS.md rules a machine can check
```
