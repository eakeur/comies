package movements

import (
	"comies/app/core/movement"
	"comies/app/core/types"
	"comies/app/data/conn"
	"context"
	"errors"

	"github.com/jackc/pgconn"
)

func Create(ctx context.Context, mov movement.Movement) (movement.Movement, error) {
	const script = `
		insert into movements (
			id,
			product_id,
			type,
			date,
			quantity,
			value,
			agent_id
		) values (
			$1, $2, $3, $4, $5, $6, $7
		)
	`

	_, err := conn.ExecFromContext(ctx, script,
		mov.ID,
		mov.ProductID,
		mov.Type,
		mov.Date,
		mov.Quantity,
		mov.PaidValue,
		mov.AgentID,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {

			if pgErr.Code == conn.NonexistentFK && pgErr.ConstraintName == conn.MovementStockIDFK {
				return movement.Movement{}, types.ErrNotFound
			}

			if pgErr.Code == conn.DuplicateError && pgErr.ConstraintName == conn.MovementIDPK {
				return movement.Movement{}, types.ErrAlreadyExists
			}
		}
		return movement.Movement{}, err
	}

	return mov, nil
}
