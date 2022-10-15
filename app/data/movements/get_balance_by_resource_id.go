package movements

import (
	"comies/app/core/menu"
	"comies/app/core/types"
	"comies/app/data/conn"
	"comies/app/data/query"
	"context"
)

func GetBalance(ctx context.Context, filter menu.MovementFilter) (types.Quantity, error) {
	const script = `
		select
			coalesce(sum(
            case
                when m.type = 10
                    then m.quantity
                    else  -1 * m.quantity
                end
            ), 0)
		from
			movements m
		where
			%query%
	`

	q, err := query.NewQuery(script).
		Where(!filter.InitialDate.IsZero(), "m.date >= $%v", filter.InitialDate).And().
		Where(!filter.FinalDate.IsZero(), "m.date <= $%v", filter.FinalDate).And().
		OnlyWhere(filter.ProductID != 0, "m.product_id = $%v", filter.ProductID)
	if err != nil {
		return 0, err
	}

	row, err := conn.QueryRowFromContext(ctx, q.Script(), q.Args...)

	var quantity types.Quantity
	if err := row.Scan(&quantity); err != nil {
		return 0, err
	}

	return quantity, nil
}
