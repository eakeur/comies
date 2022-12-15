package conn

import (
	"comies/telemetry"
	"embed"
	"errors"
	"fmt"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jackc/pgx/v4/stdlib"
)

const MigrationPath = "migrations"

//go:embed migrations
var migrationsFS embed.FS

// Migrate writes up the schema to a database
func Migrate(pool *pgxpool.Pool) error {

	cfg := pool.Config().ConnConfig

	driver, err := postgres.WithInstance(stdlib.OpenDB(*cfg), &postgres.Config{
		DatabaseName: cfg.Database,
	})
	if err != nil {
		return fmt.Errorf("failed to get migrations postgres driver: %w", err)
	}

	source, err := httpfs.New(http.FS(migrationsFS), MigrationPath)
	if err != nil {
		return fmt.Errorf("failed to create migration httpfs driver: %w", err)
	}

	migration, err := migrate.NewWithInstance("httpfs", source, cfg.Database, driver)
	if err != nil {
		return fmt.Errorf("[migrations] failed to create migrate source instance: %w", err)
	}

	log := telemetry.SQLLogger()
	if log != nil {
		migration.Log = logger{
			log: log,
		}
	}

	if err != nil {
		return err
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

func MigrationHandler(pool *pgxpool.Pool) (*migrate.Migrate, error) {

	cfg := pool.Config().ConnConfig

	driver, err := postgres.WithInstance(stdlib.OpenDB(*cfg), &postgres.Config{
		DatabaseName: cfg.Database,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get migrations postgres driver: %w", err)
	}

	source, err := httpfs.New(http.FS(migrationsFS), MigrationPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create migration httpfs driver: %w", err)
	}

	migration, err := migrate.NewWithInstance("httpfs", source, cfg.Database, driver)
	if err != nil {
		return nil, fmt.Errorf("[migrations] failed to create migrate source instance: %w", err)
	}

	log := telemetry.SQLLogger()
	if log != nil {
		migration.Log = logger{
			log: log,
		}
	}

	return migration, nil
}
