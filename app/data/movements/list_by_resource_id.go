package movements

import (
	"comies/app/core/id"
	"comies/app/core/movement"
	"comies/app/data/conn"
	"comies/app/data/query"
	"context"
)

func ListByProductID(ctx context.Context, resourceID id.ID, filter movement.Filter) ([]movement.Movement, error) {
	const script = `
		select
			m.id,
			m.product_id,
			m.type,
			m.date,
			m.quantity,
			m.value,
			m.agent_id
		from
			movements m
		where
			%query%
	`

	q, err := query.NewQuery(script).
		Where(!filter.InitialDate.IsZero(), "m.date >= $%v", filter.InitialDate).And().
		Where(!filter.FinalDate.IsZero(), "m.date <= $%v", filter.FinalDate).And().
		OnlyWhere(resourceID != 0, "m.product_id= $%v", resourceID)

	rows, err := conn.QueryFromContext(ctx, q.Script(), q.Args...)
	if err != nil {
		return nil, err
	}

	movements := make([]movement.Movement, 0)
	for rows.Next() {
		var m movement.Movement
		if err := rows.Scan(
			&m.ID,
			&m.ProductID,
			&m.Type,
			&m.Date,
			&m.Quantity,
			&m.PaidValue,
			&m.AgentID,
		); err != nil {
			continue
		}

		m.Date = m.Date.UTC()

		movements = append(movements, m)
	}

	return movements, nil
}
