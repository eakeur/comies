package item

import (
	"comies/core/types"
	"comies/data/conn"
	"context"
)

func (a actions) SetQuantity(ctx context.Context, itemID types.ID, qt types.Quantity) error {
	const script = `update items set quantity = $1 where id = $2`

	cmd, err := conn.ExecFromContext(ctx, script, qt, itemID)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() <= 0 {
		return types.ErrNotFound
	}

	return nil
}
