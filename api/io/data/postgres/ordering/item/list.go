package item

import (
	"comies/core/ordering/item"
	"comies/core/types"
	"comies/io/data/postgres/conn"
	"context"
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

	rows, err := conn.QueryFromContext(ctx, script, orderID)
	if err != nil {
		return nil, err
	}

	return conn.ScanRows(rows,
		func(scan conn.Scan, i item.Item) error {
			return rows.Scan(
				&i.ID,
				&i.OrderID,
				&i.Status,
				&i.Value,
				&i.ProductID,
				&i.Quantity,
				&i.Observations,
			)
		},
	)

}
