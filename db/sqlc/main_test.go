package db

import (
	"context"
	"log"
	"minibank/utils"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {

	var err error

	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}
	configDB, err := pgxpool.ParseConfig(config.DBSource)
	if err != nil {
		log.Fatal("cannot parse config:", err)
	}

	testDB, err = pgxpool.NewWithConfig(context.Background(), configDB)
	if err != nil {
		log.Fatal("Cannot connect database: ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
