package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Load() Config {

	for _, filename := range []string{
		".env.local", ".env",
	} {
		err := godotenv.Load(filename)
		if err == nil {
			log.Printf("Loaded %s file", filename)
			break
		}

		log.Printf("Could not load %s file: continuing fetching variables", filename)
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
	}
}
