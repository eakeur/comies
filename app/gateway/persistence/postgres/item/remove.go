package item

import (
	"comies/app/core/throw"
	"comies/app/core/types"
	"comies/app/gateway/persistence/postgres/transaction"
	"context"
)

func (a actions) Remove(ctx context.Context, itemID types.ID) error {
	const script = `delete from items where id = $1`

	cmd, err := transaction.ExecFromContext(ctx, script, itemID)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() != 1 {
		return throw.ErrNotFound
	}

	return nil
}
