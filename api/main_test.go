package api

import (
	"os"
	"testing"

	"github.com/TemurMannonov/query_analyzer/config"
	"github.com/TemurMannonov/query_analyzer/storage"
)

func newTestServer(t *testing.T, strg storage.DBRepositoryI) *Server {
	cfg := config.ParseConfig(".")

	server := NewServer(&cfg, strg)

	return server
}

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}
