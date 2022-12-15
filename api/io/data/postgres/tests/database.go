package tests

import (
	"comies/config"
	"comies/io/data/postgres/conn"
	"context"
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Callback = func(ctx context.Context, t *testing.T)

func NewDBConn(t *testing.T) *pgxpool.Pool {
	t.Helper()

	if Connection == nil {
		t.FailNow()
	}

	test, cfg, ctx := t.Name(), Connection.Config().ConnConfig, context.Background()

	n, _ := rand.Int(rand.Reader, big.NewInt(math.MaxInt32))
	name := fmt.Sprintf("database_%d", n)

	_, _ = Connection.Exec(ctx, fmt.Sprintf("drop database if exists %s", name))

	script := fmt.Sprintf("create database %s template %s_template", name, cfg.Database)
	_, err := Connection.Exec(ctx, script)
	if err != nil {
		t.Errorf("Could not create database for test %s: %v", test, err)
		t.FailNow()
	}

	pool, err := conn.Connect(config.Database{
		User:     cfg.User,
		Password: cfg.Password,
		Host:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Name:     name,
		SSLMode:  "disable",
	})
	if err != nil {
		t.Errorf("Could not connect to database for test %s: %v", test, err)
		t.FailNow()
	}

	return pool
}

func NewTestDB(t *testing.T) func(tx bool, cbs ...Callback) context.Context {
	ctx, pool := context.Background(), NewDBConn(t)
	return func(tx bool, cbs ...Callback) context.Context {
		var db conn.Executer = pool
		if tx {
			tx, err := pool.Begin(ctx)
			if err != nil {
				t.Fatal(err)
			}

			db = tx
		}

		ctx = conn.WithContext(ctx, db)

		l := len(cbs)
		if l > 0 {
			cbs[0](ctx, t)
		}

		if l > 1 {
			t.Cleanup(func() {
				cbs[1](ctx, t)
			})
		}

		return ctx
	}
}
