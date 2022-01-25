package database

import (
	"context"
	"database/sql"
)

type (
	Transaction interface {
		Commit() error
		Rollback() error
		QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
		QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
		ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	}

	TransactionOpener interface {
		Begin() (Transaction, error)
	}

	Client interface {
		Reader() DB
		Writer() DB
		CheckDB(ctx context.Context) DBChecker
	}

	DBChecker interface {
		CanRead() bool
		CanWrite() bool
	}

	DB interface {
		Start() DB
		CloseDB()
		CheckRead() bool
		CheckWrite() bool
		Begin() (Transaction, error)
		Exec(ctx context.Context, tx Transaction, query string, returnRows bool, args ...interface{}) (int, error)
		Query(ctx context.Context, tx Transaction, query string, args ...interface{}) (*sql.Rows, error)
		QueryRow(ctx context.Context, tx Transaction, query string, args ...interface{}) (*sql.Row, error)
	}
)