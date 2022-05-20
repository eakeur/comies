package item

import (
	"context"
	"gomies/app/core/entities/item"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/session"
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
			observations,
			store_id
		from
			items
		where
			order_id = $1 and
			store_id = $2
	`

	s, err := session.FromContext(ctx)
	if err != nil {
		return nil, fault.Wrap(err)
	}

	rows, err := a.db.Query(ctx, script, orderID, s.StoreID)
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
			&it.Store.StoreID,
		)
		if err != nil {
			continue
		}

		items = append(items, it)
	}

	return items, nil

}
