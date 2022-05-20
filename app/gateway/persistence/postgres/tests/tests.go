package tests

import (
	"context"
	"gomies/app/gateway/persistence/postgres"
	"gomies/app/sdk/session"
	"gomies/app/sdk/types"
	"log"
	"testing"
)

var (
	DefaultStore = types.Store{StoreID: 1}
	container    *Container
)

func SetupTest(m *testing.M) int {

	ctx := context.Background()
	container = &Container{
		config: postgres.Config{
			User:     "postgres",
			Password: "postgres",
			Host:     "localhost",
			Name:     "postgres",
			SSLMode:  "disable",
		},
		database: &Database{
			TestName: "Main test",
			Context:  ctx,
		},
	}
	err := container.create()
	if err != nil {
		log.Fatal(err)
	}

	defer container.teardown()
	return m.Run()
}

func NewTestDatabase(t *testing.T, ctx context.Context) *Database {
	test := t.Name()

	db := &Database{
		Test:     t,
		TestName: test,
		Context:  ctx,
	}

	db.name()
	err := db.mount()
	if err != nil {
		log.Printf("Could not mount database for test %s: %v", db.TestName, err)
		return nil
	}
	err = db.connect()
	if err != nil {
		log.Printf("Could not connect to database for test %s: %v", db.TestName, err)
		return nil
	}

	return db
}

func FetchTestDB(t *testing.T, callbacks ...Callback) (context.Context, *Database) {
	ctx := session.Session{StoreID: DefaultStore.StoreID}.WithContext(context.Background())
	db := NewTestDatabase(t, ctx)
	db.runCallbacks(callbacks)
	return ctx, db
}

func FetchTestTX(t *testing.T, callbacks ...Callback) (context.Context, *Database) {
	ctx := session.Session{StoreID: DefaultStore.StoreID}.WithContext(context.Background())
	db := NewTestDatabase(t, ctx)
	ctx, _ = db.Transaction(ctx)
	db.runCallbacks(callbacks)
	return ctx, db
}
