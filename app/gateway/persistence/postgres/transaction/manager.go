package transaction

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

//go:generate moq -fmt goimports -out manager_mock.go . Manager:ManagerMock
type (
	Manager interface {
		// Begin starts a transaction and stores an object to it in this context
		Begin(context.Context) context.Context

		// Commit commits a transaction in this context
		Commit(context.Context)

		// Rollback rollbacks a transaction in this context
		Rollback(context.Context)

		// End finishes a transaction, calling commit if there are no errors or calling rollback if there are any
		End(ctx context.Context)
	}

	manager struct {
		pool *pgxpool.Pool
	}
)

func (m *manager) Begin(ctx context.Context) context.Context {
	tx, err := m.pool.Begin(ctx)
	if err != nil {
		return ctx
	}

	return context.WithValue(ctx, key, tx)
}

func (m *manager) Commit(ctx context.Context) {
	tx, _ := FromContext(ctx)
	tx.Commit(ctx)
}

func (m *manager) Rollback(ctx context.Context) {
	tx, _ := FromContext(ctx)
	tx.Rollback(ctx)
}

func (m *manager) End(ctx context.Context) {
	tx, _ := FromContext(ctx)
	err := tx.Commit(ctx)
	if err != nil {
		tx.Rollback(ctx)
	}
}

func NewManager(pool *pgxpool.Pool) Manager {
	return &manager{
		pool: pool,
	}
}
