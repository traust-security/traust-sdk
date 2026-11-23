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

Hooks, tests, generation and the other gates for working on this repo:
[CONTRIBUTING.md](CONTRIBUTING.md).

## Language SDKs

| Language | Path | Module |
|---|---|---|
| Go | [`go/`](go/) | `github.com/traust-security/traust-sdk/go` |

## License

Apache License 2.0 — see [`LICENSE`](LICENSE).
