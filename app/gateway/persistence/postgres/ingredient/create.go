package ingredient

import (
	"comies/app/core/entities/ingredient"
	"comies/app/gateway/persistence/postgres"
	"comies/app/gateway/persistence/postgres/transaction"
	"comies/app/sdk/fault"
	"context"
	"errors"

	"github.com/jackc/pgconn"
)

func (a actions) Create(ctx context.Context, i ingredient.Ingredient) (ingredient.Ingredient, error) {
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

	_, err := transaction.ExecFromContext(ctx, script,
		i.ID,
		i.ProductID,
		i.IngredientID,
		i.Quantity,
		i.Optional,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			params := map[string]interface{}{
				"product_id": i.ProductID, "ingredient_id": i.IngredientID, "quantity": i.Quantity.String(),
			}

			if pgErr.Code == postgres.NonexistentFK && pgErr.ConstraintName == postgres.IngredientProductIDFK {
				return ingredient.Ingredient{}, fault.Wrap(fault.ErrNotFound).
					Describe("the product id provided in the product id field seems to not exist").Params(params)
			}
			if pgErr.Code == postgres.NonexistentFK && pgErr.ConstraintName == postgres.IngredientIDFK {
				return ingredient.Ingredient{}, fault.Wrap(fault.ErrNotFound).
					Describe("the product id provided in the ingredient id field seems to not exist").Params(params)
			}
			if pgErr.Code == postgres.DuplicateError && pgErr.ConstraintName == postgres.IngredientProductUK {
				return ingredient.Ingredient{}, fault.Wrap(fault.ErrAlreadyExists).
					Describe("the ingredient provided seems to already exist").Params(params)
			}
		}
		return ingredient.Ingredient{}, fault.Wrap(err)
	}

	return i, nil
}
