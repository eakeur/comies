package item

import (
	"comies/app/core/entities/item"
	"comies/app/core/throw"
	"comies/app/core/types"
	"comies/app/gateway/persistence/postgres/transaction"
	"context"
)

func (a actions) SetStatus(ctx context.Context, itemID types.ID, status item.Status) error {
	const script = `update items set status = $1 where id = $2`

	cmd, err := transaction.ExecFromContext(ctx, script, status, itemID)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() <= 0 {
		return throw.ErrNotFound
	}

	return nil
}
