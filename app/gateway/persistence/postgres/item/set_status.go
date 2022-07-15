package item

import (
	"comies/app/core/entities/item"
	"comies/app/gateway/persistence/postgres/transaction"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (a actions) SetStatus(ctx context.Context, itemID types.ID, status item.Status) error {
	const script = `update items set status = $1 where id = $2`

	cmd, err := transaction.ExecFromContext(ctx, script, status, itemID)
	if err != nil {
		return throw.Error(err)
	}

	if cmd.RowsAffected() <= 0 {
		return throw.Error(throw.ErrNotFound)
	}

	return nil
}
