package order

import (
	"comies/app/core/entities/order"
	"comies/app/core/types"
	"comies/app/gateway/persistence/postgres"
	"comies/app/gateway/persistence/postgres/transaction"
	"context"
	"errors"

	"github.com/jackc/pgconn"
)

func (a actions) Create(ctx context.Context, o order.Order) (order.Order, error) {
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

	if _, err := transaction.ExecFromContext(ctx, script,
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
			if pgErr.Code == postgres.DuplicateError && pgErr.ConstraintName == postgres.OrderIDPK {
				return order.Order{}, types.ErrAlreadyExists
			}
		}

		return order.Order{}, err
	}

	return o, nil
}
