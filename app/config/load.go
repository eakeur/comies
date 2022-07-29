package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Load() (Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Could not load .env file: continuing fetching variables from environment")
	}

	return Config{
		Database: Database{
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASS"),
			Host:     os.Getenv("DB_HOST"),
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
