package item

import (
	"context"
	"gomies/app/core/entities/item"
	"gomies/app/gateway/persistence/postgres/transaction"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/session"
	"gomies/app/sdk/types"
)

func (a actions) SetStatus(ctx context.Context, itemID types.ID, status item.Status) error {
	const script = `update items set status = $1 where id = $2 and store_id = $3`

	s, err := session.FromContext(ctx)
	if err != nil {
		return fault.Wrap(err)
	}

	cmd, err := transaction.ExecFromContext(ctx, script, status, itemID, s.StoreID)
	if err != nil {
		return fault.Wrap(err)
	}

	if cmd.RowsAffected() <= 0 {
		return fault.Wrap(fault.ErrNotFound)
	}

	return nil
}
