package ingredients

import (
	"comies/app/core/ingredient"
	"comies/app/core/types"
	"comies/app/data/conn"
	"context"
)

func List(ctx context.Context, productID types.ID) ([]ingredient.Ingredient, error) {
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

	rows, err := conn.QueryFromContext(ctx, script, productID)
	if err != nil {
		return nil, err
	}

	ingredients := make([]ingredient.Ingredient, 0)
	for rows.Next() {
		var i ingredient.Ingredient
		if err := rows.Scan(
			&i.ID,
			&i.ProductID,
			&i.IngredientID,
			&i.Quantity,
			&i.Optional,
		); err != nil {
			return nil, err
		}

		ingredients = append(ingredients, i)
	}

	return ingredients, nil
}
