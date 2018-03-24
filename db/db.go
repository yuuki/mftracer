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
	"github.com/yuuki/mftracer/data"
)

const (
	DefaultDBName     = "mftracer"
	DefaultDBUserName = "mftracer"
	DefaultDBHostname = "localhost"
	DefaultDBPort     = "5432"
	ConnectTimeout    = 5
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

// Opt are options for database connection.
// https://godoc.org/github.com/lib/pq
type Opt struct {
	DBName   string
	User     string
	Password string
	Host     string
	Port     string
	SSLMode  string
}

// New creates the DB object.
func New(opt *Opt) (*DB, error) {
	var user, dbname, host, port, sslmode string
	if user = opt.User; user == "" {
		user = DefaultDBUserName
	}
	if dbname = opt.DBName; dbname == "" {
		dbname = DefaultDBName
	}
	if host = opt.Host; host == "" {
		host = DefaultDBHostname
	}
	if port = opt.Port; port == "" {
		port = DefaultDBPort
	}
	if sslmode = opt.SSLMode; sslmode == "" {
		sslmode = "disable"
	}
	db, err := sql.Open("postgres", fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s connect_timeout=%d",
		user, opt.Password, host, port, dbname, sslmode, ConnectTimeout,
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

// InsertOrUpdateHostFlows insert host flows or update it if the same flow exists.
func (db *DB) InsertOrUpdateHostFlows(flows tcpflow.HostFlows) error {
	ctx, cancel := context.WithTimeout(context.Background(), POST_TIMEOUT_SEC*time.Second)
	defer cancel()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "begin transaction error")
	}
	q1 := `
	INSERT INTO nodes (ipv4, port) VALUES ($1, $2)
	ON CONFLICT (ipv4, port) DO NOTHING
	RETURNING node_id
`
	stmt1, err := tx.PrepareContext(ctx, q1)
	if err != nil {
		return errors.Wrapf(err, "query prepare error: %s", q1)
	}
	stmtFindNodeID, err := tx.PrepareContext(ctx, `
	SELECT node_id FROM nodes WHERE ipv4 = $1 AND port = $2
`)
	if err != nil {
		return errors.Wrap(err, "query prepare error")
	}
	q2 := `
	INSERT INTO flows
	(direction, source_node_id, destination_node_id, connections, updated)
	VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP)
	ON CONFLICT (source_node_id, destination_node_id, direction) 
	DO UPDATE SET
	direction=$1, source_node_id=$2, destination_node_id=$3, connections=$4, updated=CURRENT_TIMESTAMP
`
	stmt2, err := tx.PrepareContext(ctx, q2)
	if err != nil {
		return errors.Wrapf(err, "query prepare error: %s", q2)
	}

	for _, flow := range flows {
		if flow.Local.Addr == "127.0.0.1" || flow.Local.Addr == "::1" || flow.Peer.Addr == "127.0.0.1" || flow.Peer.Addr == "::1" {
			continue
		}
		var localNodeid, peerNodeid int64
		err := stmt1.QueryRowContext(ctx, flow.Local.Addr, flow.Local.PortInt()).Scan(&localNodeid)
		if err == sql.ErrNoRows {
			err = stmtFindNodeID.QueryRowContext(ctx, flow.Local.Addr, flow.Local.PortInt()).Scan(&localNodeid)
		}
		if err != nil {
			return errors.Wrapf(err, "query error")
		}
		err = stmt1.QueryRowContext(ctx, flow.Peer.Addr, flow.Peer.PortInt()).Scan(&peerNodeid)
		if err == sql.ErrNoRows {
			err = stmtFindNodeID.QueryRowContext(ctx, flow.Peer.Addr, flow.Peer.PortInt()).Scan(&peerNodeid)
		}
		if err != nil {
			return errors.Wrapf(err, "query error")
		}
		if flow.Direction == tcpflow.FlowActive {
			_, err = stmt2.ExecContext(ctx, flow.Direction.String(), localNodeid, peerNodeid, flow.Connections)
		} else if flow.Direction == tcpflow.FlowPassive {
			_, err = stmt2.ExecContext(ctx, flow.Direction.String(), peerNodeid, localNodeid, flow.Connections)
		}
		if err != nil {
			return errors.Wrapf(err, "query error")
		}
	}
	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "transaction commit error")
	}
	return nil
}
