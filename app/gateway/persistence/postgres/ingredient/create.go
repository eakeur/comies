package ingredient

import (
	"context"
	"gomies/app/core/entities/ingredient"
	"gomies/app/gateway/persistence/postgres/transaction"
	"gomies/app/sdk/fault"
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
		return ingredient.Ingredient{}, fault.Wrap(err)
	}

	return i, nil
}
