package stock

import (
	"comies/app/core/entities/stock"
	"comies/app/sdk/fault"
	"comies/app/sdk/types"
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
)

func (a actions) GetByID(ctx context.Context, resourceID types.ID) (stock.Stock, error) {
	const script = `
		select
			s.id,
			s.target_id,
			s.minimum_quantity,
			s.maximum_quantity,
			s.location
		from
			stocks s
		where
			s.target_id = $1
	`

	row := a.db.QueryRow(ctx, script, resourceID)

	var st stock.Stock
	if err := row.Scan(
		&st.ID,
		&st.TargetID,
		&st.MinimumQuantity,
		&st.MaximumQuantity,
		&st.Location,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return stock.Stock{}, fault.Wrap(fault.ErrNotFound).
				Describe("the resource id provided seems to not exist").Params(map[string]interface{}{
				"id": resourceID,
			})
		}
		return stock.Stock{}, fault.Wrap(err)
	}

	return st, nil
}
