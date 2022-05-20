package tests

import (
	"context"
	"crypto/rand"
	"fmt"
	"gomies/app/gateway/persistence/postgres"
	"gomies/app/gateway/persistence/postgres/transaction"
	"log"
	"math"
	"math/big"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

type (
	Database struct {
		Name     string
		TestName string
		Pool     *pgxpool.Pool
		Context  context.Context
		Test     *testing.T
		txEnd    func(context.Context)
	}

	Callback = func(ctx context.Context, db *Database, t *testing.T)
)

func (d *Database) Transaction(ctx context.Context) (context.Context, func(context.Context)) {
	man := transaction.NewManager(d.Pool)
	d.Context = man.Begin(ctx)
	d.txEnd = man.End
	return d.Context, man.End
}

func (d Database) Drop(callbacks ...Callback) {
	if d.txEnd != nil {
		d.txEnd(d.Context)
	}

	d.runCallbacks(callbacks)

	const script = `drop database if exists %s`
	d.Pool.Close()

	_, err := container.database.Pool.Exec(d.Context, fmt.Sprintf(script, d.Name))
	if err != nil {
		log.Printf("Could not drop database for test %s: %v", d.TestName, err)
	}
}

func (d *Database) name() {
	n, _ := rand.Int(rand.Reader, big.NewInt(math.MaxInt32))
	d.Name = fmt.Sprintf("database_%d", n)
}

func (d *Database) mount() error {
	_, err := container.database.Pool.Exec(d.Context, "create database "+d.Name)
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) connect() error {
	cfg := container.config
	cfg.Name = d.Name

	p, err := postgres.ConnectAndMount(d.Context, cfg)
	if err != nil {
		return err
	}

	d.Pool = p

	return nil
}

func (d *Database) runCallbacks(cb []Callback) {
	for _, callback := range cb {
		if callback != nil {
			callback(d.Context, d, d.Test)
		}
	}
}
