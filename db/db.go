//go:generate go-bindata -o ../data/bindata.go -pkg data ../data/schema/
package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/yuuki/lstf/tcpflow"
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

const (
	POST_TIMEOUT_SEC = 10
)

func (db *DB) PostHostFlows(flows tcpflow.HostFlows) error {
	ctx, cancel := context.WithTimeout(context.Background(), POST_TIMEOUT_SEC*time.Second)
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		cancel()
		return errors.Wrap(err, "begin transaction error")
	}
	q1 := `
	INSERT INTO nodes (ipv4, port) VALUES ($1, $2)
	ON CONFLICT (ipv4, port) DO NOTHING
	RETURNING *
`
	stmt1, err := tx.PrepareContext(ctx, q1)
	if err != nil {
		cancel()
		return errors.Wrapf(err, "query prepare error: %s", q1)
	}
	q2 := `
	INSERT INTO flows
	(direction, source_node_id, destination_node_id, connections, updated)
	VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP)
	ON CONFLICT ON CONSTRAINT source_dest_direction_idx DO UPDATE
`
	stmt2, err := tx.PrepareContext(ctx, q2)
	if err != nil {
		cancel()
		return errors.Wrapf(err, "query prepare error: %s", q2)
	}

	for _, flow := range flows {
		var localNodeid, peerNodeid int64
		err := stmt1.QueryRowContext(ctx, q1, flow.Peer.Addr, flow.Peer.Port).Scan(&peerNodeid)
		if err != nil {
			cancel()
			return errors.Wrapf(err, "query error")
		}
		_, err = stmt2.ExecContext(ctx, q2, flow.Direction, localNodeid, peerNodeid, flow.Connections)
		if err != nil {
			cancel()
			return errors.Wrapf(err, "query error")
		}
	}
	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "transaction commit error")
	}
	return nil
}
