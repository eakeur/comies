package items

import (
	"comies/app/core/id"
	"comies/app/core/ordering"
	"comies/app/core/types"
	"comies/app/data/conn"
	"context"
)

func SetStatus(ctx context.Context, itemID id.ID, status ordering.Status) error {
	const script = `update items set status = $1 where id = $2`

	cmd, err := conn.ExecFromContext(ctx, script, status, itemID)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() <= 0 {
		return types.ErrNotFound
	}

	return nil
}
