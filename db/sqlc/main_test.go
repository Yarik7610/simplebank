package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/Yarik7610/simplebank/utils"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var conn *sql.DB

func TestMain(t *testing.M) {
	var err error

	config, err := utils.LoadConfig("../../")
	if err != nil {
		log.Fatalf("can't load config: %v", err)
	}

	conn, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("can't connect to db: %s", err)
	}

	if err = conn.Ping(); err != nil {
		log.Fatalf("can't ping db: %s", err)
	}

	testQueries = New(conn)

	os.Exit(t.Run())
}
