package ingredient

import (
	"comies/app/core/throw"
	"comies/app/core/types"
	"comies/app/gateway/persistence/postgres/transaction"
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
