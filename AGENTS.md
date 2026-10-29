# traust-sdk

The Go framework for running traust work. A consumer (SCI, a worker, a CI job) writes an `Executor` and a little config; the SDK runs the lifecycle around it against traust-core's REST API. Go module under `go/`.

## Using the SDK

One client for everything you ask traust; one runner for the work you do for it.

```go
c, err := api.NewBoundClient(api.ClientConfig{BaseURL: coreURL, Tokens: api.FileTokenSource{Path: tokenPath}})
pack, err := ops.LoadPack(packYAML)
r, err := ops.NewRunner(ops.RunnerConfig{
    Client: c, Pack: pack, Executor: myExecutor, Assets: []string{"ledger_audit"},
    LeaseSeconds: 30, PollInterval: time.Second,
})
err = r.Run(ctx)
```

| I want to… | Use | Core operations |
|---|---|---|
| onboard a product and its repositories | `c` | `PlanOnboarding` (dry run) → `Onboard` |
| browse inventory | `c` | `ListProducts`, `ListRepositories`, `ListRefs`, `ListOwners`, `ListProductsOf` |
| run queued fleet commands | `ops.Runner` + your `Executor` | `Claim` → your run → `Complete` |
| record an explicitly dispatched run | `ops.RecordingRun` + your `Executor` | `Start` → `Prepare*` → your run → `Record*` → `Complete` |
| read what a run produced | `c` | `ListRepositoryAssessments`, `GetAssessment`, `GetArtifact`, `GetArtifactContent` |
| work findings (a person's token) | `c` | `ReviewQueue`, `GetFinding`, `Decide`, `ListDecisions`, `Restate` |
| check the ledger | `c` | `ListLayers`, `ListEvents`, `VerifyLayer`, `Checkpoint`; core writes ledger events when records and decisions land, you never append them |
| find exposure | `c` | `GetAdvisory`, `Affected`, `PackagesAffected`, `Dependents`, `Dependencies` |
| react to changes | `feed` + your cursor store | `ReadFeed` |

Method names are core's operation ids. A long or unclear name is fixed in core's route, never by a wrapper here.

## How it works

```mermaid
sequenceDiagram
    participant R as SDK recording run (v1/ops)
    participant API as traust-core REST API
    participant X as consumer Executor
    R->>API: start a dispatched order
    R->>API: prepare → skill input
    R->>X: run(input), heartbeating the lease
    X-->>R: output document
    R->>R: validate against schema
    R->>API: record(document)
    R->>API: complete(JobResult)
```

Each step is one core endpoint. A step core doesn't serve is not a step: never stand one in locally.

| Step | Core endpoint |
|---|---|
| claim / start | `POST /v1/orders/claim` (filtered by `assets`; 204 = nothing queued) · `POST /v1/orders/start` (an explicitly dispatched order) |
| heartbeat | `POST /v1/orders/{order_id}/heartbeat`; a 409 means the lease is lost: stop, record nothing |
| prepare | `GET /v1/assessments/{assessment_id}/triage-input`, and each domain's input route as core adds it |
| record | `POST /v1/assessments/{assessment_id}/triage-records` · `/v1/validation-records` · `/v1/verification-records` · `/v1/assessment-records` |
| complete | `POST /v1/orders/{order_id}/complete` with `JobResult` (outcome, degraded, usage); repeat-safe |
| feed | `GET /v1/feed`, cursor-paged |
| config resolve, preflight | not served by core: no such step |

| Package | Owns | Consumer writes |
|---|---|---|
| `v1/api` | the client, types and enums generated from core's OpenAPI; auth, read retries, typed problem errors, one page iterator | a `TokenSource` |
| `v1/ops` | the runner: claim, heartbeat, lease loss, prepare, validate, record, complete | an `Executor`; skill registrations |
| `v1/feed` | reading `GET /v1/feed` once per event, resuming from a cursor | a cursor store |
| `v1/<pkg>/<pkg>test` | test doubles for each extension point; the conformance suite every `Executor` passes against `traust serve` | — |

Packages not in this table (`storage`, `ledger`, `skills`, `types`, `enums`, `validate`, the contracts generator in `internal/generate`) predate this shape: fix bugs, add nothing.

## Rules

- **Every change is one of three things:** an extension point the SDK calls, a lifecycle step in `ops`, or a regenerated `v1/api`. Anything else is glue; don't add it. A missing endpoint is a traust-core change, not a workaround here
- **Core decides; the SDK enforces.** Authorization, ordering, budget, classification, identity, IDs and timestamps come from core; the SDK never derives, caches or defaults them. Approvals and decisions come from verified people through core; the runner's identity only claims, heartbeats, records and completes. The SDK's one local check is schema validation, which core re-checks
- **Core's REST API is the only way in.** No database handle, DDL, contracts tables, URL building or path splicing: call the generated method
- **Generated, never committed by hand.** `make api-generate` builds core's OpenAPI at the pinned core commit, generates `client.gen.go`, and deletes the spec. Commit the generated Go, the codegen config and the pin, nothing else. Never edit `*.gen.go`; never hand-write a type or enum core already publishes; never wrap generated methods one-for-one
- **Fail closed.** No answer means no run: 5xx, timeout, network failure, malformed body, an unrecognised problem all stop with a typed error. Errors map on HTTP status and, where core sends one, the problem `code`; never on `detail` text. Retry only idempotent reads
- **No fallbacks.** A missing or invalid dependency is a constructor error, never a default. Unknown is `nil`, never a zero value
- **One config.** The consumer passes core URL, credential source and stage identity; nothing else. Library code never reads env vars or files. No settings that choose between behaviours
- **Nothing consumer-specific.** No logic shaped for one consumer and no security policy. Consumers own their infrastructure and their own extension data, joined to traust by id

Before writing Go, load `.agents/skills/go-patterns/SKILL.md`. `.claude/skills` links to `.agents/skills`; edit only `AGENTS.md` and `.agents/skills/`.

## Docs and comments

- Code explains itself. A comment only for a non-obvious *why*, one line. Doc comments 1–2 lines, no parameter lists or narration
- READMEs say how to use what exists. No status, review notes, delivery plans or PR history; those go in the PR
- `go/CHANGELOG.md`: one bullet per consumer-visible change
- No new markdown files
- Public repo: no deployment-specific hostnames, figures or internal URLs (the pre-commit check enforces part of this)

## Workflow

| Lint | Test | Security | Generate | Rules |
|---|---|---|---|---|
| `make lint` | `make test` (network-free); `make test-integration` for storage; `make e2e` against core at `CORE_REF` | `make security` (gosec, govulncheck) | `make -C go generate api-generate`, then `make -C go check-drift` clean | `make guardrails` |

CI runs each of these on every PR. A new `go/v1` package is added to `ci/guardrails.sh` in the same PR.

- A failing test means fix the code. Changing an assertion, deleting a test or adding `t.Skip` needs a person's OK and a **Changed tests** entry in the PR (test · old → new · why)
- Commits: `type(scope): subject`, scope `go` for SDK changes
- Releases: `go/VERSION`, `## [X.Y.Z]` in `go/CHANGELOG.md`, tag `go/vX.Y.Z` (`ci/languages.json`)
- `make setup` once per clone for the hooks; never `--no-verify`. More in [CONTRIBUTING.md](CONTRIBUTING.md)
- Storage integration Postgres (fixed local test-only credentials): `podman run --name traust-postgres --rm -d -e POSTGRES_USER=traust -e POSTGRES_PASSWORD=traust-test-only -e POSTGRES_DB=traust_test -p 127.0.0.1:5432:5432 -v traust-postgres-data:/var/lib/postgresql/data docker.io/library/postgres:16`
