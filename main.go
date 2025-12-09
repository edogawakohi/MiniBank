package main

import (
	"context"
	"log"
	"minibank/api"
	db "minibank/db/sqlc"
	"minibank/utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	dbConfig, err := pgxpool.ParseConfig(config.DBSource)
	if err != nil {
		log.Fatal("cannot parse config:", err)
	}

	conn, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		log.Fatal("Cannot connect database: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server ", err)
	}
}
