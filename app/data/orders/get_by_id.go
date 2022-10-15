package orders

import (
	"comies/app/core/id"
	"comies/app/core/ordering"
	"comies/app/core/types"
	"comies/app/data/conn"
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
)

func GetByID(ctx context.Context, id id.ID) (ordering.Order, error) {
	const script = `
		select
			o.id,
        	max(o.placed_at),
        	max(o.delivery_mode),
        	max(o.observations),
			max(s.status),
			max(o.customer_name),
			max(o.customer_phone),
			max(o.customer_address)
		from
			orders o
		inner join 
			orders_statuses s on o.id = s.order_id
		-- left join items i on o.id = i.order_id
		where
			o.id = $1
		group by 
			o.id
	`

	row, err := conn.QueryRowFromContext(ctx, script, id)
	if err != nil {
		return ordering.Order{}, nil
	}

	var o ordering.Order
	if err := row.Scan(
		&o.ID,
		&o.PlacedAt,
		&o.DeliveryType,
		&o.Observations,
		&o.Status,
		&o.Customer.Name,
		&o.Customer.Phone,
		&o.Customer.Address,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ordering.Order{}, types.ErrNotFound
		}
		return ordering.Order{}, err
	}

	o.PlacedAt = o.PlacedAt.UTC()
	return o, nil
}
