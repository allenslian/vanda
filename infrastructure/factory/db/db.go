package db

import (
	"context"
	"database/sql"
)

type (
	//Option describes some database settings, such as connection string.
	Option struct {
		DefaultURI   string
		ReadonlyURI  string
		MaxOpenConns int
		MaxIdleConns int
	}

	//Factory produces database instance.
	Factory interface {
		GetReadonlyDB() func() (ReadonlyDB, error)
		GetDefaultDB() func() (DefaultDB, error)
		Close() []error
	}

	//ReadonlyDB indicates it only may read from database.
	ReadonlyDB interface {
		Prepare(query string) (*sql.Stmt, error)
		PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)

		Query(query string, args ...interface{}) (*sql.Rows, error)
		QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)

		QueryRow(query string, args ...interface{}) *sql.Row
		QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	}

	//DefaultDB indicates it may read from or write to database.
	DefaultDB interface {
		Begin() (*sql.Tx, error)
		BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)

		Exec(query string, args ...interface{}) (sql.Result, error)
		ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)

		Prepare(query string) (*sql.Stmt, error)
		PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)

		Query(query string, args ...interface{}) (*sql.Rows, error)
		QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)

		QueryRow(query string, args ...interface{}) *sql.Row
		QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	}
)

var (
	//New creates one database factory.
	New func(option *Option) Factory
)

var (
	factory Factory
)

//InitializeDB sets some database settings.
func InitializeDB(option *Option) {
	factory = New(option)
}

//GetDefaultDB will return one writable database instance.
func GetDefaultDB() (DefaultDB, error) {
	if factory == nil {
		return nil, ErrMissingDBOption
	}
	return factory.GetDefaultDB()()
}

//GetReadonlyDB will return one readonly database instance.
func GetReadonlyDB() (ReadonlyDB, error) {
	if factory == nil {
		return nil, ErrMissingDBOption
	}
	return factory.GetReadonlyDB()()
}

//Close will close all the instances.
func Close() []error {
	if factory == nil {
		return []error{ErrMissingDBOption}
	}
	return factory.Close()
}
