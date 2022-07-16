package product

import (
	"comies/app/core/entities/product"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
)

func (a actions) GetSaleInfoByID(ctx context.Context, productID types.ID) (product.Sale, error) {
	const script = `
		select
			cost_price,
			sale_price,
			sale_unit,
			minimum_sale
		from
			products p
		where
			p.id = $1
	`

	row := a.db.QueryRow(ctx, script, productID)

	var p product.Sale
	if err := row.Scan(
		&p.CostPrice,
		&p.SalePrice,
		&p.SaleUnit,
		&p.MinimumSale,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return product.Sale{}, throw.Error(product.ErrNotFound).
				Describe("the product id provided seems to not exist").Params(map[string]interface{}{
				"productID": productID,
			})
		}
		return product.Sale{}, throw.Error(err)
	}

	return p, nil
}
