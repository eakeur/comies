package product

import (
	"context"
	"gomies/app/core/entities/catalog/product"
	"gomies/app/sdk/fault"
)

func (w workflow) GetProduct(ctx context.Context, ext product.Key) (product.Product, error) {
	const operation = "Workflows.Product.GetCategory"

	prod, err := w.products.GetProducts(ctx, ext)
	if err != nil {
		return product.Product{}, fault.Wrap(err, operation)
	}
	return prod, nil
}
