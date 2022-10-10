package ingredients

import (
	"comies/app/core/id"
	"comies/app/core/types"
	"comies/app/data/conn"
	"context"
)

func Remove(ctx context.Context, ingredientID id.ID) error {
	const script = `delete from ingredients where id = $1`

	cmd, err := conn.ExecFromContext(ctx, script, ingredientID)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() != 1 {
		return types.ErrNotFound
	}

	return nil
}
