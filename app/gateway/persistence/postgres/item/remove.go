package item

import (
	"context"
	"gomies/app/gateway/persistence/postgres/transaction"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/session"
	"gomies/app/sdk/types"
)

func (a actions) Remove(ctx context.Context, itemID types.ID) error {
	const script = `delete from items where id = $1 and store_id = $2`

	s, err := session.FromContext(ctx)
	if err != nil {
		return fault.Wrap(err)
	}

	cmd, err := transaction.ExecFromContext(ctx, script, itemID, s.StoreID)
	if err != nil {
		return fault.Wrap(err)
	}

	if cmd.RowsAffected() <= 0 {
		return fault.Wrap(fault.ErrNotFound)
	}

	return nil
}
