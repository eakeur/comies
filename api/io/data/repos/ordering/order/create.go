package order

import (
	"comies/core/ordering/order"
	"comies/core/types"
	"comies/io/data/conn"
	"context"
	"errors"

	"github.com/jackc/pgconn"
)

func (a actions) Create(ctx context.Context, o order.Order) error {
	const script = `
		insert into orders (
			id, 
			placed_at,
			delivery_type,
			observations,
			customer_name,
			customer_address,
			customer_phone
		) values (
			$1, $2, $3, $4, $5, $6, $7
		)
	`

	if _, err := conn.ExecFromContext(ctx, script,
		o.ID,
		o.PlacedAt,
		o.DeliveryType,
		o.Observations,
		o.CustomerName,
		o.CustomerAddress,
		o.CustomerPhone,
	); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if conn.IsCode(pgErr, conn.DuplicateError) && conn.IsConstraint(pgErr, conn.OrderIDPK) {
				return types.ErrAlreadyExists
			}
		}

		return err
	}

	return nil
}
