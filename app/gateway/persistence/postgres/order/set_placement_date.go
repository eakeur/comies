package order

import (
	"comies/app/gateway/persistence/postgres/transaction"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
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
		return throw.Error(err)
	}

	if cmd.RowsAffected() != 1 {
		return throw.Error(throw.ErrNotFound)
	}

	return nil
}
