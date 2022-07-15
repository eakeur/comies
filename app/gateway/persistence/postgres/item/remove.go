package item

import (
	"comies/app/gateway/persistence/postgres/transaction"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (a actions) Remove(ctx context.Context, itemID types.ID) error {
	const script = `delete from items where id = $1`

	cmd, err := transaction.ExecFromContext(ctx, script, itemID)
	if err != nil {
		return throw.Error(err)
	}

	if cmd.RowsAffected() != 1 {
		return throw.Error(throw.ErrNotFound)
	}

	return nil
}
