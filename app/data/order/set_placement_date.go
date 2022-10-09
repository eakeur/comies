package order

import (
	"comies/app/core/types"
	"comies/app/gateway/persistence/postgres/transaction"
	"context"
	"time"
)

func (a actions) SetPlacementDate(ctx context.Context, id types.ID, date time.Time) error {
	const script = `
		update
			orders
		set
			placed_at = $1
		where
			id = $2
	`

	cmd, err := transaction.ExecFromContext(ctx, script, date, id)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() != 1 {
		return types.ErrNotFound
	}

	return nil
}
