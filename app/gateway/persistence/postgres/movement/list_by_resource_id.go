package movement

import (
	"comies/app/core/entities/movement"
	"comies/app/gateway/persistence/postgres/query"
	"comies/app/sdk/fault"
	"comies/app/sdk/types"
	"context"
)

func (a actions) ListByResourceID(ctx context.Context, resourceID types.ID, filter movement.Filter) ([]movement.Movement, error) {
	const script = `
		select
			m.id,
			m.stock_id,
			m.type,
			m.date,
			m.quantity,
			m.value,
			m.agent_id
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

	rows, err := a.db.Query(ctx, q.Script(), q.Args...)
	if err != nil {
		return nil, fault.Wrap(err)
	}

	movements := make([]movement.Movement, 0)
	for rows.Next() {
		var m movement.Movement
		if err := rows.Scan(
			&m.ID,
			&m.StockID,
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
