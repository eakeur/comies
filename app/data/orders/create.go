package orders

import (
	"comies/app/core/order"
	"comies/app/core/types"
	"comies/app/data/conn"
	"context"
	"errors"

	"github.com/jackc/pgconn"
)

func Create(ctx context.Context, o order.Order) (order.Order, error) {
	const script = `
		insert into orders (
			id,
			identification, 
			placed_at,
			delivery_mode,
			observations,
			address,
			phone
		) values (
			$1, $2, $3, $4, $5, $6, $7
		)
	`

	if _, err := conn.ExecFromContext(ctx, script,
		o.ID,
		o.Identification,
		o.PlacedAt,
		o.DeliveryMode,
		o.Observations,
		o.Address,
		o.Phone,
	); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == conn.DuplicateError && pgErr.ConstraintName == conn.OrderIDPK {
				return order.Order{}, types.ErrAlreadyExists
			}
		}

		return order.Order{}, err
	}

	return o, nil
}
