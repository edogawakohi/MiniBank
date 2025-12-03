package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbSource = "postgres://root:123@localhost:5432/mini_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {

	var err error

	config, err := pgxpool.ParseConfig(dbSource)
	if err != nil {
		log.Fatal("cannot parse config:", err)
	}

	testDB, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal("Cannot connect database: ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
