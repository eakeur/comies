package ingredient

import (
	"context"
	"gomies/app/gateway/persistence/postgres/transaction"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (a actions) Remove(ctx context.Context, ingredientID types.ID) error {
	const script = `delete from ingredients where id = $1`

	cmd, err := transaction.ExecFromContext(ctx, script, ingredientID)
	if err != nil {
		return fault.Wrap(err)
	}

	if cmd.RowsAffected() != 1 {
		return fault.Wrap(fault.ErrNotFound)
	}

	return nil
}
