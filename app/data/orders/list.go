package orders

import (
	"comies/app/core/ordering"
	"comies/app/data/conn"
	"comies/app/data/query"
	"context"
)

func List(ctx context.Context, filter ordering.OrderFilter) ([]ordering.Order, error) {
	const script = `
		select
			o.id,
        	max(o.placed_at),
        	max(o.delivery_mode),
        	max(o.observations),
			max(s.status),
			max(c.customer_name),
			max(c.customer_phone),
			max(c.customer_address),
		from
			orders o
		inner join 
			orders_statuses s on o.id = s.order_id
		-- left join items i on o.id = i.order_id
		%where_query%
		group by 
			o.id
	`
	var statuses []interface{}
	if l := len(filter.Status); l > 0 {
		statuses = make([]interface{}, l)
		for i, status := range filter.Status {
			statuses[i] = status
		}
	}

	q := query.NewQuery(script).
		Where(!filter.PlacedAfter.IsZero(), "o.placed_at >= $%v", filter.PlacedAfter).And().
		Where(!filter.PlacedBefore.IsZero(), "o.placed_at <= $%v", filter.PlacedBefore).And().
		Where(len(statuses) > 0, "s.status in (%v)", statuses...).And().
		Where(filter.DeliveryType != 0, "o.delivery_mode = $%v", filter.DeliveryType)

	rows, err := conn.QueryFromContext(ctx, q.Script(), q.Args...)
	if err != nil {
		return nil, err
	}

	items := make([]ordering.Order, 0)
	for rows.Next() {
		var o ordering.Order
		if err := rows.Scan(
			&o.ID,
			&o.PlacedAt,
			&o.DeliveryType,
			&o.Observations,
			&o.Status,
			&o.Customer.Name,
			&o.Customer.Phone,
			&o.Customer.Address,
		); err != nil {
			return nil, err
		}

		o.PlacedAt = o.PlacedAt.UTC()
		items = append(items, o)
	}

	return items, nil
}
