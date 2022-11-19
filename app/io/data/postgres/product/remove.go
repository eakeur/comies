package product

import (
	"comies/app/core/menu/product"
	"comies/app/core/types"
	"comies/app/gateway/data/postgres/conn"
	"context"
)

func (a actions) Remove(ctx context.Context, id types.ID) error {
	const script = `delete from products p where p.id = $1`

	cmd, err := conn.ExecFromContext(ctx, script, id)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() != 1 {
		return product.ErrNotFound
	}

	return nil
}
