package orders

import (
	"comies/app/core/ordering"
	"comies/app/core/types"
	"comies/app/data/conn"
	"context"
	"errors"

	"github.com/jackc/pgconn"
)

func UpdateFlow(ctx context.Context, f ordering.Flow) error {
	const script = `
		insert into orders_flow (
			order_id, 
			occurred_at,
			status
		) values (
			$1, $2, $3, $4
		)
	`
	_, err := conn.ExecFromContext(ctx, script,
		f.OrderID,
		f.OccurredAt,
		f.Status,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == conn.DuplicateError &&
				(pgErr.ConstraintName == conn.OrderFlowPK || pgErr.ConstraintName == conn.OrderStatusUK) {
				return types.ErrAlreadyExists
			}

			if pgErr.Code == conn.NonexistentFK && pgErr.ConstraintName == conn.OrderFlowFK {
				return types.ErrNotFound
			}
		}
		return err
	}

	return nil
}
