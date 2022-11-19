package ingredient

import (
	"comies/app/core/menu/ingredient"
	"comies/app/gateway/data/postgres/conn"
	"context"
)

func (a actions) Create(ctx context.Context, i ingredient.Ingredient) error {
	const script = `
		insert into ingredients (
			product_id,
			ingredient_id,
			quantity,
			optional
		) values (
			$1, $2, $3, $4
		)
	`
	if _, err := conn.ExecFromContext(ctx, script,
		i.ProductID,
		i.IngredientID,
		i.Quantity,
		i.Optional,
	); err != nil {
		// TODO: validate Postgres constraint errors
		return err
	}

	return nil
}
