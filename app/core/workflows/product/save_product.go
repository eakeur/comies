package product

import (
	"context"
	"gomies/app/core/entities/catalog/category"
	"gomies/app/core/entities/catalog/product"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/types"
)

func (w workflow) SaveProduct(ctx context.Context, input product.Product, flag ...types.WritingFlag) (product.Product, error) {
	const operation = "Workflows.Product.SaveProduct"

	if err := input.Validate(); err != nil {
		return product.Product{}, fault.Wrap(err, operation)
	}

	// If there is a category external ID assigned to the input, retrieves its internal ID
	if !input.CategoryExternalID.Empty() {
		c, err := w.categories.Get(ctx, category.Key{ID: input.CategoryExternalID})
		if err != nil {
			return product.Product{}, fault.Wrap(err, operation)
		}
		input.CategoryID = c.Entity.ID
	} else {
		input.CategoryID = 0
	}

	input, err := w.products.Save(ctx, input, flag...)
	if err != nil {
		return product.Product{}, fault.Wrap(err, operation)
	}

	return input, nil
}
