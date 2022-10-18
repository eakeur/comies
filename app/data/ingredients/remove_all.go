package ingredients

import (
	"comies/app/core/types"
	"comies/app/data/conn"
	"context"
)

func RemoveAll(ctx context.Context, productID types.ID) error {
	const script = `delete from ingredients where product_id = $1`

	cmd, err := conn.ExecFromContext(ctx, script, productID)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() <= 0 {
		return types.ErrNotFound
	}

	return nil
}
