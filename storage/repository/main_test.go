package repository_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/TemurMannonov/query_analyzer/config"
	"github.com/TemurMannonov/query_analyzer/storage"
	"github.com/TemurMannonov/query_analyzer/storage/repository"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	postgresConn *sqlx.DB
	err          error
	cfg          config.Config
	strg         storage.DBRepositoryI
)

func TestMain(m *testing.M) {
	cfg = config.ParseConfig("../..")

	conStr := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
	)

	postgresConn, err = sqlx.Open("postgres", conStr)

	if err != nil {
		log.Fatal(err)
	}

	strg = repository.NewDBRepository(postgresConn)

	os.Exit(m.Run())
}
