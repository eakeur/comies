package ingredient

import (
	"comies/app/core/throw"
	"comies/app/core/types"
	"comies/app/gateway/persistence/postgres/transaction"
	"context"
)

func (a actions) Remove(ctx context.Context, ingredientID types.ID) error {
	const script = `delete from ingredients where id = $1`

	cmd, err := transaction.ExecFromContext(ctx, script, ingredientID)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() != 1 {
		return throw.ErrNotFound
	}

	return nil
}
