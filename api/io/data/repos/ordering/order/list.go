package order

import (
	"comies/core/ordering/order"
	"comies/io/data/conn"
	"comies/io/data/query"
	"context"
)

func (a actions) List(ctx context.Context, filter order.Filter) ([]order.Order, error) {
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
		%where_query%
	`
	statuses := make([]interface{}, len(filter.Status))
	for i, status := range filter.Status {
		statuses[i] = status
	}

	q := query.NewQuery(script).
		Where(!filter.PlacedAfter.IsZero(), "o.placed_at >= $%v", filter.PlacedAfter).And().
		Where(!filter.PlacedBefore.IsZero(), "o.placed_at <= $%v", filter.PlacedBefore).And().
		Where(len(statuses) > 0, "s.status in (%v)", filter.Status).And().
		Where(filter.DeliveryType != 0, "o.delivery_mode = $%v", filter.DeliveryType)

	rows, err := conn.QueryFromContext(ctx, q.Script(), q.Args...)
	if err != nil {
		return nil, err
	}

	return conn.ScanRows(rows,
		func(scan conn.Scan, o order.Order) error {
			return scan(
				&o.ID,
				&o.PlacedAt,
				&o.DeliveryType,
				&o.Observations,
				&o.CustomerName,
				&o.CustomerAddress,
				&o.CustomerPhone,
			)
		},
	)
}
