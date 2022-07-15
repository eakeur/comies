package order

import (
	"comies/app/gateway/persistence/postgres/transaction"
	"comies/app/sdk/fault"
	"comies/app/sdk/types"
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
		return fault.Wrap(err)
	}

	if cmd.RowsAffected() != 1 {
		return fault.Wrap(fault.ErrNotFound)
	}

	return nil
}
