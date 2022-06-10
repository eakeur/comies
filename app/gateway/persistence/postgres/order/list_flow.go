package order

import (
	"context"
	"gomies/app/core/entities/order"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (a actions) ListFlow(ctx context.Context, orderID types.ID) ([]order.FlowUpdate, error) {
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

	rows, err := a.db.Query(ctx, script, orderID)
	if err != nil {
		return nil, fault.Wrap(err)
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
