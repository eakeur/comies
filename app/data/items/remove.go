package items

import (
	"comies/app/core/id"
	"comies/app/core/types"
	"comies/app/data/conn"
	"context"
)

func Remove(ctx context.Context, itemID id.ID) error {
	const script = `delete from items where id = $1`

	cmd, err := conn.ExecFromContext(ctx, script, itemID)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() != 1 {
		return types.ErrNotFound
	}

	return nil
}
