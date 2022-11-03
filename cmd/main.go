package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/TemurMannonov/query_analyzer/api"
	"github.com/TemurMannonov/query_analyzer/config"
	"github.com/TemurMannonov/query_analyzer/storage/repository"
)

func main() {
	cfg := config.ParseConfig()

	psqlUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
	)

	psqlConn, err := sqlx.Connect("postgres", psqlUrl)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	strg := repository.NewDBRepository(psqlConn)

	apiServer := api.NewServer(&cfg, strg)

	err = apiServer.Start(cfg.HttpPort)
	if err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

	log.Print("Server stopped")
}
