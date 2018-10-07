package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type postgresFactory struct {
	option *Option
}

func (factory *postgresFactory) GetReadonlyDB() func() (ReadonlyDB, error) {
	var db, err = sql.Open("postgres", factory.option.DefaultURI)
	if err == nil {
		db.SetMaxIdleConns(factory.option.MaxIdleConns)
		db.SetMaxOpenConns(factory.option.MaxOpenConns)
	}

	return func() (ReadonlyDB, error) {
		return db, err
	}
}

func (factory *postgresFactory) GetDefaultDB() func() (DefaultDB, error) {
	var db, err = sql.Open("postgres", factory.option.DefaultURI)
	if err == nil {
		db.SetMaxIdleConns(factory.option.MaxIdleConns)
		db.SetMaxOpenConns(factory.option.MaxOpenConns)
	}

	return func() (DefaultDB, error) {
		return db, err
	}
}

func (factory *postgresFactory) Close() []error {
	var errors = make([]error, 0, 4)
	defaultDB, err := factory.GetDefaultDB()()
	if err != nil {
		errors = append(errors, err)
	} else {
		var db, ok = defaultDB.(*sql.DB)
		if ok {
			if err = db.Close(); err != nil {
				errors = append(errors, err)
			}
		}
	}

	readonlyDB, err := factory.GetReadonlyDB()()
	if err != nil {
		errors = append(errors, err)
	} else {
		var db, ok = readonlyDB.(*sql.DB)
		if ok {
			if err = db.Close(); err != nil {
				errors = append(errors, err)
			}
		}
	}

	return errors
}

func newPostgresFactory(option *Option) Factory {
	return &postgresFactory{
		option: option,
	}
}

func init() {
	New = newPostgresFactory
}
