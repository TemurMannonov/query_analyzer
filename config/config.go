package config

import (
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	HttpPort string
	Postgres PostgresConfig
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func ParseConfig() Config {
	godotenv.Load() // load .env file if it exists

	conf := viper.New()
	conf.AutomaticEnv()
	conf.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	cfg := Config{
		HttpPort: conf.GetString("http.port"),
		Postgres: PostgresConfig{
			Host:     conf.GetString("postgres.host"),
			Port:     conf.GetString("postgres.port"),
			User:     conf.GetString("postgres.user"),
			Password: conf.GetString("postgres.password"),
			Database: conf.GetString("postgres.database"),
		},
	}

	return cfg
}
