# traust-core REST API clients

Generated Go clients for traust-core's REST adapter (`traust-server`), plus
hand-written bindings in `client.go` (bearer-token auth from a rotating file,
bounded retries on transient GET failures, typed problem+json errors, cursor
pagination).

## Input

`openapi.json` is traust-core's `make api-check` output, checked in so
generation is reproducible:

- **traust-core commit:** `38a165e` (development, 2026-10-28) — 37 paths,
  67 schemas (assessments, exposure, findings, health, identity, ledger,
  ops, portfolio)
- Regenerate: check out that traust-core, run `make api-check`, copy
  `.openapi/openapi.json` here, bump the commit above, then `make
  api-generate` and commit the result.

`normalize.py` rewrites the two 3.1-isms kin-openapi cannot read (numeric
`exclusiveMinimum`/`exclusiveMaximum`; pydantic's `anyOf: [schema,
{type: "null"}]` Optional) into their 3.0 forms, producing
`openapi-3.0.json`; the document is otherwise untouched. When core's
OpenAPI publishing lands with each release, this directory tracks the
released artifact instead.

## Generating

```sh
make -C .. api-generate   # from go/, or: make api-generate from go/
```

runs `normalize.py` then `go tool oapi-codegen` (pinned in go.mod's tool
directives) over `oapi-codegen.yaml`. `client.gen.go` is generated; never
edit it — extend `client.go` instead. CI's drift check regenerates and fails
on a diff.
