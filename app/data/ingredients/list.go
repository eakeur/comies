package ingredients

import (
	"comies/app/core/id"
	"comies/app/core/ingredient"
	"comies/app/data/conn"
	"context"
)

func List(ctx context.Context, productID id.ID) ([]ingredient.Ingredient, error) {
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
