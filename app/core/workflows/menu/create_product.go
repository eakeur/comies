package menu

import (
	"comies/app/core/entities/product"
	"comies/app/core/throw"
	"context"
)

func (w workflow) CreateProduct(ctx context.Context, prd product.Product) (product.Product, error) {

	w.id.Create(&prd.ID)

	if err := prd.Validate(); err != nil {
		return product.Product{}, throw.Error(err).Params(map[string]interface{}{
			"minimum_quantity": prd.MinimumSale,
			"cost_price":       prd.CostPrice,
			"code":             prd.Code,
			"name":             prd.Name,
			"type":             prd.Type,
		})
	}

	prd, err := w.products.Create(ctx, prd)
	if err != nil {
		return product.Product{}, throw.Error(err)
	}

	return prd, nil
}
