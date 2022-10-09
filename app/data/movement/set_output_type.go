package movement

import (
	"comies/app/core/entities/movement"
	"comies/app/core/types"
	"comies/app/gateway/persistence/postgres/transaction"
	"context"
)

func (a actions) SetOutputType(ctx context.Context, agentID types.ID) error {
	const script = `
		update 
			movements
		set
			type = $1
		where 
			id = $2
	`

	cmd, err := transaction.ExecFromContext(ctx, script, movement.OutputType, agentID)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() != 1 {
		return types.ErrNotFound
	}

	return nil
}
