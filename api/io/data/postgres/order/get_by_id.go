package order

import (
	"comies/core/ordering/order"
	"comies/core/types"
	"comies/io/data/postgres/conn"
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
)

func (a actions) GetByID(ctx context.Context, id types.ID) (order.Order, error) {
	const script = `
		select
			o.id,
        	o.placed_at,
        	o.delivery_type,
        	o.observations,
			o.customer_name,
			o.customer_address,
			o.customer_phone
		from
			orders o
		where
			o.id = $1
	`

	row, err := conn.QueryRowFromContext(ctx, script, id)
	if err != nil {
		return order.Order{}, err
	}

	var o order.Order
	if err := row.Scan(
		&o.ID,
		&o.PlacedAt,
		&o.DeliveryType,
		&o.Observations,
		&o.CustomerName,
		&o.CustomerAddress,
		&o.CustomerPhone,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return order.Order{}, types.ErrNotFound
		}
		return order.Order{}, err
	}

	return o, nil
}
