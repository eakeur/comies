package postgres

import (
	"comies/config"
	"comies/io/data/postgres/conn"
	"context"
	"testing"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DatabaseContextBuilder func(t *testing.T, useTX bool, callbacks ...Callback) context.Context

type Callback = func(ctx context.Context, t *testing.T)

func builder(c *pgx.Conn, port, templateName string) DatabaseContextBuilder {
	return func(t *testing.T, useTX bool, callbacks ...Callback) context.Context {
		t.Helper()

		test, ctx, name := t.Name(), context.Background(), createName()

		err := useTemplate(ctx, c, name, templateName)

		if err != nil {
			t.Errorf("%s: %v", test, err)
			t.FailNow()
		}

		pool, err := conn.Connect(config.Database{
			URL: url(port, name),
		})

		if err != nil {
			t.Errorf("Could not connect to database for test %s: %v", test, err)
			t.FailNow()
		}

		var db conn.Executer = pool
		if useTX {
			tx, err := pool.Begin(ctx)
			if err != nil {
				t.Fatal(err)
			}

			db = tx
		}

		ctx = conn.WithContext(ctx, db)

		l := len(callbacks)
		if l > 0 {
			callbacks[0](ctx, t)
		}

		if l > 1 {
			t.Cleanup(func() {
				callbacks[1](ctx, t)
			})
		}

		return ctx
	}
}

func (b DatabaseContextBuilder) Connection(t *testing.T) *pgxpool.Pool {
	t.Helper()

	exec, err := conn.FromContext(b(t, false))
	if err != nil {
		t.Fatal(err)
	}

	return exec.(*pgxpool.Pool)
}

func (b DatabaseContextBuilder) TX(t *testing.T, callback ...Callback) context.Context {
	t.Helper()

	return b(t, true, callback...)
}

func (b DatabaseContextBuilder) Pool(t *testing.T, callback ...Callback) context.Context {
	t.Helper()

	return b(t, false, callback...)
}
