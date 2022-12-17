package item

import (
	"comies/core/types"
	"comies/data/conn"
	"context"
)

func (a actions) SumByBillID(ctx context.Context, billID types.ID) (types.Amount, error) {
	const script = `
		select
			sum(i.unit_price * i.quantity) as value,
			sum(i.discounts) as discounts,
			sum(i.unit_price * i.quantity) - sum(i.discounts) as net
		from
			bill_items i
		where bill_id = $1
		group by bill_id
	`

	row, err := conn.QueryRowFromContext(ctx, script, billID)
	if err != nil {
		return types.Amount{}, err
	}

	var amount types.Amount
	if err := row.Scan(&amount.Value, &amount.Discounts, &amount.Net); err != nil {
		return types.Amount{}, err
	}

	return amount, nil
}
