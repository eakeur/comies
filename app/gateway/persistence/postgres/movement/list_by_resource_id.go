package movement

import (
	"context"
	"gomies/app/core/entities/movement"
	"gomies/app/gateway/persistence/postgres/query"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (a actions) ListByResourceID(ctx context.Context, resourceID types.ID, filter movement.Filter) ([]movement.Movement, int, error) {
	const script = `
		select
			m.id,
			m.stock_id,
			m.type,
			m.date,
			m.quantity,
			m.value,
			m.agent
		from
			movements m
		inner join
			stocks s on s.id = m.stock_id
		where
			%query%
	`

	q := query.NewQuery(script).
		Where(filter.ResourceID != 0, "s.target_id = $%v", resourceID).And().
		Where(!filter.InitialDate.IsZero(), "m.date >= $%v", filter.InitialDate).And().
		Where(!filter.FinalDate.IsZero(), "m.date <= $%v", filter.FinalDate)

	rows, err := a.db.Query(ctx, q.Script(), q.Args)
	if err != nil {
		return nil, 0, fault.Wrap(err)
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

		movements = append(movements, m)
	}

	return movements, 0, nil
}
