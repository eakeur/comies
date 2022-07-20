package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Load() (Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return Config{}, err
	}

	return Config{
		Database: Database{
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASS"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Name:     os.Getenv("DB_NAME"),
			SSLMode:  os.Getenv("DB_SSL"),
			URL:      os.Getenv("DATABASE_URL"),
		},
		Server: Server{
			ListenPort: os.Getenv("PORT"),
		},
		Logger: Logger{
			Environment: os.Getenv("ENVIRONMENT"),
		},
		IDGeneration: IDGeneration{
			NodeNumber: os.Getenv("ID_GENERATION_NODE_NUMBER"),
		},
	}, nil
}
