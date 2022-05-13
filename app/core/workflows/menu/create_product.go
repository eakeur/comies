package menu

import (
	"context"
	"gomies/app/core/entities/product"
	"gomies/app/sdk/fault"
)

func (w workflow) CreateProduct(ctx context.Context, prd product.Product) (product.Product, error) {
	const operation = "Workflows.Product.CreateProduct"

	if err := prd.Validate(); err != nil {
		return product.Product{}, fault.Wrap(err, operation, fault.AdditionalData{
			"minimum_quantity": prd.MinimumSale,
			"cost_price":       prd.CostPrice,
			"code":             prd.Code,
			"name":             prd.Name,
			"type":             prd.Type,
		})
	}

	prd, err := w.products.CreateProduct(ctx, prd)
	if err != nil {
		return product.Product{}, fault.Wrap(err, operation)
	}

	return prd, nil
}
