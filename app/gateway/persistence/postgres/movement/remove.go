package movement

import (
	"context"
	"gomies/app/gateway/persistence/postgres/transaction"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (a actions) Remove(ctx context.Context, movementID types.ID) error {
	const script = `delete from movements m where m.id = $1`

	cmd, err := transaction.ExecFromContext(ctx, script, movementID)
	if err != nil {
		return fault.Wrap(err)
	}

	if cmd.RowsAffected() != 1 {
		return fault.Wrap(fault.ErrNotFound)
	}

	return nil
}
