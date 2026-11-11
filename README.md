# traust-sdk

The Go SDK for traust: read its inventory, findings and insights, keep inventory current, and run
work for it, all through traust-core's REST API. Start with [go/README.md](go/README.md).

## Purpose

A consumer imports the SDK, points it at traust-core with a token, and works with traust's data
without tables, schemas or identity logic of its own. traust-core owns the data and every
decision; the consumer keeps only its own extension data, joined to traust by id.

| Layer | Owns |
|---|---|
| [traust-core](https://github.com/traust-security/traust-core) | the data and its services, served over REST |
| traust-sdk (this repo) | the client generated from core's OpenAPI, the `(T, error)` calls and iterators on it, the run lifecycle (`ops`), and a fake core for tests (`traustest`) |
| consumer | its own extension data, keyed by traust ids |

When several consumers build the same thing on top of traust, it belongs in core (a route) or
here (an extension point).

## Setup

```bash
make setup   # enable git hooks (once per clone)
make test              # network-free Go unit tests
make test-integration  # storage: SQLite + optional local PostgreSQL
make status            # per-language versions and tags
```

Each language has its own `{lang}/VERSION` and git tag (`{tag_prefix}/vX.Y.Z` in `ci/languages.json`).

## Language SDKs

| Language | Path | Module |
|---|---|---|
| Go | [`go/`](go/) | `github.com/traust-security/traust-sdk/go` |

## Architecture

```
traust-core   ← data, services, REST API, OpenAPI
traust-sdk    ← generated client + calls + run lifecycle (this repo)
consumer      ← extension data only
```

## Regenerating

```bash
cd go/
make api-generate   # client, calls and pack schema from core at CORE_REF
make generate       # the legacy contracts-generated packages at CONTRACTS_REF
make test           # network-free unit tests
make e2e            # the client and runner against a real traust-server
```

For storage integration coverage, `make test-integration` always exercises SQLite
and attempts PostgreSQL at one fixed local test DSN. Start that database with:

```bash
podman run --name traust-postgres --rm -d -e POSTGRES_USER=traust -e POSTGRES_PASSWORD=traust-test-only -e POSTGRES_DB=traust_test -p 127.0.0.1:5432:5432 -v traust-postgres-data:/var/lib/postgresql/data docker.io/library/postgres:16
```

These are fixed local test-only credentials, **not deployed credentials or
secrets**; the suite does not support a DSN override. An absent, unreachable, or
authentication-rejecting PostgreSQL skips that portion with a message. Once a
connection succeeds, Open, initialization, and protocol failures fail the test.

## License

Apache License 2.0 — see [`LICENSE`](LICENSE).
