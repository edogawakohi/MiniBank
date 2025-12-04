package main

import (
	"context"
	"log"
	"minibank/api"
	db "minibank/db/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbSource      = "postgres://root:123@localhost:5432/mini_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {

	config, err := pgxpool.ParseConfig(dbSource)
	if err != nil {
		log.Fatal("cannot parse config:", err)
	}

	conn, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal("Cannot connect database: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Cannot start server ", err)
	}
}
