package order

import (
	"comies/app/core/throw"
	"comies/app/core/types"
	"comies/app/gateway/persistence/postgres/transaction"
	"context"
)

func (a actions) SetObservation(ctx context.Context, id types.ID, observation string) error {
	const script = `
		update
			orders
		set
			observations = $1
		where
			id = $2
	`

	cmd, err := transaction.ExecFromContext(ctx, script, observation, id)
	if err != nil {
		return throw.Error(err)
	}

	if cmd.RowsAffected() != 1 {
		return throw.Error(throw.ErrNotFound)
	}

	return nil
}
