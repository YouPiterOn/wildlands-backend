package env

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	return godotenv.Load(".env")
}

type EnvKey string

const (
	PostgresURL EnvKey = "POSTGRES_URL"
)

func (k EnvKey) GetValue() string {
	return os.Getenv(string(k))
}
