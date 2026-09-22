package storage

import (
	"context"
	"database/sql"
	"strings"
)

func (s *sqlStore) requireNoStorageObjects(ctx context.Context, conn *sql.Conn, temporaryOnly bool) error {
	if s.dialect == dialectPostgres {
		if temporaryOnly {
			return nil
		}
		var occupied bool
		err := conn.QueryRowContext(ctx, "SELECT EXISTS (SELECT 1 FROM pg_catalog.pg_class c JOIN pg_catalog.pg_namespace n ON n.oid=c.relnamespace WHERE n.nspname='traust_storage')").Scan(&occupied)
		if err != nil {
			return wrap(OperationInit, PhaseRevision, err)
		}
		if occupied {
			return wrap(OperationInit, PhaseRevision, ErrIncompatibleRevision)
		}
		return nil
	}
	statement := "SELECT name, tbl_name FROM temp.sqlite_schema"
	if !temporaryOnly {
		statement = "SELECT name, tbl_name FROM main.sqlite_schema UNION ALL SELECT name, tbl_name FROM temp.sqlite_schema"
	}
	rows, err := conn.QueryContext(ctx, statement)
	if err != nil {
		return wrap(OperationInit, PhaseRevision, err)
	}
	defer func() { _ = rows.Close() }()
	for rows.Next() {
		var name, table string
		if err := rows.Scan(&name, &table); err != nil {
			return wrap(OperationInit, PhaseRevision, err)
		}
		if storageObjectNames[strings.ToLower(name)] || storageObjectNames[strings.ToLower(table)] {
			return wrap(OperationInit, PhaseRevision, ErrIncompatibleRevision)
		}
	}
	if err := rows.Err(); err != nil {
		return wrap(OperationInit, PhaseRevision, err)
	}
	return nil
}
