package tests

import (
	"comies/app/gateway/persistence/postgres"
	"comies/app/gateway/persistence/postgres/transaction"
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math"
	"math/big"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

type (
	Database struct {
		*pgxpool.Pool
	}

	Callback = func(ctx context.Context, db *Database, t *testing.T)
)

func NewTestDatabase(t *testing.T, ctx context.Context, bef Callback, aft Callback, withTX bool) (*Database, context.Context) {
	t.Helper()

	if conn == nil {
		t.FailNow()
	}

	var (
		tx   transaction.Manager
		db   *Database
		test = t.Name()
		cfg  = conn.Config().ConnConfig
	)

	n, _ := rand.Int(rand.Reader, big.NewInt(math.MaxInt32))
	name := fmt.Sprintf("database_%d", n)

	_, _ = conn.Exec(ctx, fmt.Sprintf("drop database if exists %s", name))

	script := fmt.Sprintf("create database %s template %s_template", name, conn.Config().ConnConfig.Database)
	_, err := conn.Exec(ctx, script)
	if err != nil {
		t.Errorf("Could not create database for test %s: %v", test, err)
		t.FailNow()
	}

	pool, err := postgres.NewConnection(ctx, postgres.CreateDatabaseURL(cfg.User, cfg.Password, fmt.Sprintf("%s:%d", cfg.Host, cfg.Port), name, "disable"))
	if err != nil {
		t.Errorf("Could not connect to database for test %s: %v", test, err)
		t.FailNow()
	}

	db = &Database{
		Pool: pool,
	}

	if withTX {
		tx = transaction.NewManager(pool)
		ctx = tx.Begin(ctx)
	}

	if bef != nil {
		bef(ctx, db, t)
	}

	t.Cleanup(func() {
		if withTX {
			tx.End(ctx)
		}

		if aft != nil {
			aft(ctx, db, t)
		}

		const script = `drop database %s`
		pool.Close()

		_, err := conn.Exec(ctx, fmt.Sprintf(script, name))
		if err != nil {
			log.Printf("Could not drop database for test %s: %v", test, err)
		}
	})

	return db, ctx
}

func FetchTestDB(t *testing.T, callbacks ...Callback) (context.Context, *Database) {
	t.Helper()
	ctx := context.Background()
	bef, aft := checkCallbacks(callbacks)

	db, ctx := NewTestDatabase(t, ctx, bef, aft, false)
	return ctx, db
}

func FetchTestTX(t *testing.T, callbacks ...Callback) (context.Context, *Database) {
	t.Helper()
	ctx := context.Background()
	bef, aft := checkCallbacks(callbacks)

	db, ctx := NewTestDatabase(t, ctx, bef, aft, true)
	return ctx, db
}

func checkCallbacks(cbs []Callback) (bef Callback, aft Callback) {
	if len(cbs) > 0 {
		bef = cbs[0]
	}

	if len(cbs) > 1 {
		aft = cbs[1]
	}

	return
}
