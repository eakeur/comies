package item

import (
	"context"
	"gomies/app/core/entities/item"
	"gomies/app/gateway/persistence/postgres/transaction"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (a actions) SetStatus(ctx context.Context, itemID types.ID, status item.Status) error {
	const script = `update items set status = $1 where id = $2`

	cmd, err := transaction.ExecFromContext(ctx, script, status, itemID)
	if err != nil {
		return fault.Wrap(err)
	}

	if cmd.RowsAffected() <= 0 {
		return fault.Wrap(fault.ErrNotFound)
	}

	return nil
}
