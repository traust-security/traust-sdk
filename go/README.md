# traust-sdk (Go)

Program against traust: read its inventory, exposure, findings and insights, keep inventory
current, and submit and follow its work. Every call goes to traust-core's REST API; the SDK
handles auth, retries, paging and typed errors. One package per domain carries the calls; the
root keeps the transport, `Poller`, the errors, the generated models and `Whoami`/`Healthz`.
`go doc` on a package shows every signature; the examples run.

```bash
go get github.com/traust-security/traust-sdk/go
```

## Connect

```go
client, err := traust.NewClient(traust.Config{
    BaseURL: "https://traust-core.example.com",
    Tokens:  traust.FileTokenSource{Path: "/var/run/secrets/traust-core/token"},
})
url := "https://github.com/example/api"
repos := inventory.New(client)
page, err := repos.ListRepositories(ctx, &traust.ListRepositoriesParams{Url: &url})
```

`traust.FileTokenSource{Path}` for services — re-reads the file on every request, so rotated
tokens just work — and `traust.StaticToken("…")` for scripts, local development and tests.

`Config` also takes your own `*http.Client`, the retry count and the backoff. Reads (GET)
retry on 429, 502, 503, 504 and transport errors; writes are never retried.

`traust.LoadConfig` builds the `Config` from a profile in `~/.traust/config.yaml` — the one
`traust.Profile("name")` names, else `TRAUST_PROFILE`, else the file's default; `TRAUST_URL`
and `TRAUST_TOKEN_FILE` override it. A profile names exactly one credential: a `token_file`,
or a machine identity — `token_url`, `client_id`, `client_secret_file` (a leading `~/`
expands, optional `scopes`) — the secret is read from its file, never the config.

A BFF keeps one client under its own identity and sends each user's calls as that user, so
core authorizes and attributes them to the person:

```go
ctx := traust.OnBehalfOf(r.Context(), userToken) // an empty token is an error, never a fallback
order, err := runs.New(client).Submit(ctx, req)
```

## inventory

Products, repositories, refs and owners — core's portfolio and registry routes. Writes are
idempotent, and core stamps when ownership opens and closes, so "who owned this on date X"
stays answerable.

```go
inv := inventory.New(client)
inv.Onboard(ctx, traust.OnboardingRequest{Product: "api"})       // product, repositories, refs, teams
inv.RegisterRef(ctx, repoID, traust.RefRegistration{Sha: sha})    // a commit you're about to scan
inv.SetOwners(ctx, repoID, traust.Owners{Teams: []string{"red"}}) // the teams your source reports
```

Every paged listing also has an `…All` iterator that follows the cursor and stops at the
first error: `repos.ListRepositoriesAll(ctx, nil)`.

## runs

Orders: submit, claim, run and complete work, and query history — core's ops routes. A worker
only claims, heartbeats, records and completes; approvals come from verified people through
core.

```go
poller, err := runs.New(client).BeginSkill(ctx, "scan") // submit a run by asset name and follow it
```

`QueryOrders(ctx, traust.OrderQueryRequest{…})` posts filters, search, sort and bounded paging
to `POST /v1/orders/query`; each item carries an opaque cursor for resuming with filters and
sort unchanged — not a transactional snapshot. `QueryOrdersAll` iterates every matching item,
stopping when the consumer stops. As a POST, `QueryOrders` is not retried.

## findings

Findings, assessments, artifacts and validation, and the decisions people make on them.

```go
fs := findings.New(client)
finding, err := fs.GetFinding(ctx, findingID)                  // one resource
for f, err := range fs.ListFindingsAll(ctx, repoID, nil) { … } // every page
```

People `Decide` and `Restate`; workers `RecordTriage` and `RecordAssessment`.

## remediation

Remediation efforts, work items and their verification.

```go
rem := remediation.New(client)
page, err := rem.ListWorkItems(ctx, repoID, nil) // …All iterates every page
_, err = rem.RecordVerification(ctx, traust.VerificationRecordBody{…})
```

## exposure

Advisories, packages, dependencies and the reach of both.

```go
exp := exposure.New(client)
advisory, err := exp.FindAdvisory(ctx, &traust.FindAdvisoryParams{Identifier: "GHSA-xxxx-yyyy-zzzz"})
rows, err := exp.SummarizeExposure(ctx, nil) // rollups per dimension; ExposureTimeline for series
```

## ledger

Ledger layers, events and their verification.

```go
led := ledger.New(client)
layers, err := led.ListLayers(ctx, repoID)
result, err := led.VerifyLayer(ctx, layerID) // an admin may Checkpoint
```

## insight

Dashboards and reports. One query counts what you ask, broken down by up to three dimensions,
over the repositories (or products, teams) the person may see; rows read by typed key:

```go
result, err := insight.New(client).QueryDashboard(ctx, traust.DashboardQuery{
    Subject:  traust.SubjectFindings,
    Measures: []traust.Measure{traust.MeasureFindings},
    GroupBy:  &[]traust.Dimension{traust.DimensionSeverity},
})
for _, row := range result.Rows {
    severity, count := insight.Group(row, traust.DimensionSeverity), insight.Measure(row, traust.MeasureFindings)
}
```

Counts come from `QueryDashboard`, `SummarizeFindings` and `SummarizeExposure`, never from
listings.

A report is a job core builds; follow the build, then download the file:

```go
format := traust.ReportFormatMarkdown
reports := insight.New(client)
poller, err := reports.BeginRequestReport(ctx, traust.ReportRequest{Report: traust.FindingsRollup, Format: &format})
report, err := poller.Wait(ctx, 2*time.Second)          // ready or failed; ResumeReport(token) after a restart
file, err := reports.GetReportFile(ctx, report.OrderId) // the stored bytes, whatever the format
```

## Errors

A refusal matches its status with `errors.Is`, and core's deny code matches too; the problem
document itself is a `*traust.ProblemError`, a transport failure is the underlying error.

```go
errors.Is(err, traust.ErrNotFound)     // also ErrUnauthorized, ErrForbidden, ErrConflict, ErrInvalid
errors.Is(err, traust.BudgetExhausted) // a DenyCode core sent, as the problem's code or an issue's
if p, ok := errors.AsType[*traust.ProblemError](err); ok { log.Print(p.Problem.RequestId) }
```

## Poller

`traust.Poller[T]` follows work core finishes later: `Wait(ctx, every)` returns the finished
thing or its failure, `Poll` checks once, `ResumeToken` persists across a restart. The domain
packages build one for orders (`runs.BeginSubmit`, `ResumeOrder`) and reports
(`insight.BeginRequestReport`, `ResumeReport`); core's `finished_at` says when it is done.

## The example app

[`examples/webapp`](examples/webapp) is a small web app on traust (net/http, templ, htmx),
one package per feature. From a blank core: onboard repositories (with a dry-run preview),
scan and triage one, settle the review queue, decide and restate findings, open and follow
remediation efforts, read advisories and exposure, watch the dashboard, verify (and, as an
admin, checkpoint) the ledger, and build and download a report. It acts as each signed-in
person with `OnBehalfOf`, submits runs by asset name with `BeginSkill`, and follows runs and
report builds with the `Poller`; the box's claim worker, not the app, runs the work.

`make webapp` runs the whole stack locally: traust-server beside its report builder (core at
`CORE_REF`), traust-engine planning the shipped security pack's work every minute, and the
engine's claim worker running each order — the pack's deterministic assets as materializers
and its llm assets as skill sessions in the traust harness, which needs a Claude login. Sign
in as the person it prints and onboard any public repository; the engine's next pass scans
and triages it. It needs `uv`, `cosign` (the local ledger is signed), `gitleaks` and a Claude
login. `make webapp STORE=byos` brings its own storage the way a deployment does, Postgres
for records and an S3 store for artifacts, in throwaway containers (needs `podman`).

`make e2e` runs the SDK's and the app's suites end to end through the same box: onboarding
reads a real public GitHub repository (needs network; a `GITHUB_TOKEN` in the environment, if
any, lifts its rate limit) and a stub claude runs canned work, so no Claude login is needed.

`make webapp-compose` runs the same box as containers: `podman compose` brings up Postgres
(throwaway), traust-server and its reports role, the engine's planner and claim worker (the
traust harness and the claude CLI in its image), and the app, all built from the pinned sources.
It needs only podman and, for live scans, a Claude login in `CLAUDE_RUNNER` (or
`ANTHROPIC_API_KEY`); `podman compose -f examples/webapp/compose.yaml down -v` tears it down.

## Test your code

`traustest` is a fake core: register the responses your code should see and check what it sent.

```go
core := traustest.New(t)
core.Handle("GET /v1/repositories/{id}", traustest.JSON(200, traust.Repository{Url: "https://github.com/example/api"}))
core.Handle("POST /v1/orders/claim", traustest.Status(204))
client := core.Client() // … exercise your code …
for _, r := range core.Requests() { … }
```

An unregistered route answers with a 404 problem. Use `traustest.Start()` (and `Close`)
outside a test, as the runnable examples do.
