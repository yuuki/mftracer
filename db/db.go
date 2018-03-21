//go:generate go-bindata -o ../data/bindata.go -pkg data ../data/schema/
package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/yuuki/mkr-flow-tracer/data"
)

const (
	DefaultDBUserName = "mtracer"
	DefaultDBName     = "mtracer"
)

var (
	Schemas = []string{
		"../data/schema/flows.sql",
	}
)

// DB represents a Database handler.
type DB struct {
	*sql.DB
}

// New creates the DB object.
func New() (*DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf(
		"user=%s dbname=%s sslmode=disable", DefaultDBUserName, DefaultDBName,
	))
	if err != nil {
		return nil, errors.Wrap(err, "postgres open error")
	}
	if err = db.Ping(); err != nil {
		return nil, errors.Wrap(err, "postgres ping error")
	}
	return &DB{db}, nil
}

// CreateSchema creates the table schemas defined by the paths including Schemas.
func (db *DB) CreateSchema() error {
	for _, schema := range Schemas {
		sql, err := data.Asset(schema)
		if err != nil {
			return errors.Wrapf(err, "get schema error: %v", schema)
		}
		_, err = db.Exec(fmt.Sprintf("%s", sql))
		if err != nil {
			return errors.Wrapf(err, "exec schema error: %s", sql)
		}
	}
	return nil
}
