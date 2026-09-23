package storage

import (
	"context"
	"database/sql"
	"errors"
	"path/filepath"
	"testing"
)

func TestStorageBaselineRejectsIncompatibleMetadata(t *testing.T) {
	for _, mutation := range []string{
		"UPDATE traust_storage_meta SET baseline_id='old-baseline'",
		"UPDATE traust_storage_meta SET revision=15",
		"UPDATE traust_storage_meta SET revision=16",
		"UPDATE traust_storage_meta SET contract_version='v0'",
		"DELETE FROM traust_storage_meta",
		"DROP TABLE traust_storage_meta",
		"ALTER TABLE traust_storage_meta DROP COLUMN baseline_id",
		"CREATE TEMP TABLE traust_storage_meta (id INTEGER)",
	} {
		t.Run(mutation, func(t *testing.T) {
			client := openTestStorage(t)
			if _, err := sqlDB(client).Exec(mutation); err != nil {
				t.Fatal(err)
			}
			if err := client.Init(context.Background()); err == nil {
				t.Fatal("incompatible store was accepted")
			}
		})
	}
}

func TestUnstampedStorageRejected(t *testing.T) {
	for _, setup := range []string{
		"CREATE TABLE report (marker TEXT)",
		"CREATE VIEW current_binding AS SELECT 1",
		"CREATE VIEW CURRENT_BINDING AS SELECT 1",
		"CREATE TEMP TABLE REPORT (marker TEXT)",
		"CREATE TEMP TABLE report (marker TEXT)",
		"CREATE TABLE traust_storage_meta (id INTEGER, contract_version TEXT, revision INTEGER, applied_at TEXT)",
	} {
		t.Run(setup, func(t *testing.T) {
			db, err := sql.Open("sqlite", ":memory:")
			if err != nil {
				t.Fatal(err)
			}
			defer func() { _ = db.Close() }()
			db.SetMaxOpenConns(1)
			if _, err := db.Exec(setup); err != nil {
				t.Fatal(err)
			}
			client, err := NewClient(context.Background(), db)
			if err != nil {
				t.Fatal(err)
			}
			if err := client.Init(context.Background()); err == nil {
				t.Fatal("unstamped store was adopted")
			}
		})
	}
}

func TestStorageBaselinePreservesMetadata(t *testing.T) {
	client := openTestStorage(t)
	db := sqlDB(client)
	var version, baseline, applied string
	var revision int
	if err := db.QueryRow("SELECT contract_version, revision, baseline_id, applied_at FROM traust_storage_meta").Scan(&version, &revision, &baseline, &applied); err != nil {
		t.Fatal(err)
	}
	if version != storageFormatVersion || revision != 1 || baseline != storageBaselineID {
		t.Fatalf("unexpected baseline: %s %d %s", version, revision, baseline)
	}
	if err := client.Init(context.Background()); err != nil {
		t.Fatal(err)
	}
	var after string
	if err := db.QueryRow("SELECT applied_at FROM traust_storage_meta").Scan(&after); err != nil || after != applied {
		t.Fatalf("initialization changed timestamp: %v", err)
	}
}

func TestStorageRejectsShadowOnReusedConnection(t *testing.T) {
	ctx := context.Background()
	db, err := sql.Open("sqlite", filepath.Join(t.TempDir(), "pooled.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = db.Close() }()
	db.SetMaxOpenConns(2)
	shadow, err := db.Conn(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = shadow.Close() }()
	if _, err := shadow.ExecContext(ctx, "CREATE TEMP TABLE REPORT (marker TEXT)"); err != nil {
		t.Fatal(err)
	}
	client, err := NewClient(ctx, db)
	if err != nil {
		t.Fatal(err)
	}
	if err := client.Init(ctx); err != nil {
		t.Fatal(err)
	}
	clean, err := db.Conn(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = clean.Close() }()
	if err := shadow.Close(); err != nil {
		t.Fatal(err)
	}
	if _, err := client.GetEvidence(ctx, "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"); !errors.Is(err, ErrIncompatibleRevision) {
		t.Fatalf("shadowing pooled connection not rejected by baseline guard: %v", err)
	}
}

func TestPostgresBaselineRejectsIncompatibleMetadata(t *testing.T) {
	for _, mutation := range []string{
		"UPDATE traust_storage.traust_storage_meta SET baseline_id='old-baseline'",
		"DELETE FROM traust_storage.traust_storage_meta",
		"DROP TABLE traust_storage.traust_storage_meta",
		"ALTER TABLE traust_storage.traust_storage_meta DROP COLUMN baseline_id",
	} {
		t.Run(mutation, func(t *testing.T) {
			client, db := openPostgresStorage(t)
			if _, err := db.Exec(mutation); err != nil {
				t.Fatal(err)
			}
			if err := client.Init(context.Background()); err == nil {
				t.Fatal("incompatible PostgreSQL store was accepted")
			}
		})
	}
}
