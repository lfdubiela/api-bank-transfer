package database

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lfdubiela/api-bank-transfer/pkg/infra/env"
	"log"
	"strings"
	"time"
)

const (
	defaultConnMaxLifetime = time.Minute * 60
	MysqlAdapterName       = "mysql"
	HealthCheckSql         = "SHOW VARIABLES LIKE '%innodb_read_only%'"
)

type (
	db struct {
		db               *sql.DB
		poolSize         int
		driver           string
		connectionString string
	}

	client struct {
		ReadDB  DB
		WriteDB DB
	}

	DBTx struct {
		*sql.Tx
	}

	DBVariable struct {
		VariableName string
		Value        string
	}

	DBCheck struct {
		CanRead  bool
		CanWrite bool
	}
)

func NewMySqlAdapter(readConn, writeConn string, readPoolSize, writePoolSize int) Client {
	return &client{
		ReadDB:  newMySqlDB(readConn, readPoolSize).Start(),
		WriteDB: newMySqlDB(writeConn, writePoolSize).Start(),
	}
}

func newMySqlDB(connectionString string, poolSize int) DB {
	return &db{poolSize: poolSize, driver: MysqlAdapterName, connectionString: connectionString}
}

func (d client) Reader() DB {
	return d.ReadDB
}

func (d client) Writer() DB {
	return d.WriteDB
}

func Config() Client {
	db := NewMySqlAdapter(
		fmt.Sprintf(env.Get().DbConnStr,
			env.Get().DbReadUser,
			env.Get().DbReadPass,
			env.Get().DbReadHost,
			env.Get().DbReadName),
		fmt.Sprintf(env.Get().DbConnStr,
			env.Get().DbWriteUser,
			env.Get().DbWritePass,
			env.Get().DbWriteHost,
			env.Get().DbWriteName),
		env.Get().DbReadPool,
		env.Get().DbReadPool,
	)

	return db
}

func (d client) CheckDB(ctx context.Context) DBCheck {
	read := d.ReadDB.CheckRead(ctx)
	write := d.WriteDB.CheckWrite(ctx)

	return DBCheck{
		CanRead:  read,
		CanWrite: write,
	}
}

func (m db) CheckRead(ctx context.Context) bool {
	var dbVariable DBVariable

	row, err := m.QueryRow(ctx, nil, HealthCheckSql)
	if err != nil {
		return false
	}

	err = row.Scan(&dbVariable.VariableName, &dbVariable.Value)
	if err != nil {
		return false
	}

	return true
}

func (m db) CheckWrite(ctx context.Context) bool {
	var dbVariable DBVariable

	row, err := m.QueryRow(ctx, nil, HealthCheckSql)
	if err != nil {
		return false
	}

	err = row.Scan(&dbVariable.VariableName, &dbVariable.Value)
	if err != nil {
		return false
	}

	if strings.ToUpper(dbVariable.Value) != "OFF" {
		return false
	}

	return true
}

func (m *db) Start() DB {
	if m.db != nil {
		if err := m.db.Ping(); err == nil {
			return nil
		}
		m.CloseDB()
	}

	db, err := sql.Open(
		m.driver,
		m.connectionString)

	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(m.poolSize)
	db.SetMaxIdleConns(m.poolSize / 2)
	db.SetConnMaxLifetime(defaultConnMaxLifetime)

	m.db = db

	return m
}

func (m *db) CloseDB() {
	err := m.db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func (m *db) Begin() (Transaction, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return nil, ErrBeginTransaction.WithCause(err)
	}

	return tx, nil
}

func (m db) Exec(ctx context.Context, tx Transaction, query string, returnRows bool, args ...interface{}) (int, error) {
	result, err := m.exec(ctx, tx, query, args...)
	if err != nil {
		return 0, err
	}

	numRows, err := result.RowsAffected()
	switch {
	case err != nil:
		return 0, ErrRowsAffected.WithCause(err)
	case returnRows:
		return int(numRows), nil
	case numRows <= 0:
		err := WarnNoRowsAffected
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, ErrLastInsertID.WithCause(err)
	}

	return int(id), err
}

func (m db) QueryRow(ctx context.Context, tx Transaction, query string, args ...interface{}) (*sql.Row, error) {
	row := m.queryRow(ctx, tx, query, args...)

	if row != nil {
		err := row.Err()
		if err != nil {
			err = ErrQueryRow.WithCause(err)
			return nil, err
		}

		return row, nil
	}

	return nil, ErrQuery
}

func (m db) Query(ctx context.Context, tx Transaction, query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := m.queryRows(ctx, tx, query, args...)
	if err != nil {
		return nil, ErrQuery.WithCause(err)
	}

	return rows, err
}

func (m db) exec(ctx context.Context, tx Transaction, query string, args ...interface{}) (sql.Result, error) {
	if tx == nil {
		return m.db.ExecContext(ctx, query, args...)
	}

	return tx.ExecContext(ctx, query, args...)
}

func (m db) queryRow(ctx context.Context, tx Transaction, query string, args ...interface{}) *sql.Row {
	if tx == nil {
		return m.db.QueryRowContext(ctx, query, args...)
	}

	return tx.QueryRowContext(ctx, query, args...)
}

func (m db) queryRows(ctx context.Context, tx Transaction, query string, args ...interface{}) (*sql.Rows, error) {
	if tx == nil {
		return m.db.QueryContext(ctx, query, args...)
	}

	return tx.QueryContext(ctx, query, args...)
}
