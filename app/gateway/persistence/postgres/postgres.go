package postgres

import (
	"context"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"log"
)

type Config struct {
	User      string
	Password  string
	Host      string
	Port      string
	Name      string
	SSLMode   string
	Migration string
}

const MigrationPath = "../migrations"

func CreateDatabaseURL(user, password, host, port, name, SSLMode string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, password, host, port, name, SSLMode)
}

func ConnectAndMount(ctx context.Context, c Config) (*pgxpool.Pool, error) {
	url := CreateDatabaseURL(c.User, c.Password, c.Host, c.Port, c.Name, c.SSLMode)
	conn, err := NewConnection(ctx, url)
	if err != nil {
		return nil, err
	}

	var mig = MigrationPath
	if c.Migration != "" {
		mig = c.Migration
	}

	err = Migrate(url, mig)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// Migrate writes up the schema to a database
func Migrate(db string, filesPath string) error {
	path := fmt.Sprintf("file://%v", filesPath)
	migration, err := migrate.New(path, db)
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
