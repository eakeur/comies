package ingredient

import (
	"comies/app/gateway/persistence/postgres/transaction"
	"comies/app/sdk/fault"
	"comies/app/sdk/types"
	"context"
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
