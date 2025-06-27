package main

import (
	"database/sql"
	"log"

	"github.com/Yarik7610/simplebank/api"
	db "github.com/Yarik7610/simplebank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgres://root:secret@localhost:5433/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalf("can't connect to db: %s", err)
	}

	if err = conn.Ping(); err != nil {
		log.Fatalf("can't ping db: %s", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatalf("can't run server on port: %s, error: %s", serverAddress, err)
	}
}
