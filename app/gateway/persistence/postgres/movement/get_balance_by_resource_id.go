package movement

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"gomies/app/core/entities/movement"
	"gomies/app/gateway/persistence/postgres/query"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (a actions) GetBalanceByResourceID(ctx context.Context, resourceID types.ID, filter movement.Filter) (types.Quantity, error) {
	const script = `
		select
			sum(m.quantity)
		from
			movements m
		inner join
			stocks s on s.id = m.stock_id
		where
			%query%"
	`

	q, err := query.NewQuery(script).
		Where(!filter.InitialDate.IsZero(), "m.date >= $%v", filter.InitialDate).And().
		Where(!filter.FinalDate.IsZero(), "m.date <= $%v", filter.FinalDate).And().
		OnlyWhere(filter.ResourceID != 0, "s.target_id = $%v", resourceID)
	if err != nil {
		return 0, fault.Wrap(err)
	}

	row := a.db.QueryRow(ctx, q.Script(), q.Args)

	var quantity types.Quantity
	if err := row.Scan(&quantity); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, fault.Wrap(fault.ErrNotFound).
				Describe("the resource provided seems to not exist or there are no data for it").Params(map[string]interface{}{
				"resource_id":  filter.ResourceID,
				"initial_date": filter.InitialDate,
				"final_date":   filter.FinalDate,
			})
		}

		return 0, fault.Wrap(err)
	}

	return quantity, nil
}
