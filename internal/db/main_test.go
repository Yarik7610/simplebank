package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

const (
	dbDriver = "postgres"
	dbSource = "postgres://root:secret@localhost:5433/simple_bank?sslmode=disable"
)

func TestMain(t *testing.M) {
	var err error

	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalf("cannot connect to database: %s", err)
	}

	testQueries = New(testDB)

	os.Exit(t.Run())
}
