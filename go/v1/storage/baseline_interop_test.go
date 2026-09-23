package storage

import (
	"context"
	"database/sql"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestPythonGoBaselineInteroperability(t *testing.T) {
	python := os.Getenv("TRAUST_CONTRACTS_PYTHON")
	if python == "" {
		t.Skip("set TRAUST_CONTRACTS_PYTHON to a Python environment with the pinned contracts")
	}
	for _, first := range []string{"python", "go"} {
		t.Run(first+"-initializes", func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "interop.db")
			pythonInit := func() {
				t.Helper()
				command := exec.CommandContext(context.Background(), python, "-c",
					"import sqlite3,sys; from traust_contracts.v1.storage import Store; connection=sqlite3.connect(sys.argv[1]); Store(connection).init(); connection.close()", path)
				if output, err := command.CombinedOutput(); err != nil {
					t.Fatalf("Python initialization: %v: %s", err, output)
				}
			}
			if first == "python" {
				pythonInit()
			}
			db, err := sql.Open("sqlite", path)
			if err != nil {
				t.Fatal(err)
			}
			defer func() { _ = db.Close() }()
			db.SetMaxOpenConns(1)
			client, err := NewClient(context.Background(), db)
			if err != nil {
				t.Fatal(err)
			}
			if err := client.Init(context.Background()); err != nil {
				t.Fatal(err)
			}
			checkReportProjection(t, client, db, "")
			pythonInit()
			if err := client.Init(context.Background()); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestPostgresPythonGoBaselineInteroperability(t *testing.T) {
	python := os.Getenv("TRAUST_CONTRACTS_PYTHON")
	if python == "" {
		t.Skip("set TRAUST_CONTRACTS_PYTHON to a Python environment with the pinned contracts")
	}
	for _, first := range []string{"python", "go"} {
		t.Run(first+"-initializes", func(t *testing.T) {
			client, db := openPostgresStorage(t)
			pythonInit := func() {
				t.Helper()
				command := exec.CommandContext(context.Background(), python, "-c",
					"import psycopg,sys; from traust_contracts.v1.storage import Store; connection=psycopg.connect(sys.argv[1]); Store(connection).init(); connection.close()", postgresDSN)
				if output, err := command.CombinedOutput(); err != nil {
					t.Fatalf("Python PostgreSQL initialization: %v: %s", err, output)
				}
			}
			if first == "python" {
				if _, err := db.Exec("DROP SCHEMA traust_storage CASCADE"); err != nil {
					t.Fatal(err)
				}
				pythonInit()
				if err := client.Init(context.Background()); err != nil {
					t.Fatal(err)
				}
			}
			checkReportProjection(t, client, db, "traust_storage.")
			pythonInit()
			if err := client.Init(context.Background()); err != nil {
				t.Fatal(err)
			}
		})
	}
}
