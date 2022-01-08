package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/BigListRyRy/harbourlivingapi/api"
	db "github.com/BigListRyRy/harbourlivingapi/db/sqlc"
	"github.com/BigListRyRy/harbourlivingapi/util"
	_ "github.com/lib/pq"
)

func startOld() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot local config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalln("cannot connect to database, ", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(store, *config)
	if err != nil {
		log.Fatal("unable to create a new server ", err)
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = config.Port
	}
	err = server.Start(":" + port)
	if err != nil {
		log.Fatal("unable to start server", err)
	}
}
