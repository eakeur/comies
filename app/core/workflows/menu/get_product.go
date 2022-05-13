package menu

import (
	"context"
	product2 "gomies/app/core/entities/product"
	"gomies/app/sdk/fault"
)

func (w workflow) GetProduct(ctx context.Context, ext product2.Key) (product2.Product, error) {
	const operation = "Workflows.Product.GetCategory"

	prod, err := w.products.GetProducts(ctx, ext)
	if err != nil {
		return product2.Product{}, fault.Wrap(err, operation)
	}
	return prod, nil
}
