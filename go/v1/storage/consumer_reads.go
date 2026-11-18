package storage

// consumer_reads.go — the reads the interaction-boundary audit
// (traust-foundry#97, class 3) found consumers serving by reaching into
// traust_storage directly because the client that owns the schema exposed
// no method for the shape. Each moves one existing raw SQL read behind
// this package — no DDL, no shape changes — until the read-side plan
// retires the package onto core's REST surface.

import (
	"context"
	"database/sql"
	"errors"
)

// FindBindingBySubject returns the binding an artifact name holds for one
// subject and run — the lookup a report's companion artifacts (a threat
// model recorded beside its report) resolve through. When more than one
// binding matches, the newest wins; no match answers ErrBindingNotFound.
func (c *Client) FindBindingBySubject(ctx context.Context, subjectID, runID, artifactName string) (BindingRecord, error) {
	if c == nil || c.store == nil {
		return BindingRecord{}, wrap(OperationRead, PhaseInput, ErrNilDatabase)
	}
	for _, value := range []string{subjectID, runID, artifactName} {
		if value == "" || !validBindingText(value) {
			return BindingRecord{}, wrap(OperationRead, PhaseInput, ErrInvalidIdentifier)
		}
	}
	conn, err := c.store.db.Conn(ctx)
	if err != nil {
		return BindingRecord{}, wrap(OperationRead, PhaseConnect, err)
	}
	defer func() { _ = conn.Close() }()
	var bindingID string
	err = c.store.queries.artifactBindingFindBySubject(ctx, conn, artifactBindingFindBySubjectParams{
		subjectId: subjectID, runId: runID, artifactName: artifactName,
	}).Scan(&bindingID)
	if errors.Is(err, sql.ErrNoRows) {
		return BindingRecord{}, wrap(OperationRead, PhaseRead, ErrBindingNotFound)
	}
	if err != nil {
		return BindingRecord{}, wrap(OperationRead, PhaseRead, err)
	}
	return c.store.readBinding(ctx, conn, bindingID)
}

// CountProjectionRows returns how many per-finding projection rows a
// published artifact landed for its binding — the completeness check a
// consumer runs before trusting a re-read of an artifact it published.
// The artifact name must carry a per-finding projection (report, triage);
// any other answers ErrUnknownArtifact.
func (c *Client) CountProjectionRows(ctx context.Context, bindingID, artifactName string) (int64, error) {
	if c == nil || c.store == nil {
		return 0, wrap(OperationRead, PhaseInput, ErrNilDatabase)
	}
	if !digestPattern.MatchString(bindingID) {
		return 0, wrap(OperationRead, PhaseInput, ErrBindingNotFound)
	}
	var count func(context.Context, *sql.Conn, string) (*sql.Row, error)
	switch artifactName {
	case "report":
		count = func(ctx context.Context, conn *sql.Conn, bindingID string) (*sql.Row, error) {
			return c.store.queries.reportFindingCount(ctx, conn, reportFindingCountParams{bindingId: bindingID}), nil
		}
	case "triage":
		count = func(ctx context.Context, conn *sql.Conn, bindingID string) (*sql.Row, error) {
			return c.store.queries.triageVerdictCount(ctx, conn, triageVerdictCountParams{bindingId: bindingID}), nil
		}
	default:
		return 0, wrap(OperationRead, PhaseInput, ErrUnknownArtifact)
	}
	conn, err := c.store.db.Conn(ctx)
	if err != nil {
		return 0, wrap(OperationRead, PhaseConnect, err)
	}
	defer func() { _ = conn.Close() }()
	row, err := count(ctx, conn, bindingID)
	if err != nil {
		return 0, wrap(OperationRead, PhaseRead, err)
	}
	var rows int64
	if err := row.Scan(&rows); err != nil {
		return 0, wrap(OperationRead, PhaseRead, err)
	}
	return rows, nil
}

// StorageBaseline is the revision a storage schema records.
type StorageBaseline struct {
	ContractVersion string
	Revision        int
}

// StorageRevision reads the baseline the bound database records, failing
// closed when the recorded baseline is not what this build stamps and
// accepts (ErrIncompatibleRevision) or when the recorded evidence shape
// disagrees with the revision it claims (ErrIncompatibleDatabase) — the
// startup check a consumer runs beside Init.
func (c *Client) StorageRevision(ctx context.Context) (StorageBaseline, error) {
	if c == nil || c.store == nil {
		return StorageBaseline{}, wrap(OperationRead, PhaseInput, ErrNilDatabase)
	}
	conn, err := c.store.db.Conn(ctx)
	if err != nil {
		return StorageBaseline{}, wrap(OperationRead, PhaseConnect, err)
	}
	defer func() { _ = conn.Close() }()
	var version string
	var revision int
	if err := c.store.queries.traustStorageMetaGet(ctx, conn, traustStorageMetaGetParams{}).Scan(&version, &revision); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return StorageBaseline{}, wrap(OperationRead, PhaseRead, ErrNotFound)
		}
		return StorageBaseline{}, wrap(OperationRead, PhaseRead, err)
	}
	if version != storageFormatVersion || revision != contractRevision {
		return StorageBaseline{}, wrap(OperationRead, PhaseRead, ErrIncompatibleRevision)
	}
	var hasSize, hasPayload bool
	if err := c.store.queries.artifactEvidenceShape(ctx, conn, artifactEvidenceShapeParams{}).Scan(&hasSize, &hasPayload); err != nil {
		return StorageBaseline{}, wrap(OperationRead, PhaseRead, err)
	}
	if !hasSize || hasPayload {
		return StorageBaseline{}, wrap(OperationRead, PhaseRead, ErrIncompatibleDatabase)
	}
	return StorageBaseline{ContractVersion: version, Revision: revision}, nil
}

// The queries below follow the frozen queries.go shape (one const pair per
// dialect, a params struct, a method on queries); the storage generator
// cannot emit them since contracts dropped the queries/ layout, so the
// package's hand-maintained reads grow here beside their callers.

const artifactBindingFindBySubjectPostgres = "SELECT binding_id\nFROM traust_storage.artifact_binding\nWHERE subject_id = $1\n  AND run_id = $2\n  AND artifact_name = $3\nORDER BY bound_at DESC, binding_id DESC\nLIMIT 1;"
const artifactBindingFindBySubjectSQLite = "SELECT binding_id\nFROM artifact_binding\nWHERE subject_id = ?\n  AND run_id = ?\n  AND artifact_name = ?\nORDER BY bound_at DESC, binding_id DESC\nLIMIT 1;"

type artifactBindingFindBySubjectParams struct {
	subjectId    string
	runId        string
	artifactName string
}

func (q queries) artifactBindingFindBySubject(ctx context.Context, conn *sql.Conn, p artifactBindingFindBySubjectParams) *sql.Row {
	statement := artifactBindingFindBySubjectSQLite
	if q.dialect == dialectPostgres {
		statement = artifactBindingFindBySubjectPostgres
	}
	return conn.QueryRowContext(ctx, statement, p.subjectId, p.runId, p.artifactName)
}

const reportFindingCountPostgres = "SELECT count(*)\nFROM traust_storage.report_finding\nWHERE binding_id = $1;"
const reportFindingCountSQLite = "SELECT count(*)\nFROM report_finding\nWHERE binding_id = ?;"

type reportFindingCountParams struct {
	bindingId string
}

func (q queries) reportFindingCount(ctx context.Context, conn *sql.Conn, p reportFindingCountParams) *sql.Row {
	statement := reportFindingCountSQLite
	if q.dialect == dialectPostgres {
		statement = reportFindingCountPostgres
	}
	return conn.QueryRowContext(ctx, statement, p.bindingId)
}

const triageVerdictCountPostgres = "SELECT count(*)\nFROM traust_storage.triage_verdict\nWHERE binding_id = $1;"
const triageVerdictCountSQLite = "SELECT count(*)\nFROM triage_verdict\nWHERE binding_id = ?;"

type triageVerdictCountParams struct {
	bindingId string
}

func (q queries) triageVerdictCount(ctx context.Context, conn *sql.Conn, p triageVerdictCountParams) *sql.Row {
	statement := triageVerdictCountSQLite
	if q.dialect == dialectPostgres {
		statement = triageVerdictCountPostgres
	}
	return conn.QueryRowContext(ctx, statement, p.bindingId)
}

// The evidence shape the recorded revision implies: byte_size present
// (sizes are stored beside digests), payload absent (bytes live in the
// caller's evidence store, never in the database).
const artifactEvidenceShapePostgres = "SELECT EXISTS (\n    SELECT 1\n    FROM information_schema.columns\n    WHERE table_schema = 'traust_storage'\n      AND table_name = 'artifact_evidence'\n      AND column_name = 'byte_size'\n),\nEXISTS (\n    SELECT 1\n    FROM information_schema.columns\n    WHERE table_schema = 'traust_storage'\n      AND table_name = 'artifact_evidence'\n      AND column_name = 'payload'\n);"
const artifactEvidenceShapeSQLite = "SELECT EXISTS (\n    SELECT 1\n    FROM pragma_table_info('artifact_evidence')\n    WHERE name = 'byte_size'\n),\nEXISTS (\n    SELECT 1\n    FROM pragma_table_info('artifact_evidence')\n    WHERE name = 'payload'\n);"

type artifactEvidenceShapeParams struct {
}

func (q queries) artifactEvidenceShape(ctx context.Context, conn *sql.Conn, p artifactEvidenceShapeParams) *sql.Row {
	statement := artifactEvidenceShapeSQLite
	if q.dialect == dialectPostgres {
		statement = artifactEvidenceShapePostgres
	}
	return conn.QueryRowContext(ctx, statement)
}
