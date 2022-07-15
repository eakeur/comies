package stock

import (
	"comies/app/core/entities/stock"
	"comies/app/gateway/persistence/postgres"
	"comies/app/gateway/persistence/postgres/transaction"
	"comies/app/sdk/fault"
	"context"
	"errors"
	"github.com/jackc/pgconn"
)

func (a actions) Create(ctx context.Context, st stock.Stock) (stock.Stock, error) {
	const script = `
		insert into stocks (
			id,
			target_id,
			minimum_quantity,
			maximum_quantity,
			location
		) values (
			$1, $2, $3, $4, $5
		)
	`

	if _, err := transaction.ExecFromContext(ctx, script,
		st.ID,
		st.TargetID,
		st.MinimumQuantity,
		st.MaximumQuantity,
		st.Location,
	); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == postgres.DuplicateError && pgErr.ConstraintName == postgres.StockIDPK {
				return stock.Stock{}, fault.Wrap(fault.ErrAlreadyExists).
					Describe("the stock id provided seems to already exist").Params(map[string]interface{}{
					"id": st.ID,
				})
			}

			if pgErr.Code == postgres.DuplicateError && pgErr.ConstraintName == postgres.StockIDUK {
				return stock.Stock{}, fault.Wrap(fault.ErrAlreadyExists).
					Describe("the stock target id provided seems to already exist").Params(map[string]interface{}{
					"target_id": st.TargetID,
				})
			}
		}

		return stock.Stock{}, err
	}

	return st, nil
}
