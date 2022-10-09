package movement

import (
	"comies/app/core/entities/movement"
	"comies/app/core/types"
	"comies/app/gateway/persistence/postgres/query"
	"context"
)

func (a actions) GetBalanceByProductID(ctx context.Context, resourceID types.ID, filter movement.Filter) (types.Quantity, error) {
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
		OnlyWhere(resourceID != 0, "m.product_id = $%v", resourceID)
	if err != nil {
		return 0, err
	}

	row := a.db.QueryRow(ctx, q.Script(), q.Args...)

	var quantity types.Quantity
	if err := row.Scan(&quantity); err != nil {
		return 0, err
	}

	return quantity, nil
}
