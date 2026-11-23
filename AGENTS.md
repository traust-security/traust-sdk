# traust-sdk

The Go SDK for traust: it does the work for the user, like a cloud provider's SDK. Every core operation is a typed call, every listing an iterator, and the run lifecycle is the SDK's; a consumer (SCI, a worker, a CI job) writes an `Executor` and a little config. Go module under `go/`; the user guide is `go/README.md`.

## Rules

- **Every change is one of four things:** an extension point the SDK calls, a lifecycle step in `ops`, a regenerated `v1/traust`, or a helper that saves every consumer the same work (`ProcessExecutor`, `traustest`, a generated iterator). Anything else is glue; don't add it. A missing endpoint is a traust-core change, not a workaround here
- The legacy packages — `storage`, `ledger`, `skills`, `types`, `enums`, `validate`, and the contracts generator in `internal/generate` — predate this shape: fix bugs, add nothing
- **Core decides; the SDK enforces.** Authorization, ordering, budget, classification, identity, IDs and timestamps come from core; the SDK never derives, caches or defaults them. Approvals and decisions come from verified people through core; the runner's identity only claims, heartbeats, records and completes. The SDK's one local check is schema validation, which core re-checks
- **Core's REST API is the only way in.** No database handle, DDL, contracts tables, URL building or path splicing: call the generated method. A step core doesn't serve is not a step: never stand one in locally
- **Generated, never committed by hand.** Everything the SDK knows about core comes from core at `CORE_REF`: `make api-generate` builds its OpenAPI, generates `client.gen.go` and, from it, the `(T, error)` calls and `…All` iterators in `calls.gen.go` (`internal/generate/calls`), deletes the spec, and writes `v1/ops/pack.schema.json` from the contracts core validates packs with. Commit those generated files and the pin, nothing else; there is no second contracts pin. Never edit `*.gen.go`; never hand-write a type or enum core already publishes, or a method per endpoint: a better call shape is a change to the generator, and a long or unclear name is fixed in core's route, never by a wrapper here
- **Fail closed.** No answer means no run: 5xx, timeout, network failure, malformed body, an unrecognised problem all stop with a typed error. Errors map on HTTP status and, where core sends one, the problem `code`; never on `detail` text. Retry only idempotent reads
- **No fallbacks.** A missing or invalid dependency is a constructor error, never a default. Unknown is `nil`, never a zero value
- **One config.** The consumer passes core URL, credential source and stage identity; nothing else. Library code never reads env vars or files, except `traust.LoadConfig`, the one explicit loader the consumer calls. No settings that choose between behaviours
- **Nothing consumer-specific.** No logic shaped for one consumer and no security policy. Consumers own their infrastructure and their own extension data, joined to traust by id

Before writing Go, load `.agents/skills/go-patterns/SKILL.md`. `.claude/skills` links to `.agents/skills`; edit only `AGENTS.md` and `.agents/skills/`.

## Docs and comments

- Code explains itself. A comment only for a non-obvious *why*, one line. Doc comments 1–2 lines, no parameter lists or narration
- READMEs say how to use what exists. No status, review notes, delivery plans or PR history; those go in the PR
- No new markdown files
- Public repo: no deployment-specific hostnames, figures or internal URLs (the pre-commit check enforces part of this)

## Workflow

| Lint | Test | Security | Generate | Rules |
|---|---|---|---|---|
| `make lint` | `make test` (network-free); `make test-integration` for storage; `make e2e` against core at `CORE_REF` | `make security` (gosec, govulncheck) | `make -C go generate api-generate`, then `make -C go check-drift` clean | `make guardrails` |

CI runs each of these on every PR. A new `go/v1` package is added to `ci/guardrails.sh` in the same PR.

- A failing test means fix the code. Changing an assertion, deleting a test or adding `t.Skip` needs a person's OK and a **Changed tests** entry in the PR (test · old → new · why)
- Commits: `type(scope): subject`, scope `go` for SDK changes
- `make setup` once per clone for the hooks; never `--no-verify`. More in [CONTRIBUTING.md](CONTRIBUTING.md)
- Storage integration runs against a local PostgreSQL with fixed test-only credentials; [CONTRIBUTING.md](CONTRIBUTING.md#running-tests) starts it
