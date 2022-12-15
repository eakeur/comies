package item

import (
	"comies/core/types"
	"comies/io/data/conn"
	"context"
)

func (a actions) SumByBillID(ctx context.Context, billID types.ID) (types.Currency, error) {
	const script = `
		select
			sum(i.credits + i.debts)
		from
			bill_items i
		where bill_id = $1
		group by bill_id
	`

	row, err := conn.QueryRowFromContext(ctx, script, billID)
	if err != nil {
		return 0, err
	}

	var sum types.Currency
	if err := row.Scan(&sum); err != nil {
		return 0, err
	}

	return sum, nil
}
