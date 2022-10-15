package ingredients

import (
	"comies/app/core/menu"
	"comies/app/core/types"
	"comies/app/data/conn"
	"context"
	"errors"

	"github.com/jackc/pgconn"
)

func Create(ctx context.Context, i menu.Ingredient) error {
	const script = `
		insert into ingredients (
			id,
			product_id,
			ingredient_id,
			quantity,
			optional
		) values (
			$1, $2, $3, $4, $5
		)
	`

	_, err := conn.ExecFromContext(ctx, script,
		i.ID,
		i.ProductID,
		i.IngredientID,
		i.Quantity,
		i.Optional,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {

			if pgErr.Code == conn.NonexistentFK && pgErr.ConstraintName == conn.IngredientProductIDFK {
				return types.ErrNotFound
			}
			if pgErr.Code == conn.NonexistentFK && pgErr.ConstraintName == conn.IngredientIDFK {
				return types.ErrNotFound
			}
			if pgErr.Code == conn.DuplicateError && pgErr.ConstraintName == conn.IngredientProductUK {
				return types.ErrAlreadyExists
			}
		}

		return err
	}

	return nil
}
