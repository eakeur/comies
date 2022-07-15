package order

import (
	"comies/app/core/entities/order"
	"comies/app/sdk/fault"
	"comies/app/sdk/types"
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
)

func (a actions) GetByID(ctx context.Context, id types.ID) (order.Order, error) {
	const script = `
		select
			o.id,
        	max(o.identification),
        	max(o.placed_at),
        	max(o.delivery_mode),
        	max(o.observations),
			max(s.status),
			coalesce(sum(i.price), 0) as price
		from
			orders o
		inner join 
			orders_statuses s on o.id = s.order_id
		left join items i on o.id = i.order_id
		where
			o.id = $1
		group by 
			o.id
	`

	row := a.db.QueryRow(ctx, script, id)

	var o order.Order
	if err := row.Scan(
		&o.ID,
		&o.Identification,
		&o.PlacedAt,
		&o.DeliveryMode,
		&o.Observations,
		&o.Status,
		&o.FinalPrice,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return order.Order{}, fault.Wrap(fault.ErrNotFound).
				Describe("the order id provided seems to not exist").Params(map[string]interface{}{
				"id": id,
			})
		}
		return order.Order{}, fault.Wrap(err)
	}

	o.PlacedAt = o.PlacedAt.UTC()
	return o, nil
}
