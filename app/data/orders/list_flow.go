package orders

import (
	"comies/app/core/ordering"
	"comies/app/core/types"
	"comies/app/data/conn"
	"context"
)

func ListFlow(ctx context.Context, orderID types.ID) ([]ordering.Flow, error) {
	const script = `
		select
			f.order_id,
			f.occurred_at,
			f.status
		from
			orders_flow f
		where
			f.order_id = $1
	`

	rows, err := conn.QueryFromContext(ctx, script, orderID)
	if err != nil {
		return nil, err
	}

	items := make([]ordering.Flow, 0)
	for rows.Next() {
		var o ordering.Flow
		if err := rows.Scan(
			&o.OrderID,
			&o.OccurredAt,
			&o.Status,
		); err != nil {
			return nil, err
		}

		o.OccurredAt = o.OccurredAt.UTC()
		items = append(items, o)
	}

	return items, nil
}
