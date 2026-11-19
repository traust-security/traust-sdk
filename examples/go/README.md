# The Go example app

The SDK's reference consumer, in the spirit of Okta's and AWS's example apps:
a small runnable program that uses the SDK the way an adopter would, so the
team feels every gap before designing around it.

It does three things, end to end against a local `traust-server`:

1. **runs one asset** — onboards the fixture repository, submits an order
   pinned to `pack.yaml`, and runs it to completion on a worker's runner
2. **reads one dashboard screen** — live counts for its repository
3. **downloads one report** — the findings-rollup view, rendered as markdown

## Run it

```sh
make -C go example
```

That target builds core at `CORE_REF`, starts `traust-server` with a local
identity issuer, stands up the e2eforge stand-in, mints a person and a worker
token, and runs the three flows. Against your own server instead:

```sh
TRAUST_CORE_URL=https://core.example.com \
TRAUST_TOKEN_FILE=person.jwt \
TRAUST_RUNNER_TOKEN_FILE=worker.jwt \
go run .
```

## The gap log

Every place this example had to build machinery the SDK should own is marked
`GAP-n` in `main.go` and explained in [GAPLOG.md](GAPLOG.md), each with the
appliance or read-side step that closes it. The example is deliberately
written twice: this version against the SDK as it exists today (155 lines —
the baseline), and again after the appliance lands — every line the second
version keeps should be a domain decision, everything else the SDK's job.
