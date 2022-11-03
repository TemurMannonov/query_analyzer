package config

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseConfig(t *testing.T) {
	cfgExpected := getConfig()

	cfg := ParseConfig(".")
	equal := reflect.DeepEqual(cfgExpected, cfg)
	require.Equal(t, true, equal)
}

func getConfig() Config {
	cfg := Config{
		HttpPort: "8000",
		Postgres: PostgresConfig{
			Host:     "localhost",
			Port:     "5432",
			User:     "postgres",
			Password: "postgres",
			Database: "postgres",
		},
	}
	os.Setenv("POSTGRES_HOST", cfg.Postgres.Host)
	os.Setenv("POSTGRES_PORT", cfg.Postgres.Port)
	os.Setenv("POSTGRES_USER", cfg.Postgres.User)
	os.Setenv("POSTGRES_PASSWORD", cfg.Postgres.Password)
	os.Setenv("POSTGRES_DATABASE", cfg.Postgres.Database)

	os.Setenv("HTTP_PORT", cfg.HttpPort)

	return cfg
}
