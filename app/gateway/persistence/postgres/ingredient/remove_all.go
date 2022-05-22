package ingredient

import (
	"context"
	"gomies/app/gateway/persistence/postgres/transaction"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (a actions) RemoveAll(ctx context.Context, productID types.ID) error {
	const script = `delete from ingredients where product_id = $1`

	cmd, err := transaction.ExecFromContext(ctx, script, productID)
	if err != nil {
		return fault.Wrap(err)
	}

	if cmd.RowsAffected() <= 0 {
		return fault.Wrap(fault.ErrNotFound)
	}

	return nil
}
