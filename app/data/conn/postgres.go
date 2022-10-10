package conn

import (
	"comies/app/config"
	"context"
	"embed"
	"fmt"
	"log"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/pkg/errors"
)

type (
	logger struct{}
)

const MigrationPath = "migrations"

//go:embed migrations
var MigrationsFS embed.FS

var pool *pgxpool.Pool

func Connect(c config.Database) error {
	url := c.URL

	if url == "" {
		url = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s", c.User, c.Password, c.Host, c.Name, c.SSLMode)
	}

	pgxConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return err
	}

	db, err := pgxpool.ConnectConfig(context.Background(), pgxConfig)
	if err != nil {
		return err
	}

	pool = db

	return nil
}

// Migrate writes up the schema to a database
func Migrate() error {

	cfg := pool.Config().ConnConfig

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

	migration.Log = logger{}

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
func NewConnection(url string) error {
	pgxConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return err
	}

	db, err := pgxpool.ConnectConfig(context.Background(), pgxConfig)
	if err != nil {
		return err
	}

	pool = db

	return nil
}

func (l logger) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (l logger) Verbose() bool {
	return true
}
