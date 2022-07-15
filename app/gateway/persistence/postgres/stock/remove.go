package stock

import (
	"comies/app/gateway/persistence/postgres/transaction"
	"comies/app/sdk/fault"
	"comies/app/sdk/types"
	"context"
)

func (a actions) Remove(ctx context.Context, resourceID types.ID) error {
	const script = `delete from stocks s where s.target_id = $1`

	cmd, err := transaction.ExecFromContext(ctx, script, resourceID)
	if err != nil {
		return fault.Wrap(err)
	}

	if cmd.RowsAffected() != 1 {
		return fault.Wrap(fault.ErrNotFound)
	}

	return nil
}
