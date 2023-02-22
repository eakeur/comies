package postgres

import (
	"comies/data/conn/migrations"
	"errors"
	"fmt"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jackc/pgx/v4/stdlib"
)

const migration = "."

func mig(pool *pgxpool.Pool) error {

	cfg := pool.Config().ConnConfig

	driver, err := postgres.WithInstance(stdlib.OpenDB(*cfg), &postgres.Config{
		DatabaseName: cfg.Database,
	})
	if err != nil {
		return fmt.Errorf("failed to get migrations postgres driver: %w", err)
	}

	source, err := httpfs.New(http.FS(migrations.Migrations), migration)
	if err != nil {
		return fmt.Errorf("failed to create migration httpfs driver: %w", err)
	}

	migration, err := migrate.NewWithInstance("httpfs", source, cfg.Database, driver)
	if err != nil {
		return fmt.Errorf("[migrations] failed to create migrate source instance: %w", err)
	}

	defer migration.Close()

	if err != nil {
		return fmt.Errorf("could not start migration: %w", err)
	}

	if err := migration.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("error uploading migration: %w", err)
	}

	return nil
}
