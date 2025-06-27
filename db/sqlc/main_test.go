package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var conn *sql.DB

const (
	dbDriver = "postgres"
	dbSource = "postgres://root:secret@localhost:5433/simple_bank?sslmode=disable"
)

func TestMain(t *testing.M) {
	var err error

	conn, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalf("can't connect to db: %s", err)
	}

	if err = conn.Ping(); err != nil {
		log.Fatalf("can't ping db: %s", err)
	}

	testQueries = New(conn)

	os.Exit(t.Run())
}
