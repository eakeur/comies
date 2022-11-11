package ingredient

import (
	"comies/app/core/menu/ingredient"
	"comies/app/core/types"
	"comies/app/gateway/data/postgres/conn"
	"context"
)

func (a actions) ListByProductID(ctx context.Context, productID types.ID) ([]ingredient.Ingredient, error) {
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

	return conn.ScanRows(rows,
		func(scan conn.Scan, i ingredient.Ingredient) error {
			return scan(
				&i.ID,
				&i.ProductID,
				&i.IngredientID,
				&i.Quantity,
				&i.Optional,
			)
		},
	)
}
