package conn

import (
	"comies/config"
	"comies/telemetry"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/log/zapadapter"
	"github.com/jackc/pgx/v4/pgxpool"
)

func Connect(c config.Database) (*pgxpool.Pool, error) {
	url := c.URL

	if url == "" {
		url = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s", c.User, c.Password, c.Host, c.Name, c.SSLMode)
	}

	pgxConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, err
	}

	logger := telemetry.SQLLogger()
	if logger != nil {
		pgxConfig.ConnConfig.Logger = zapadapter.NewLogger(logger)
	}

	db, err := pgxpool.ConnectConfig(context.Background(), pgxConfig)
	if err != nil {
		return nil, err
	}

	return db, nil
}
