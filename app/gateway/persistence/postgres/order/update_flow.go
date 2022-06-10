package order

import (
	"context"
	"errors"
	"github.com/jackc/pgconn"
	"gomies/app/core/entities/order"
	"gomies/app/gateway/persistence/postgres"
	"gomies/app/gateway/persistence/postgres/transaction"
	"gomies/app/sdk/fault"
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
				return order.FlowUpdate{}, fault.Wrap(fault.ErrAlreadyExists).
					Describe("the flow id or status provided seems to already exist").Params(map[string]interface{}{
					"id": f.ID, "order_id": f.OrderID, "status": f.Status,
				})
			}

			if pgErr.Code == postgres.NonexistentFK && pgErr.ConstraintName == postgres.OrderFlowFK {
				return order.FlowUpdate{}, fault.Wrap(fault.ErrNotFound).
					Describe("the order id provided seems to not exist").Params(map[string]interface{}{
					"id": f.ID, "order_id": f.OrderID, "status": f.Status,
				})
			}
		}
		return order.FlowUpdate{}, err
	}

	return f, nil
}
