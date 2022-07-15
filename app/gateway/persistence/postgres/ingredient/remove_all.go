package ingredient

import (
	"comies/app/gateway/persistence/postgres/transaction"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (a actions) RemoveAll(ctx context.Context, productID types.ID) error {
	const script = `delete from ingredients where product_id = $1`

	cmd, err := transaction.ExecFromContext(ctx, script, productID)
	if err != nil {
		return throw.Error(err)
	}

	if cmd.RowsAffected() <= 0 {
		return throw.Error(throw.ErrNotFound)
	}

	return nil
}
