package products

import (
	"comies/app/core/menu"
	"comies/app/core/types"
	"comies/app/data/conn"
	"context"
)

func Remove(ctx context.Context, id types.ID) error {
	const script = `delete from products p where p.id = $1`

	cmd, err := conn.ExecFromContext(ctx, script, id)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() != 1 {
		return menu.ErrNotFound
	}

	return nil
}
