package ingredient

import (
	"comies/app/core/entities/ingredient"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (a actions) List(ctx context.Context, productID types.ID) ([]ingredient.Ingredient, error) {
	const script = `
		select
			id,
			product_id,
			ingredient_id,
			quantity,
			optional
		from
			ingredients
		where
			product_id = $1
	`

	rows, err := a.db.Query(ctx, script, productID)
	if err != nil {
		return nil, throw.Error(err)
	}

	ingredients := make([]ingredient.Ingredient, 0)
	for rows.Next() {
		var it ingredient.Ingredient
		err := rows.Scan(
			&it.ID,
			&it.ProductID,
			&it.IngredientID,
			&it.Quantity,
			&it.Optional,
		)
		if err != nil {
			continue
		}

		ingredients = append(ingredients, it)
	}

	return ingredients, nil
}
