package postgres

import (
	"context"
	"crypto/rand"
	"fmt"
	"math"
	"math/big"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type executor interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
}

func url(port string, database string) string {
	return fmt.Sprintf(dockerPostgresURL, port, database)
}

func useConnection(url string) (*pgx.Conn, error) {
	ctx := context.Background()

	pg, err := pgx.Connect(ctx, url)
	if err != nil {
		return nil, err
	}

	return pg, pg.Ping(ctx)
}

func createName() string {
	n, _ := rand.Int(rand.Reader, big.NewInt(math.MaxInt32))
	return fmt.Sprintf("database_%d", n)
}

func createDatabase(ctx context.Context, conn executor, name string) error {
	_ = removeDatabase(ctx, conn, name)

	if _, err := conn.Exec(context.Background(), fmt.Sprintf(`create database %s;`, name)); err != nil {
		return fmt.Errorf("failed creating test database %s: %w", name, err)
	}
	return nil
}

func createTemplate(ctx context.Context, conn executor, name string) error {
	_ = removeDatabase(ctx, conn, name)

	if _, err := conn.Exec(context.Background(), fmt.Sprintf("create database %[1]s_template template %[1]s;", name)); err != nil {
		return fmt.Errorf("failed creating template database %s: %w", name, err)
	}

	return nil
}

func useTemplate(ctx context.Context, conn executor, name, template string) error {
	_ = removeDatabase(ctx, conn, name)

	if _, err := conn.Exec(context.Background(), fmt.Sprintf(`create database %s template %s_template;`, name, template)); err != nil {
		return fmt.Errorf("failed creating test database %s from template %s: %w", name, template, err)
	}
	return nil
}

func removeDatabase(ctx context.Context, conn executor, name string) error {
	if _, err := conn.Exec(context.Background(), fmt.Sprintf(`drop database if exists %s;`, name)); err != nil {
		return fmt.Errorf("failed dropping test database %s: %w", name, err)
	}

	return nil
}
