package orders

import (
	"comies/app/core/id"
	"comies/app/core/order"
	"comies/app/data/conn"
	"context"
)

func ListFlow(ctx context.Context, orderID id.ID) ([]order.FlowUpdate, error) {
	const script = `
		select
			f.id,
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

	items := make([]order.FlowUpdate, 0)
	for rows.Next() {
		var o order.FlowUpdate
		if err := rows.Scan(
			&o.ID,
			&o.OrderID,
			&o.OccurredAt,
			&o.Status,
		); err != nil {
			continue
		}
		o.OccurredAt = o.OccurredAt.UTC()
		items = append(items, o)
	}

	return items, nil
}
