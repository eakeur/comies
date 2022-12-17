package ingredient

import (
	"comies/core/menu/ingredient"
	"comies/core/types"
	"comies/data/conn"
	"context"
)

func (a actions) ListByProductID(ctx context.Context, productID types.ID) ([]ingredient.Ingredient, error) {
	const script = `
		select
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
		func(scan conn.Scan, i *ingredient.Ingredient) error {
			return scan(
				&i.ProductID,
				&i.IngredientID,
				&i.Quantity,
				&i.Optional,
			)
		},
	)
}
