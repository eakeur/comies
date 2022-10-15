package ingredients

import (
	"comies/app/core/id"
	"comies/app/core/menu"
	"comies/app/core/types"
	"comies/app/data/conn"
	"context"
)

func List(ctx context.Context, productID id.ID) ([]menu.Ingredient, error) {
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

	ingredients := make([]menu.Ingredient, 0)
	for rows.Next() {
		var i menu.Ingredient
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
