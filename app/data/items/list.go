package items

import (
	"comies/app/core/ordering"
	"comies/app/core/types"
	"comies/app/data/conn"
	"context"
)

func List(ctx context.Context, orderID types.ID) ([]ordering.Item, error) {
	const script = `
		select
			id,
			order_id,
			status,
			product_id,
			quantity,
			observations
		from
			items
		where
			order_id = $1
	`

	rows, err := conn.QueryFromContext(ctx, script, orderID)
	if err != nil {
		return nil, err
	}

	items := make([]ordering.Item, 0)
	var it ordering.Item

	for rows.Next() {
		if err := rows.Scan(
			&it.ID,
			&it.OrderID,
			&it.Status,
			&it.ProductID,
			&it.Quantity,
			&it.Observations,
		); err != nil {
			return nil, err
		}

		items = append(items, it)
	}

	return items, nil

}
