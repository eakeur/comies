package movements

import (
	"comies/app/core/movement"
	"comies/app/data/conn"
	"comies/app/data/query"
	"context"
)

func List(ctx context.Context, filter movement.Filter) ([]movement.Movement, error) {
	const script = `
		select
			m.id,
			m.product_id,
			m.type,
			m.date,
			m.agent_id,
			m.quantity
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

	movements := make([]movement.Movement, 0)
	var m movement.Movement

	for rows.Next() {

		if err := rows.Scan(
			&m.ID,
			&m.ProductID,
			&m.Type,
			&m.Date,
			&m.AgentID,
			&m.Quantity,
		); err != nil {
			return nil, err
		}

		m.Date = m.Date.UTC()
		movements = append(movements, m)
	}

	return movements, nil
}
