package ingredient

import (
	"context"
	"gomies/app/core/entities/ingredient"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
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
		return nil, fault.Wrap(err)
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
