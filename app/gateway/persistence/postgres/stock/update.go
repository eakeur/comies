package stock

import (
	"comies/app/core/entities/stock"
	"comies/app/gateway/persistence/postgres/transaction"
	"comies/app/sdk/throw"
	"context"
)

func (a actions) Update(ctx context.Context, st stock.Stock) error {
	const script = `
		update 
			stocks
		set
			minimum_quantity = $1,
			maximum_quantity = $2,
			location = $3
		where 
			target_id = $4
	`

	cmd, err := transaction.ExecFromContext(ctx, script,
		st.MinimumQuantity,
		st.MaximumQuantity,
		st.Location,
		st.TargetID,
	)
	if err != nil {
		return throw.Error(err)
	}

	if cmd.RowsAffected() != 1 {
		return throw.Error(throw.ErrNotFound)
	}

	return nil
}
