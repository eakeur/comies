package order

import (
	"context"
	"gomies/app/core/entities/order"
	"gomies/app/gateway/persistence/postgres/query"
	"gomies/app/sdk/fault"
)

func (a actions) List(ctx context.Context, filter order.Filter) ([]order.Order, int, error) {
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
		%where_query%
		group by 
			o.id
	`
	q := query.NewQuery(script).
		Where(!filter.PlacedAfter.IsZero(), "o.placed_at >= $%v", filter.PlacedAfter).And().
		Where(!filter.PlacedBefore.IsZero(), "o.placed_at <= $%v", filter.PlacedBefore).And().
		Where(filter.Status != "", "s.status = $%v", filter.Status).And().
		Where(filter.DeliveryMode != "", "o.delivery_mode = $%v", filter.DeliveryMode)

	rows, err := a.db.Query(ctx, q.Script(), q.Args...)
	if err != nil {
		return nil, 0, fault.Wrap(err)
	}

	items := make([]order.Order, 0)
	for rows.Next() {
		var o order.Order
		if err := rows.Scan(
			&o.ID,
			&o.Identification,
			&o.PlacedAt,
			&o.DeliveryMode,
			&o.Observations,
			&o.Status,
			&o.FinalPrice,
		); err != nil {
			continue
		}
		o.PlacedAt = o.PlacedAt.UTC()
		items = append(items, o)
	}

	return items, 0, nil
}
