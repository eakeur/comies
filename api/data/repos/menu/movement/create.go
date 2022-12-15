package movement

import (
	"comies/core/menu/movement"
	"comies/core/types"
	"comies/data/conn"
	"context"
	"errors"

	"github.com/jackc/pgconn"
)

func (a actions) Create(ctx context.Context, mov movement.Movement) error {
	const script = `
		insert into movements (
			id,
			product_id,
			type,
			date,
			quantity,
			agent_id
		) values (
			$1, $2, $3, $4, $5, $6
		)
	`

	_, err := conn.ExecFromContext(ctx, script,
		mov.ID,
		mov.ProductID,
		mov.Type,
		mov.Date,
		mov.Quantity,
		mov.AgentID,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {

			if conn.IsCode(pgErr, conn.NonexistentFK) && conn.IsConstraint(pgErr, conn.MovementStockIDFK) {
				return types.ErrNotFound
			}

			if conn.IsCode(pgErr, conn.DuplicateError) && conn.IsConstraint(pgErr, conn.MovementIDPK) {
				return types.ErrAlreadyExists
			}
		}
		return err
	}

	return nil
}
