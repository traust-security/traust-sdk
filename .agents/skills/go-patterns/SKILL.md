---
name: go-patterns
description: How Go is written in traust-sdk - package layout, extension points, constructors, errors, generated code, concurrency, tests. Load before editing any .go file in this repo.
---

# Go patterns (traust-sdk)

The SDK is a library other teams import: every exported symbol is a promise. Smaller is the goal.

## Layout

```text
go/
├── v1/traust/              generated client (client.gen.go, calls.gen.go) + bindings (client.go); traustest
├── v1/ops/                 the runner and the Executor extension point
├── v1/feed/                feed receiver over GET /v1/feed
├── v1/<pkg>/<pkg>test/     test doubles for <pkg>'s extension points
└── internal/generate/      codegen only; nothing under v1/ imports it
```

- One package per responsibility, named for what it provides. No `util`, `common`, `helpers`, `models`
- File order: `package` → `import` → `const` → `var` → `type` → functions/methods
- `errors.go` holds a package's sentinels and error types; `doc.go` holds the package comment (2–3 lines, no tutorials: those go in the README)
- Everything public is under `v1/`

## Extension points and interfaces

| Need | Do |
|---|---|
| the SDK calls consumer code (`Executor`, `TokenSource`) | interface declared in the package that calls it; 1–3 methods |
| the consumer calls traust | the generated client method; no interface |
| a test needs to fake an extension point | a double in `<pkg>test` (`NewStaticX`, `MemoryX`), shipped for consumers too |
| "might want to swap it later" | nothing: no interface, factory or registry until a second implementation exists |

- Accept interfaces, return concrete types
- Implementations of extension points live with the consumer, not here (except the doubles)

## Constructors and config

- `NewX(required..., opts ...Option) (*X, error)` or `NewX(cfg XConfig) (*X, error)`; pick one per type, never both
- Validate every required dependency: nil or empty → error naming the field (`"api: BaseURL is required"`), never a default
- Defaults only for transport tuning, as named `const`s; anything deployment-specific (URLs, credentials, paths) is passed in
- No package-level mutable state, no `init()` side effects, no `os.Getenv`, no reading files the caller didn't name
- Clock and randomness are injected where behaviour depends on them (`Now func() time.Time`, as in `traust.Config`); IDs and recorded timestamps come from core, never minted here

## Context and concurrency

- `ctx context.Context` is the first parameter of anything that does I/O or blocks; never stored in a struct
- Cancellation is honoured everywhere: every goroutine exits on `ctx.Done()`, and a function that starts one waits for it before returning
- Executor contract: on cancel it stops the work (including child processes), waits, returns

| Need | Pattern |
|---|---|
| parallel ops that can fail | `errgroup.WithContext` |
| heartbeat / periodic work | `time.Ticker` in a goroutine selecting on `ctx.Done()` |
| cancel work when a lease is lost | derived `context.WithCancelCause`; the cause is the typed error |
| shared state | `sync.Mutex`, smallest critical section |
| one-time init | `sync.Once` |

## Errors

- Wrap with context: `fmt.Errorf("recording triage for order %s: %w", id, err)`; what was being done, on what
- Sentinels (`ErrX`) for conditions a caller branches on; typed structs (`*traust.ProblemError`, `*ops.LeaseLostError`) when the caller needs data; both reachable through `errors.Is` / `errors.As`
- Problem responses map once, in `v1/traust`, keyed on HTTP status and, where core sends one, the problem `code`; nothing anywhere matches on `detail` or `err.Error()` text
- Unknown code, malformed body, transport failure → an error, never a guess or a zero value
- No `panic` in library code; no `log` or `fmt.Print` in library code: return the error. A long-running component that must report (the runner) takes a `*slog.Logger` from the consumer
- Never put tokens, credentials or document bodies in errors or logs: IDs, counts, digests

## HTTP and generated code

- All traust calls go through `traust.Client` (auth, retries, problem mapping bound once). No `http.NewRequest`, no URL string building, no path splicing in other packages
- `*.gen.go` is never edited; extend the generated type in a hand-written file in the same package
- Don't wrap generated methods one-for-one. Hand-written code only where the SDK adds behaviour, and generic over the endpoints it serves (one paging helper over `Page`, not one `EachX` per resource)
- Retries: bounded, backoff, idempotent reads only; POSTs are never retried blindly
- The OpenAPI spec is a build input: built or fetched at the pinned core commit during `make api-generate`, then deleted; never committed

## Types and data

- Request and response types are the generated ones; don't re-declare a struct that mirrors one
- No `map[string]any` or `interface{}` across a function boundary; `json.RawMessage` only for opaque bytes passed through unchanged (an executor's report payload)
- Closed sets are the generated enum types, never bare strings
- Unknown or unmeasured is a nil pointer, never `0` or `""`

## Tests

- Table-driven: `tests := []struct{ name string; ...; want ... }{...}` with `t.Run(tt.name, ...)`
- Client behaviour against `httptest.Server`; extension points through the `<pkg>test` doubles
- `go test ./...` is network-free; anything needing a live service is an explicit target, never on by default
- Cover the edges, not only the happy path: nil / missing dependency, cancellation mid-run, lease loss, each problem code, unknown code, 5xx, timeout, malformed body, empty page, re-run
- Test names say the behaviour they protect
- A failing test means fix the code. Changing an assertion, deleting a test or adding `t.Skip` needs a person's OK and a line under **Changed tests** in the PR

## Security

- `make security` (gosec, govulncheck) is clean; `#nosec` only with a one-line why on the same line
- No secrets in code, tests or fixtures; test data is sanitized (no real hosts, orgs, repos, people, tokens)

## Checklist

- Is each new exported symbol an extension point or backed by one core endpoint? If neither, delete it
- Required dependencies validated; no default standing in for a missing one?
- `ctx` first; every goroutine exits on cancel?
- Errors wrapped with what/on-what; typed where a caller branches; no string matching?
- No `os.Getenv`, global state, `time.Now()` in logic, hand-built URLs?
- No edit to `*.gen.go`, no committed spec, no one-for-one wrapper?
- Comments only for a *why*; doc comments 1–2 lines?
- Table-driven tests covering the edges; `go test ./...` network-free?
- `make lint`, `make test`, `make security`, `make check-drift` clean?
