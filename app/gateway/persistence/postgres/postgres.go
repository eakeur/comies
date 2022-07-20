package postgres

import (
	"context"
	"embed"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
	SSLMode  string
}

const MigrationPath = "migrations"

//go:embed migrations
var MigrationsFS embed.FS

func CreateDatabaseURL(user, password, host, port, name, SSLMode string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, password, host, port, name, SSLMode)
}

func ConnectAndMount(ctx context.Context, c Config) (*pgxpool.Pool, error) {
	url := CreateDatabaseURL(c.User, c.Password, c.Host, c.Port, c.Name, c.SSLMode)
	return ConnectAndMountURL(ctx, url)
}

func ConnectAndMountURL(ctx context.Context, url string) (*pgxpool.Pool, error) {
	conn, err := NewConnection(ctx, url)
	if err != nil {
		return nil, err
	}

	err = Migrate(conn.Config().ConnConfig)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// Migrate writes up the schema to a database
func Migrate(cfg *pgx.ConnConfig) error {
	driver, err := postgres.WithInstance(stdlib.OpenDB(*cfg), &postgres.Config{
		DatabaseName: cfg.Database,
	})
	if err != nil {
		return fmt.Errorf("failed to get migrations postgres driver: %w", err)
	}

	source, err := httpfs.New(http.FS(MigrationsFS), MigrationPath)
	if err != nil {
		return fmt.Errorf("failed to create migration httpfs driver: %w", err)
	}

	migration, err := migrate.NewWithInstance("httpfs", source, cfg.Database, driver)
	if err != nil {
		return fmt.Errorf("[migrations] failed to create migrate source instance: %w", err)
	}

	if err != nil {
		return err
	}
	defer func(migration *migrate.Migrate) {
		err, _ := migration.Close()
		if err != nil {
			log.Printf("Error deferring migration resources: %v", err)
		}

	}(migration)

	if err != nil {
		return errors.Wrap(err, "could not start migration")
	}

	if err := migration.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return errors.Wrap(err, "error uploading migration")
	}

	return nil
}

// NewConnection creates a connection object
func NewConnection(ctx context.Context, url string) (*pgxpool.Pool, error) {
	pgxConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, err
	}

	db, err := pgxpool.ConnectConfig(ctx, pgxConfig)
	if err != nil {
		return nil, err
	}

	return db, nil
}
