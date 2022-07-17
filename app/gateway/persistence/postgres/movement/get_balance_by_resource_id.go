package movement

import (
	"comies/app/core/entities/movement"
	"comies/app/gateway/persistence/postgres/query"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (a actions) GetBalanceByProductID(ctx context.Context, resourceID types.ID, filter movement.Filter) (types.Quantity, error) {
	const script = `
		select
			coalesce(sum(m.quantity), 0)
		from
			movements m
		inner join
			products p on p.id = m.product_id
		where
			%query%
	`

	q, err := query.NewQuery(script).
		Where(!filter.InitialDate.IsZero(), "m.date >= $%v", filter.InitialDate).And().
		Where(!filter.FinalDate.IsZero(), "m.date <= $%v", filter.FinalDate).And().
		OnlyWhere(resourceID != 0, "p.id = $%v", resourceID)
	if err != nil {
		return 0, throw.Error(err)
	}

	row := a.db.QueryRow(ctx, q.Script(), q.Args...)

	var quantity types.Quantity
	if err := row.Scan(&quantity); err != nil {
		return 0, throw.Error(err)
	}

	return quantity, nil
}
