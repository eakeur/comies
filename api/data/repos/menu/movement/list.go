package movement

import (
	"comies/core/menu/movement"
	"comies/data/conn"
	"comies/data/query"
	"context"
)

func (a actions) List(ctx context.Context, filter movement.Filter) ([]movement.Movement, error) {
	const script = `
		select
			m.id,
			m.product_id,
			m.type,
			m.date,
			m.quantity,
			m.agent_id
		from
			movements m
		where
			%query%
	`

	q, err := query.NewQuery(script).
		Where(!filter.InitialDate.IsZero(), "m.date >= $%v", filter.InitialDate).And().
		Where(!filter.FinalDate.IsZero(), "m.date <= $%v", filter.FinalDate).And().
		OnlyWhere(filter.ProductID != 0, "m.product_id= $%v", filter.ProductID)
	if err != nil {
		return nil, err
	}

	rows, err := conn.QueryFromContext(ctx, q.Script(), q.Args...)
	if err != nil {
		return nil, err
	}

	return conn.ScanRows(rows,
		func(scan conn.Scan, m movement.Movement) error {
			return scan(
				&m.ID,
				&m.ProductID,
				&m.Type,
				&m.Date,
				&m.Quantity,
				&m.AgentID,
			)
		},
	)
}
