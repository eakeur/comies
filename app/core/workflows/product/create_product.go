package product

import (
	"context"
	"gomies/app/core/entities/catalog/product"
	"gomies/app/sdk/fault"
)

func (w workflow) CreateProduct(ctx context.Context, input product.Product) (product.Product, error) {
	const operation = "Workflows.Product.CreateProduct"

	if err := input.Validate(); err != nil {
		return product.Product{}, fault.Wrap(err, operation)
	}

	input, err := w.products.CreateProduct(ctx, input)
	if err != nil {
		return product.Product{}, fault.Wrap(err, operation)
	}

	return input, nil
}
