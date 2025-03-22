package config

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Server struct {
		Port int
	}
	Database struct {
		HOST string
		PORT int
		USER string
		PASS string
		NAME string
	}
	Logging struct {
		Level string
	}
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		slog.Error(fmt.Sprintf("Error occured while loading config: %s", err))
	}

	config := &Config{}

	port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		port = 8080
	}

	config.Server.Port = port

	config.Database.HOST = os.Getenv("DATABASE_HOST")
	config.Database.PORT, err = strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		port = 5432
	}

	config.Database.USER = os.Getenv("DATABASE_USER")
	config.Database.PASS = os.Getenv("DATABASE_PASS")
	config.Database.NAME = os.Getenv("DATABASE_DB")

	config.Logging.Level = os.Getenv("LOGGING_LEVEL")

	return config, nil
}
