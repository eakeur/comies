package orders

import (
	"comies/app/core/order"
	"comies/app/core/types"
	"comies/app/data/conn"
	"context"
	"errors"

	"github.com/jackc/pgconn"
)

func UpdateFlow(ctx context.Context, f order.FlowUpdate) (order.FlowUpdate, error) {
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
	_, err := conn.ExecFromContext(ctx, script,
		f.ID,
		f.OrderID,
		f.OccurredAt,
		f.Status,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == conn.DuplicateError &&
				(pgErr.ConstraintName == conn.OrderFlowPK || pgErr.ConstraintName == conn.OrderStatusUK) {
				return order.FlowUpdate{}, types.ErrAlreadyExists
			}

			if pgErr.Code == conn.NonexistentFK && pgErr.ConstraintName == conn.OrderFlowFK {
				return order.FlowUpdate{}, types.ErrNotFound
			}
		}
		return order.FlowUpdate{}, err
	}

	return f, nil
}
