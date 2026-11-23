# traust-sdk (Go)

Program against traust: read its inventory, findings, dashboards and reports, keep inventory
current, and submit and follow its work. Every call goes to traust-core's REST API; the SDK
handles auth, retries, paging and typed errors.

```bash
go get github.com/traust-security/traust-sdk/go
```

## Quickstart

```go
import "github.com/traust-security/traust-sdk/go/v1/traust"

client, err := traust.NewClient(traust.Config{
    BaseURL: "https://traust-core.example.com",
    Tokens:  traust.FileTokenSource{Path: "/var/run/secrets/traust-core/token"},
})

url := "https://github.com/example/api"
repos, err := client.ListRepositories(ctx, &traust.ListRepositoriesParams{Url: &url})
repo := repos.Items[0]

counts, err := client.SummarizeFindings(ctx, *repo.Id,
    &traust.SummarizeFindingsParams{GroupBy: traust.FindingDimensionSeverity})
```

Every operation is a method named after core's operation — `client.GetRepository(ctx, id)` returns
`(*traust.Repository, error)`; `go doc ./v1/traust` shows every signature and the runnable examples.

## Connect

| Token | Use |
|---|---|
| `traust.FileTokenSource{Path}` | services: re-reads the file on every request, so rotated tokens just work |
| `traust.StaticToken("…")` | scripts, local development and tests |

`Config` also takes your own `*http.Client`, the retry count and the backoff. Reads (GET)
retry on 429, 502, 503, 504 and transport errors; writes are never retried.

A `LoadConfig` profile can name a machine identity instead of a `token_file`: `token_url`, `client_id`,
`client_secret_file` (a leading `~/` expands), optional `scopes` — the secret is read from the file, never the config.

## Read

```go
finding, err := client.GetFinding(ctx, id)            // one resource
page, err := client.ListFindings(ctx, repoID, nil)     // one page
for f, err := range client.ListFindingsAll(ctx, repoID, nil) { … } // every page
```

Each paged listing has an `…All` iterator that follows the cursor and stops at the first error. Counts and dashboards
come from `QueryDashboard`, `SummarizeFindings` and `SummarizeExposure`, reports from `BeginRequestReport` — never from listings.

### Dashboards

One query answers a dashboard card: what to count, broken down by up to three dimensions, over
the repositories (or products, teams) the person may see. Rows read by typed key:

```go
result, err := client.QueryDashboard(ctx, traust.DashboardQuery{
    Subject:       traust.SubjectFindings,
    Measures:      []traust.Measure{traust.MeasureFindings},
    GroupBy:       &[]traust.Dimension{traust.DimensionSeverity},
    RepositoryIds: &repoIDs,
})
for _, row := range result.Rows {
    severity, count := row.Group(traust.DimensionSeverity), row.Measure(traust.MeasureFindings)
}
```

### Reports

A report is a job core builds; follow the build, then download the file. A `json` report is the
view itself, and `ReportView` decodes it as its generated `…Insight` type, refusing a field
outside the contract:

```go
format := traust.ReportFormatMarkdown
poller, err := client.BeginRequestReport(ctx, traust.ReportRequest{Report: traust.FindingsRollup, Format: &format})
report, err := poller.Wait(ctx, 2*time.Second)         // ready or failed; ResumeReport(token) after a restart
file, err := client.GetReportFile(ctx, report.OrderId) // the stored bytes, whatever the format

rollup, err := traust.ReportView[traust.FindingsRollupInsight](ctx, client, report) // a ready json report
```

A view composed from others uses their types: `ExecutiveSummaryInsight.Rollup` is a
`FindingsRollupInsight`, its `Pqc` a `*PqcPortfolioInsight` (nil when the panel was not asked
for), its isolation panel's `ThreatRegister` a `ThreatRegisterInsight`, and so on. Shared
contract definitions generate once and are reused: the finding records views embed are
`FindingRecordsInsightFindingView`, `FindingRecordsInsightOccurrence`,
`FindingRecordsInsightDecision` (`ContainerFindingsInsight` rows carry them).

### Query orders

`client.QueryOrders(ctx, traust.OrderQueryRequest{…})` posts filters, search, sort and bounded
paging to `POST /v1/orders/query`. Each item contains the order, its repository URL (nullable)
and an opaque cursor. Follow the page's `Next` for a whole page, or the last consumed item's
`Cursor` when merging only part of a page with another source. Keep filters and sorting
unchanged when resuming; the limit may change (1–1000). Malformed or mismatched cursors return
a typed validation problem.

`q` is a literal case-insensitive substring search across the displayed repository reference
(URL, else repository ID, else order ID), asset and state. Normal sorts use the requested
direction, then order ID, with nulls last. `sort_by=ids` preserves the supplied distinct ID
sequence regardless of direction; `ids=null` is unrestricted and `ids=[]` matches nothing.
Each ID selector accepts at most 500 values; oversized requests fail validation
rather than silently dropping IDs. Live mutable sort/filter values can change
between pages: these cursors are not transactional snapshots.
Use these server-side queries rather than walking order history to filter or sort it.
As a POST, `QueryOrders` is not automatically retried.

`client.QueryOrdersAll(ctx, request)` iterates the remaining matching items
from `request.Cursor`, following page boundaries and yielding the first error.
It stops fetching when the consumer stops. Use it for intentional full-history
processing; interactive tables and mixed-source merges should keep using
bounded `QueryOrders` pages rather than draining this iterator before rendering.

## Keep inventory current

```go
client.Onboard(ctx, traust.OnboardingRequest{…})                     // product, repositories, refs, teams
client.RegisterRef(ctx, repoID, traust.RefRegistration{Sha: sha})     // a commit you're about to scan
client.SetOwners(ctx, repoID, traust.Owners{Teams: []string{"red"}})  // the teams your source reports
```

Writes are idempotent: re-sending the same request changes nothing. Core stamps when ownership
opens and closes, so "who owned this on date X" stays answerable.

## Errors

A refusal matches its status with `errors.Is`, and core's deny code matches too. For the problem
document itself, take the `*traust.ProblemError`; a transport failure is the underlying error.

```go
errors.Is(err, traust.ErrNotFound)     // also ErrUnauthorized, ErrForbidden, ErrConflict, ErrInvalid
errors.Is(err, traust.BudgetExhausted) // a DenyCode core sent, as the problem's code or an issue's
if p, ok := errors.AsType[*traust.ProblemError](err); ok { log.Print(p.Problem.RequestId) }
```

## Act for a person

A service serving many users (a BFF) keeps one client under its own identity and sends each
user's calls as that user, so core authorizes and attributes them to the person:

```go
ctx := traust.OnBehalfOf(r.Context(), userToken) // an empty token is an error, never a fallback
order, err := client.Submit(ctx, req)
```

## Example app

[`examples/webapp`](examples/webapp) is a small web app on traust (net/http, templ, htmx). From a
blank core: onboard repositories (with a dry-run preview), scan one and triage it, settle what the
triage leaves in the review queue, decide and restate findings, watch the dashboard count it all,
verify (and, as an admin, checkpoint) each repository's ledger, and build and download a report. It acts as each signed-in person
with `OnBehalfOf`, submits runs by asset name with `BeginSkill`, follows runs and report builds
with `Poller`; the box's claim worker, not the app, runs the work.

`make webapp` runs the whole stack locally: traust-server (reaching real forges), its report
builder, and traust-engine, which plans a scan of each repository under the app's pack. Sign in
as the person it prints (or the admin, who may restate decisions) and onboard any public
repository; the engine's next pass scans it, then triages it.
It needs `uv` and `cosign` (the local ledger is signed). `make webapp STORE=byos` brings its own
storage the way a deployment does, Postgres for records and an S3 store for artifacts, in throwaway
containers (needs `podman`). `make e2e` runs the same flows end to end against a forge stand-in.

The box runs the pack's real skills already: each order runs the pack's command, the traust
harness's skill session, which needs a harness checkout and a logged-in agent on the worker.

## Test your code

`traustest` is a fake core: register the responses your code should see and check what it sent.

```go
core := traustest.New(t)
core.Handle("GET /v1/repositories/{id}", traustest.JSON(200, traust.Repository{Url: "https://github.com/example/api"}))
core.Handle("POST /v1/orders/claim", traustest.Status(204))
client := core.Client()
// … exercise your code …
sent := core.Requests()
```

An unregistered route answers with a 404 problem. Use `traustest.Start()` (and `Close`) outside a test,
as the runnable examples do. `make e2e` runs the SDK itself against a real `traust-server`.

