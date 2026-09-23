# Changelog

All notable changes to the Go SDK are documented here.

## [0.15.0]

### Changed

- Generate storage format, revision and baseline identity from pinned contracts
  metadata. The new `traust-storage-20260922` baseline starts at revision 1.
- Reject legacy, mismatched, empty or missing compatibility metadata on existing
  stores; initialization never rewrites history or relabels an old schema.
- Add bidirectional Python/Go initialization tests for SQLite and PostgreSQL.
  A clean bootstrap/re-import or separately reviewed migration is required;
  older clients must be fenced out before reset. No live database reset occurs.

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
