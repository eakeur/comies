package movement

import (
	"comies/core/menu/movement"
	"comies/core/types"
	"comies/data/conn"
	"comies/data/query"
	"context"
)

func (a actions) Balance(ctx context.Context, filter movement.Filter) (qt types.Quantity, err error) {
	const script = `
		select coalesce(sum(quantity),0) from movements where %query%
	`

	q, err := query.NewQuery(script).
		Where(!filter.InitialDate.IsZero(), "date >= $%v", filter.InitialDate).And().
		Where(!filter.FinalDate.IsZero(), "date <= $%v", filter.FinalDate).And().
		OnlyWhere(filter.ProductID != 0, "product_id = $%v", filter.ProductID)
	if err != nil {
		return 0, err
	}

	row, err := conn.QueryRowFromContext(ctx, q.Script(), q.Args...)
	if err != nil {
		return 0, err
	}

	return qt, row.Scan(&qt)
}
