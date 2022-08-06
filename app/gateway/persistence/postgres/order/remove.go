package order

import (
	"comies/app/core/throw"
	"comies/app/core/types"
	"comies/app/gateway/persistence/postgres/transaction"
	"context"
)

func (a actions) Remove(ctx context.Context, id types.ID) error {
	const script = `
		delete 
		from 
			orders
		where
			id = $1
	`

	cmd, err := transaction.ExecFromContext(ctx, script, id)
	if err != nil {
		return throw.Error(err)
	}

	if cmd.RowsAffected() != 1 {
		return throw.Error(throw.ErrNotFound)
	}

	return nil
}
