package movement

import (
	"context"
	"gomies/app/core/entities/movement"
	"gomies/app/gateway/persistence/postgres/transaction"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
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

	cmd, err := transaction.ExecFromContext(ctx, script, movement.OutputMovement, agentID)
	if err != nil {
		return fault.Wrap(err)
	}

	if cmd.RowsAffected() != 1 {
		return fault.Wrap(fault.ErrNotFound)
	}

	return nil
}
