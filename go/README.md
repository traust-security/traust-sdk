# traust-sdk (Go)

```bash
go get github.com/traust-security/traust-sdk/go@v0.9.0
```

## Work routing: primary scan selection

`work.DecidePrimary(PrimaryInput)` selects primary scan work for one repository
from supplied inventory/audit membership, observation status, measurements, and
already-correlated events. It applies event precedence, missing-SHA and
quota-deferred guards, the existing change table, and never-audited bootstrap:

```go
import "github.com/traust-security/traust-sdk/go/v1/routing/work"

result, err := work.DecidePrimary(work.PrimaryInput{
    Inventory: &work.InventoryInput{Designation: work.DesignationExternal},
    // Audit is nil: this repository is known to have no HEAD code-audit baseline.
})
if err != nil {
    return err
}
// result.Decisions[0]: rule-1-bootstrap, full-audit, P2, private-external
```

The caller obtains facts from its authoritative views, deduplicates audit
baselines, resolves event URL/key aliases and designation rules, and collects
comparison measurements. Nil `Audit`/`Inventory` means known population absence;
an unavailable view must be reported as a collection failure instead. An existing
audit without a recoverable SHA remains an `Audit` with `StatusNoPinnedSHA`.

`EventsHonored` records recognized, unconsumed events, including those without an
audited target or without a queued lane. Unknown sources and consumed events are
ignored. Indexes refer to the original input event slice; multiple queued events
retain their order and suppress ordinary table selection. A nonqueued event does
not suppress it. The SDK never consumes events or mutates input facts. Preserve
the original event payloads for subsequent companion processing.

`Audit.Change == nil` distinguishes an unavailable comparison from measured zero.
Explicit observation statuses are required. A changed repository with `StatusOK`
needs comparison measurements unless a queued event supersedes the table. Other
failure statuses retain Python's age-rule fallback and remain visible on every
decision; `StatusQuotaDeferred` and `StatusNoPinnedSHA` take their dedicated paths.
Validation rejects invalid supplied measurements even if an event would match.

### Lower-level change table

`v1/routing/work` ports Traust's repository risk classification and ordered rescan
decision table into Go. It accepts normalized measurements and returns the matched
rule, lane, and explanation without I/O or dispatch:

```go
import "github.com/traust-security/traust-sdk/go/v1/routing/work"

decision, err := work.EvaluateTable(work.TableInput{
    RiskTier:     work.RiskP1,
    Exposure:     work.ExposurePrivateInternal,
    ChangedLines: 8000,
})
if err != nil {
    return err
}
// decision.Rule == "rule-2", decision.Lane == work.LaneFullAudit
```

Use `ClassifyRisk(RiskInput)` to derive the risk band from the live critical/high
finding count and archival/dormancy state. `P0MinLive` mirrors Python's
`P0_MIN_LIVE` policy constant (5); it is not a runtime configuration setting.
Invalid inputs return `*InputError`.
Optional ratio, audit-age, and commits-ahead pointers distinguish unknown values
from measured zero. Other zero/false measurements must represent collected data.

**This is not the complete work router.** `DecidePrimary` handles the primary
event/bootstrap/comparison flow; `EvaluateTable` evaluates only the ordinary table.
IaC/threat-model companion work, tripwire/refusal pre-routing, fleet ordering,
quarterly drain selection, and budget/model gates remain separate worklist stages.
A selected lane is not dispatch authorization; `diff-scan-quarterly` belongs to a
deferred pool and `release-passthrough` hands off to branch/container/RPM processing.
Release-driven threat-model companion work remains a separate stage. No configuration
file, scheduler, or execution provider is introduced by this package. Its typed
inputs are Go call parameters, not a new published JSON artifact contract.

The fixed policy matches
[Traust's Python implementation at the pinned revision](https://github.com/traust-security/traust/blob/0f95c95f48d36509b70234d18b500c0b738e23e8/src/traust/cli/build_rescan_worklist.py).
Offline tests compare full decisions with outputs from that Python source. To
regenerate those fixtures, supply the source file from that exact revision:

```bash
python3 v1/routing/work/testdata/generate.py /path/to/build_rescan_worklist.py
python3 v1/routing/work/testdata/generate_primary.py /path/to/build_rescan_worklist.py
go test ./v1/routing/work
```

Both generators verify the source SHA-256. The first evaluates pure reference
definitions; the primary generator executes the original event/comparison loops
and bootstrap function using temporary event/graph files. Companion stages are
excluded, and canonical identities are supplied. Normal Go tests need neither
Python nor a Traust checkout. The fixtures cover selection, not identity matching,
measurement collection, fleet ordering, or the complete CLI pipeline.

## Skills SDK

Import a skill, plug in your provider, call it — typed input in, typed result out:

```go
import "github.com/traust-security/traust-sdk/go/v1/skills"

provider := myK8sProvider(...)  // you implement skills.Provider

artifact, err := skills.Scan.Run(ctx, provider, skills.ScanInput{
    Repo: "https://github.com/org/repo",
    Ref:  "main",
})
report, err := artifact.Value()
if err != nil {
    return err
}
for _, finding := range report.Findings {
    fmt.Println(finding.Id, finding.Severity, finding.Title)
}
```

`skills.Provider` is the only interface you implement — it runs a skill somewhere
(K8s Job, Docker, local subprocess, ...) and returns raw bytes:

```go
type Provider interface {
    Execute(ctx context.Context, skill SkillMeta, input []byte) ([]byte, error)
}
```

The SDK validates the response and retains its exact bytes with the generated
Go type. `Artifact.Value` decodes typed data; `Artifact.Payload` returns owned,
byte-identical evidence suitable for storage or export.

### Calling multiple skills

Bind the provider once with a Client:

```go
client := skills.NewClient(provider)
report, _ := client.Scan(ctx, skills.ScanInput{...})
triage, _ := client.Triage(ctx, skills.TriageInput{...})
```

### Testing

```go
import "github.com/traust-security/traust-sdk/go/v1/skills/skillstest"

provider := skillstest.NewStaticProvider().
    WithScanResult(skillstest.FixtureReport())

report, _ := skills.Scan.Run(ctx, provider, skills.ScanInput{
    Repo: "https://github.com/org/repo",
    Ref:  "main",
})
```

## Ledger SDK

Submit reports and human events to the ledger service:

```go
import "github.com/traust-security/traust-sdk/go/v1/ledger"

client := ledger.NewHTTPClient("https://ledger.example.com",
    ledger.WithBearerToken(os.Getenv("LEDGER_TOKEN")),
)

// Machine lane — triage report (derives disposition events)
resp, err := client.SubmitTriageReport(ctx, ledger.TriageReportInput{
    ReportMeta: ledger.ReportMeta{
        LayerID:    "repo-a",
        SourceRef:  "findings/repo-a/triage.json",
        RecordedAt: time.Now().UTC().Format(time.RFC3339),
    },
    Report: triageReport,
})

// Human lane — countersign
resp, err = client.SubmitCountersign(ctx, ledger.CountersignInput{
    EventMeta: ledger.EventMeta{
        LayerID:    "repo-a",
        RecordedAt: time.Now().UTC().Format(time.RFC3339),
    },
    FindingRef:    "FIND-001",
    Verdict:       "true_positive",
    Justification: "Reviewed source and confirmed exploit path.",
})

// Sign a layer's merkle tree
signResp, err := client.SignLayer(ctx, "repo-a", ledger.SignOpts{})
```

Machine-lane report kinds: `triage`, `validation`. Human-lane event kinds:
`countersign`, `severity`. Additional operations: `BatchSubmit`,
`ResolveReviewItem`, `ComputeFingerprints`, `StampEventIdentities`, and
`SignLayer`.

The built-in HTTP transport handles envelope formatting, lane routing, and auth.
For custom transports, implement `ledger.Provider` and use `ledger.NewClient(provider)`.

### Testing

```go
import "github.com/traust-security/traust-sdk/go/v1/ledger/ingesttest"

provider := ingesttest.NewStaticProvider().
    WithTriageResponse(ingesttest.FixtureTriageResponse())
client := ledger.NewClient(provider)
```

## Query SDK

Read layers, findings, and verification results from the ledger service:

```go
import "github.com/traust-security/traust-sdk/go/v1/ledger"

client := ledger.NewHTTPClient("https://ledger.example.com",
    ledger.WithBearerToken(os.Getenv("LEDGER_TOKEN")),
)

// Read resolved findings for a layer
findings, err := client.GetFindings(ctx, "repo-a")
// Paginated bulk findings
page, err := client.ListFindings(ctx, ledger.ListFindingsOpts{Limit: 50})
// Verify layer integrity
result, err := client.VerifyLayer(ctx, "repo-a", ledger.VerifyOpts{CheckSignatures: true})
// Resolve the authenticated actor without recording an event
actor, err := client.Whoami(ctx)
```

The built-in HTTP transport handles auth and path construction.
For custom transports, implement `ledger.Provider` and use `ledger.NewClient(provider)`.

### Testing

```go
import "github.com/traust-security/traust-sdk/go/v1/ledger/querytest"

provider := querytest.NewStaticProvider().
    WithFindingsResponse("repo-a", querytest.FixtureFindingsResponse())
client := ledger.NewClient(provider)
```

## Storage SDK

Storage retains exact artifact evidence and its canonical relational projection in
one SQLite or PostgreSQL transaction. PostgreSQL relations use the fixed
`traust_storage` schema so application tables and `search_path` cannot redirect
storage operations. SQLite uses the caller-selected database file as its physical
namespace; a dedicated file is recommended. The caller owns the driver, DSN, pool
settings, health checks, and database closure:

```go
import (
    "database/sql"

    _ "github.com/jackc/pgx/v5/stdlib"
    "github.com/traust-security/traust-sdk/go/v1/storage"
)

db, err := sql.Open("pgx", dsn)
if err != nil {
    return err
}
defer db.Close()

client, err := storage.NewClient(ctx, db)
if err != nil {
    return err
}
if err = client.Init(ctx); err != nil {
    return err
}
subjectID := "sci:inventory-item:" + inventoryItemID
runID := "sci:scan-result:" + scanResultID
result, err := client.SaveVulnFindings(ctx, storage.SaveVulnFindingsInput{
    Binding: storage.Binding{
        ScopeID:   "local",
        SubjectID: &subjectID,
        RunID:     &runID,
    },
    Artifact: artifact,
})
```

Every schema has named typed save and read operations. A named save validates the
retained source bytes, computes their SHA-256 digest, stores globally deduplicated
`artifact_evidence`, and creates a context-specific `artifact_binding` atomically.
`SaveResult.AlreadyBound` reports only whether that binding existed; evidence-level
deduplication remains private.

The exact-evidence boundary is the named Save call. Callers may perform optional
processing, such as asking Ledger to stamp finding fingerprints, before creating
the artifact passed to storage. Storage preserves those accepted bytes but does
not compute fingerprints or treat them as storage-owned identity. When Ledger is
present, consumers join disposition data by the binding's `layer_id` and the
artifact-relative finding ID; fingerprints remain report evidence or
Ledger-owned state rather than storage join keys.

Every artifact contract has a schema-specific SQL projection. Root fields become
typed columns while nested values remain JSON until a query justifies child tables.
Scoped views exercise cross-artifact binding relationships; all projections retain
generated save and smoke-test coverage. Run-bound saves require caller-supplied
subject and run IDs; layer saves require a Ledger layer ID. Identifiers are opaque
UTF-8 strings, and omitted scope defaults to `local`.

Typed reads use `BindingID` to guard the stored schema interpretation. Use
`GetEvidence(ctx, digest)` only when raw bytes without a type claim are intended.
Scoped view reads require an explicit scope list.

## Data-only usage

For consumers who just need types or validation without invoking skills:

| Package | Contents |
|---|---|
| [`v1/types`](v1/types) | Generated Go structs for every schema in `schemas/v1` |
| [`v1/enums`](v1/enums) | Generated shared vocabularies (severity, validity, verdicts, ...) |
| [`v1/validate`](v1/validate) | Validate raw JSON bytes against an embedded v1 schema |

## Regenerating types/enums/validate/storage

```bash
make generate     # verified pinned contracts commit → generated Go outputs
make check-drift  # fails if generated output is stale
```

No sibling checkout is needed. `make generate` obtains the full immutable
`CONTRACTS_REF` from `CONTRACTS_REPO` and validates that the checkout matches it
exactly; it never generates from latest `main`. The checkout is reused from
`CONTRACTS_CACHE` (by default the user cache) at `CONTRACTS_SOURCE`. All four
variables are caller-overridable for fork, CI-cache, and pinned-ref workflows.
Normal `make test` uses only repository fixtures and requires no network.

Run `make test-integration` for the storage suite against SQLite and PostgreSQL.
The PostgreSQL suite uses one fixed local test DSN, recreates `traust_storage`
for each test, and verifies that same-named application tables remain untouched.
Start the database with:

```bash
podman run --name traust-postgres --rm -d -e POSTGRES_USER=traust -e POSTGRES_PASSWORD=traust-test-only -e POSTGRES_DB=traust_test -p 127.0.0.1:5432:5432 -v traust-postgres-data:/var/lib/postgresql/data docker.io/library/postgres:16
```

These are fixed local test-only credentials, **not production/deployed secrets**,
and no DSN override is supported. `make test-integration` disables the Go test cache so PostgreSQL availability is
checked on every run. PostgreSQL absence, connectivity failure, or authentication
failure produces a visible skip while SQLite still executes. Once PostgreSQL
accepts a connection, all version, initialization, and protocol failures fail.
