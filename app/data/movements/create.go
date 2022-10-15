package movements

import (
	"comies/app/core/menu"
	"comies/app/core/types"
	"comies/app/data/conn"
	"context"
	"errors"

	"github.com/jackc/pgconn"
)

func Create(ctx context.Context, m menu.Movement) error {
	const script = `
		insert into movements (
			id,
			product_id,
			type,
			date,
			agent_id,
			quantity
		) values (
			$1, $2, $3, $4, $5, $6, $7
		)
	`

	_, err := conn.ExecFromContext(ctx, script,
		m.ID,
		m.ProductID,
		m.Type,
		m.Date,
		m.Date,
		m.AgentID,
		menu.MovementQuantity(m),
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == conn.NonexistentFK && pgErr.ConstraintName == conn.MovementStockIDFK {
				return types.ErrNotFound
			}

			if pgErr.Code == conn.DuplicateError && pgErr.ConstraintName == conn.MovementIDPK {
				return types.ErrAlreadyExists
			}
		}

		return err
	}

	return nil
}
