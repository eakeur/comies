package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Load() Config {

	for _, filename := range []string{
		".env.local", ".env",
	} {
		err := godotenv.Load(filename)
		if err == nil {
			break
		}
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
			Address:     os.Getenv("SERVER_ADDRESS"),
			CORSOrigins: os.Getenv("SERVER_CORS_ORIGINS"),
		},
		Logger: Logger{
			Environment: os.Getenv("ENVIRONMENT"),
		},
		IDGeneration: IDGeneration{
			NodeNumber: os.Getenv("ID_GENERATION_NODE_NUMBER"),
		},
	}
}
