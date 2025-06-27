package main

import (
	"database/sql"
	"log"

	"github.com/Yarik7610/simplebank/api"
	db "github.com/Yarik7610/simplebank/db/sqlc"
	"github.com/Yarik7610/simplebank/utils"
	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatalf("can't load config: %v", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("can't connect to db: %s", err)
	}

	if err = conn.Ping(); err != nil {
		log.Fatalf("can't ping db: %s", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatalf("can't run server on port: %s, error: %s", config.HTTPServerAddress, err)
	}
}
