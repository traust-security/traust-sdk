# Changelog

All notable changes to the Go SDK are documented here.

## [0.23.0]

### Added

- **`QueryCurrentFindingDetail`: every current finding at listing grain.**
  The open/hardening cuts answer "how much exposure"; a cross-repo findings
  worklist needs the rows themselves. Returns the `current_finding_detail`
  view (new in contracts 0.61.0): the spine plus `cvss_score`, the triage
  verdict and triage-assessed severity of the newest triage for the same
  scope, subject and run, and the census `exposure_class` per row. Unlike
  `QueryOpenFindings`/`QueryHardeningFindings` it is deliberately unfiltered
  on disposition -- callers choose their cut through `ExposureClass`.
- The spine now carries `cvss_score` (extracted from the report finding's
  `cvss` object; NULL for the policy family), so views above it no longer
  re-join the finding tables to reach it.

### Fixed

- **`GetProductScopes` query bindings are generated, not hand-edited.** The
  `product_scope` read now has a canonical contracts source
  (`product_scope.get.sql`, both dialects), and the generator treats it as a
  multi-row keyed read per the `.get.sql` convention. Regenerating from
  contracts no longer drops `GetProductScopes`, so `make check-drift` passes
  again.

### Changed

- Generated storage assets are pinned to contracts commit `3ed6815`
  (feat/current-finding-detail), absorbing contracts 0.57–0.61: review
  surfaces on reports/verifications/threat models, model pricing tiers,
  portfolio forge facts and the current_finding_detail view.

## [0.22.0]

### Added

- **Registry enumeration.** `QueryProducts` and `QueryProductRepos` enumerate
  the product/repo registry through new `.search.sql` contract reads (filtered
  by product slug or repo-URL LIKE pattern, with limit/offset). Until now the
  registry was register-plus-point-lookup only, which forced consumers that
  need a working set (e.g. Source Code Intelligence) to mirror registry rows
  in their own tables.
- `RegistryFilter` validates limits; `ErrRegistryLimitInvalid` reports a
  non-positive limit or negative offset.

### Changed

- Generated storage assets are pinned to contracts commit `a150b2b5` (the
  storage-revision-1 registry-enumeration backport). Storage revision stays
  1: existing databases boot on this SDK without migration.

## [0.21.0]

### Added

- **Typed, pluggable layer-listing filters.** `ListLayers` accepts any
  combination of `LayerListFilter` values: `ReviewStatus`/`ReviewStatusIn`,
  `ReviewQueueReason`, `LayerID`/`LayerIDs`, `ProductRepoID`,
  `MerkleEpochAtLeast`/`MerkleEpochAtMost`, `UpdatedSince`/`UpdatedUntil`,
  `WithReviews` (embed `needs_review` queues), and `RawFilter` — the escape
  hatch for registry keys this SDK version does not model. Each filter
  renders into the ledger's allowlisted `filter=key:op:value` registry;
  existing `ListLayers(ctx)` call sites compile unchanged. Invalid filters
  (empty values, list separators inside `in` members, zero timestamps) are
  rejected before reaching the wire. **Requires traust-ledger#86** (commit
  `ebd5f4b`; the ledger is pinned by commit downstream) for the filter
  endpoint; unfiltered listings and `FindLayer` are unchanged.
- `LayerRef.NeedsReview []types.ReviewItem`: matching queue entries are
  embedded in listing responses when review filters or `WithReviews`
  request them.

## [0.20.0]

### Changed (breaking)

- **Layer calls are query-addressed.** `GetLayer`, `GetFindings`, `ListEvents`,
  `VerifyLayer`, `BatchSubmit`, `ResolveReviewItem`, `StampEventIdentities` and
  `SignLayer` now call `/v1/ledger/layer<op>?layer_id=…` instead of splicing the
  ID into `/v1/ledger/layers/{layer_id}…`. The ID is query-encoded, so opaque
  database IDs with `:` and `/` (migrated corpus layers such as
  `corpus:layer:<subject>`) reach the right layer. **Requires traust-ledger
  0.10.0** (traust-ledger#71); older services do not serve these routes.
  Countersign/severity (`POST /v1/ledger/events`) and `InitializeLayer` are
  unchanged.

### Added

- `ledger.LayerPath(op, layerID, extra)` and `LayerOp*` constants: the single
  route builder, also used by the `querytest`/`ingesttest` doubles so canned
  responses match the exact path the client sends.

## [0.19.0]

Unreleased registry reader review candidate; no release or ledger activation.

### Added

- Registry-driven `enums.Normalize` reader candidate with private generated
  replacement metadata, ordered value access, drop flags and retired keys
  independent of constants. Known-enum calls allocate nothing and expose no
  mutable registry slices. Neutral shared cases exercise generated behavior.
  Placement/affected-contract review remains pending; the published contracts
  pin, existing symbols, ledger activation and historical payloads are unchanged.

## [0.18.1]

### Fixed

- **Storage revision 2.** Generated against traust-contracts 0.48.1, and
  `Init` now stamps and requires storage revision 2. 0.18.0 changed the
  schema (`artifact_role`, `artifact_location`) but still accepted a
  revision-1 database created by 0.17.x, which then failed its first read or
  write with `no such column: artifact_role`. Such a database is now refused
  on open with `ErrIncompatibleRevision`; recreate it.

## [0.18.0]

### Changed (breaking)

- **Storage never writes artifact bytes.** `storage.ObjectStore{Put,Get}` is
  replaced by `storage.Resolver{Fetch(ctx, reference, meta)}`. The producer (the
  harness) is the only byte writer; a Save registers where the bytes already
  are. `NewClient(ctx, db, opts...)` takes `storage.WithResolver(r)`, optional
  for register-only loaders. Typed reads try references in order and verify
  size and SHA-256; a client with no resolver returns `ErrNoResolver`.
- `GetEvidence(ctx, digest)` is removed in favour of
  `GetPayload(ctx, bindingID)`: references belong to a binding, so a digest
  alone no longer names a location.
- `storagetest.MemoryStore` is replaced by `storagetest.MemoryResolver`
  (`Write` plays the producer, `Fetch` the resolver).
- `ErrNilObjectStore` is replaced by `ErrNoResolver`.

### Added

- Generated against traust-contracts 0.48.0.
- Every `Save*Input` has `References []string`, recorded in `artifact_location`;
  a retry may add references.
- `Binding.Role`, validated against the contracts profile (`report`:
  `baseline`, `cumulative`; `ErrRoleNotAllowed` otherwise). Encoded trailing and
  present-only, so bindings without a role keep their IDs; matches the Python
  golden vector.
- `BindingRecord.References` and `BindingRecord.ByteSize`.
- `ErrInvalidReference` for empty or NUL references, rejected before any write.

## [0.17.2]

### Fixed

- **`ConvertValidationReport` honours the report's soundness and evidence
  grades.** It mapped every `refuted` verdict to a `false_positive` event and
  ignored `soundness_flag` and `evidence_grade`, so a refutation the
  validation engine had marked unsound (the probe errored, enumerated no
  subjects, or the target was never deployed) entered the countersign queue
  as if it were sound. It now matches the Python emitter:
  - a `refuted` or `inconclusive` finding with a `soundness_flag` queues an
    `unsound_refutation` needs_review item and emits no event;
  - an `E3` (inference-only) confirmation queues `weak_confirmation` instead
    of a `confirmed` event;
  - every emitted event carries its `evidence_grade`, which the ledger uses to
    demote E2/E3 execution evidence below class 1.

  Reports with no flags or grades convert exactly as before.

## [0.17.1]

### Changed

- Docs: the object store belongs to the Traust deployment (where its
  `locations.analysis_results` points). The SDK ships no store, and a
  consumer such as SCI only adapts to that store; it never provisions its own.
  Replaces "store implementations live in consumers".

## [0.17.0]

### Changed

- **Generated against traust-contracts 0.47.0.** The threat model's
  `Assets`, `EntryPoints`, `Deprioritized`, `AttackScenarios` and
  `UpdateHistory` are typed structs (`Asset`, `EntryPoint`,
  `DeprioritizedThreat`, `AttackScenario`, `HistoryEntry`) instead of untyped
  objects. Asset `Sensitivity` is typed by the `chain-severity` enum.

## [0.16.0]

### Changed

- **Generated against traust-contracts 0.46.0.**
  - Threats carry the OWASP Risk Rating Methodology `RiskRating`, and the
    legacy `Impact` and `Likelihood` are now optional pointers (contracts
    0.45.0).
  - The `threat` projection writes `risk_rating` and its severity, scores,
    levels and basis. It writes the legacy labels and score only when the
    threat has them (an OWASP-rated threat has neither), matching the Python
    projector.
  - `threat_current` and `threat_exposure` read the new columns (contracts
    0.46.0).

## [0.15.0]

### Changed

- **Generated against traust-contracts 0.44.0** (was a pre-0.37 commit).
  - Every registered enum is now a Go type, and matching fields are typed
    `enums.*` instead of `string`.
  - `refuted-register` gains Save/Get operations.
  - The `layer_metadata` projection is gone (contracts #11); the findings
    summary's repository now comes from subject ownership.
- **storage/v1 no longer retains artifact bytes** (contracts #9). Save records
  each artifact's digest and byte size, its binding, and its projections.
  - The read-back collision check is gone, because the digest is the key.
  - Typed `Get*` and `GetEvidence` return `ErrArtifactBytesNotRetained` until
    a caller-supplied object store backs them.

## Unreleased

### Added

- Artifact saves now require a caller-provided digest-addressed object store
  at client construction; missing stores fail with `ErrNilObjectStore`.
- Typed reads verify external bytes against the database's digest and size.
  An exported in-memory store supports tests without an external bucket.
- External bytes are written before the SQL transaction so failed writes can
  leave orphaned objects rather than committed bindings without bytes.

### Changed

- Storage clients without an object store must be updated before upgrading.
  Existing databases with retained payloads require a coordinated byte backfill
  and schema migration before switching consumers to external storage.

## [0.14.2]

### Fixed

- **Enum codegen no longer emits invalid Go for registry values that are not
  identifier-safe.** A value such as `source+runtime` produced
  `...Source+runtime` and aborted generation. Characters that cannot appear in
  a Go identifier now act as word separators (`...SourceRuntime`), and two
  values in one enum that map to the same constant name are a generation error
  instead of a silent redeclaration. Values that were already identifier-safe
  generate the same constant names as before.
- **Storage projections compile whether or not an enum field is retyped.** The
  type generator retypes a schema enum field to its `enums.*` registry type when
  a registry file matches its values, and leaves it `string` otherwise. Generated
  projectors and the report finding projector now pass enum fields through
  `projectionEnum` (nullable) or `string(...)`, which accept both forms. So a
  contracts release that registers more enums does not break the storage
  package on regeneration.

## [0.14.1]

### Fixed

- Report saves now populate per-finding storage rows in the same transaction as
  the report and exact evidence, making saved findings available to SQL readers.
  Dispositions, optional flags and analytical fields retain their original values.
- Regeneration preserves the secondary report projection. Regression coverage
  checks populated and empty reports, idempotent retries and child-write rollback.
- Previously saved bindings are not backfilled by retrying: `AlreadyBound`
  returns before projection. Existing stores require a separately reviewed
  backfill or a clean re-import; no schema revision reset is included.

## [0.3.0]

## Changes

- **Reverted the 0.13.0 storage changes.** The storage bindings track
  traust-contracts 0.35.0, whose storage contract is the 0.33.0 definition
  (REVISION 15). go/v0.13.0 remains tagged and should not be used.

## [0.2.0]

## Changes

- Storage uses the fixed PostgreSQL `traust_storage` schema and fully qualified
  canonical SQL, preventing application-table collisions and `search_path`
  redirection. SQLite continues to use the caller-selected database file as its
  physical namespace.

- **Two write-path verbs the Python client already had are now on the Go
  client**, reaching the ledger's new REST endpoints (traust-ledger >= 0.3.0):
  - `ledger.Client.StampEventIdentities(ctx, layerID, StampInput{Fingerprints})`
    — POST `/v1/ledger/layers/{layer_id}/stamp`. Backfills event fingerprints
    from a `finding_ref -> fingerprint` map and re-signs the layer; the ledger
    never overwrites an existing fingerprint. Returns the new Merkle root and
    the count stamped.
  - `ledger.Client.Whoami(ctx)` — GET `/v1/ledger/whoami`. Returns the
    token-verified `types.Actor` without recording anything.

## [0.1.1]

## Changes

- `occurred_at` no longer gets a midnight suffix appended to a value that
  already carries a time. Three duplicate `*OccurredAt` helpers collapse into
  one `eventOccurredAt`; a report date with a time component passes through
  unchanged, and only a bare date is padded. Previously a timestamped report
  date produced `"2026-01-16T00:00:00ZT00:00:00+00:00"`, which the ledger's
  dict-only write path stores without complaint and then cannot read back.

- The schema compiler now calls `AssertFormat()`. Without it `format` is
  annotation-only, so `format: date` and `format: date-time` accepted any
  string and report validation passed values the Python side rejects.

### Note

`traust-ledger` >= 0.1.1 enforces RFC 3339 on event timestamps when reading a
layer, but its write path does not validate, so a malformed `occurred_at`
submitted by an older SDK is accepted and only fails on the next read. Upgrade
any Go producer before it writes.

## [0.1.0]

**First public release.**

Typed Go client for the traust-ledger service: convert and submit
triage/validation/verification reports as disposition events, request
layer merkle signing, and query layers/findings/events. Generated
`v1/types`, `v1/enums`, and `v1/validate` track `traust-contracts` schemas.
