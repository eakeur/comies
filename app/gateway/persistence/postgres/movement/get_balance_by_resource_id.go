package movement

import (
	"context"
	"gomies/app/core/entities/movement"
	"gomies/app/gateway/persistence/postgres/query"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (a actions) GetBalanceByResourceID(ctx context.Context, resourceID types.ID, filter movement.Filter) (types.Quantity, error) {
	const script = `
		select
			coalesce(sum(m.quantity), 0)
		from
			movements m
		inner join
			stocks s on s.id = m.stock_id
		where
			%query%
	`

	q, err := query.NewQuery(script).
		Where(!filter.InitialDate.IsZero(), "m.date >= $%v", filter.InitialDate).And().
		Where(!filter.FinalDate.IsZero(), "m.date <= $%v", filter.FinalDate).And().
		OnlyWhere(resourceID != 0, "s.target_id = $%v", resourceID)
	if err != nil {
		return 0, fault.Wrap(err)
	}

	row := a.db.QueryRow(ctx, q.Script(), q.Args...)

	var quantity types.Quantity
	if err := row.Scan(&quantity); err != nil {
		return 0, fault.Wrap(err)
	}

	return quantity, nil
}
