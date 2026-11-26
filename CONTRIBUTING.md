# Contributing

## Setup

```bash
git clone <repo-url> && cd traust-sdk
make setup
```

Enables git hooks. One time per clone. Go SDK requires **Go 1.26+** (`toolchain` in `go/go.mod` pins the patched release for `govulncheck`).

## Regenerating

```bash
make -C go api-generate     # the client and one calls.gen.go per domain, from core at its pin
make -C go generate         # the example app's templ templates
make -C go check-drift      # regenerates everything, fails when committed output is stale
```

`api-generate` builds core's OpenAPI at the pinned core commit, generates
`v1/traust/client.gen.go`, the `(T, error)` calls and `…All` iterators in one `calls.gen.go`
per domain package under `v1/traust/`, and deletes the spec. Core's OpenAPI is the SDK's only
source of types.

The pins — core, engine and traust, each a full commit sha with a one-line why — live in
`pins.yaml` at the repo root: every language SDK reads that one file and starts the one box
in `box/` (`go/Makefile` reads it today). To bump one, edit its `ref` there and run
`make -C go api-generate check-drift e2e`; commit the regenerated files together with the pin.

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
| `e2e` | `make e2e`: the SDK and the example app end to end in the box — core, the engine's claim worker and the harness — at the pins. Runs only when `pins.yaml` or `box/` changes, or on demand (Actions → e2e → Run workflow) |
| `guardrails` | `make guardrails`: the AGENTS.md rules a machine can check, plus commit subjects |

`core-head.yml` runs nightly against core's `development` tip; red means bump core's ref in `pins.yaml`. Paths in `CODEOWNERS` wait for a code owner.

## Running tests

```bash
make test              # network-free unit tests
make security          # govulncheck + gosec (required in CI)
make e2e               # the SDK and the example app end to end in the box at the pinned core commit (needs uv and cosign)
make guardrails        # AGENTS.md rules a machine can check
```
