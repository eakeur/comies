package order

import (
	"comies/app/core/entities/order"
	"comies/app/core/throw"
	"comies/app/gateway/persistence/postgres"
	"comies/app/gateway/persistence/postgres/transaction"
	"context"
	"errors"

	"github.com/jackc/pgconn"
)

func (a actions) UpdateFlow(ctx context.Context, f order.FlowUpdate) (order.FlowUpdate, error) {
	const script = `
		insert into orders_flow (
			id,
			order_id, 
			occurred_at,
			status
		) values (
			$1, $2, $3, $4
		)
	`
	_, err := transaction.ExecFromContext(ctx, script,
		f.ID,
		f.OrderID,
		f.OccurredAt,
		f.Status,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == postgres.DuplicateError &&
				(pgErr.ConstraintName == postgres.OrderFlowPK || pgErr.ConstraintName == postgres.OrderStatusUK) {
				return order.FlowUpdate{}, throw.ErrAlreadyExists
			}

			if pgErr.Code == postgres.NonexistentFK && pgErr.ConstraintName == postgres.OrderFlowFK {
				return order.FlowUpdate{}, throw.ErrNotFound
			}
		}
		return order.FlowUpdate{}, err
	}

	return f, nil
}
