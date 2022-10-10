package menu

import (
	"comies/app/core/product"
	"comies/app/data/products"
	"context"
)

func UpdateProduct(ctx context.Context, prd product.Product) error {

	if err := prd.Validate(); err != nil {
		return err
	}

	err := products.Update(ctx, prd)
	if err != nil {
		return err
	}

	return nil
}
