package item

import (
	"context"
	"gomies/app/core/entities/item"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (a actions) List(ctx context.Context, orderID types.ID) ([]item.Item, error) {
	const script = `
		select
			id,
			order_id,
			status,
            price,
			product_id,
			quantity,
			observations
		from
			items
		where
			order_id = $1
	`

	rows, err := a.db.Query(ctx, script, orderID)
	if err != nil {
		return nil, fault.Wrap(err)
	}

	items := make([]item.Item, 0)
	for rows.Next() {
		var it item.Item
		err := rows.Scan(
			&it.ID,
			&it.OrderID,
			&it.Status,
			&it.Price,
			&it.ProductID,
			&it.Quantity,
			&it.Observations,
		)
		if err != nil {
			continue
		}

		items = append(items, it)
	}

	return items, nil

}
