package configs

import (
	"os"
)

type Config struct {
	DBURL string
}

func (config *Config) LoadConfig() {
	config.DBURL = os.Getenv("DB_URL")
}
