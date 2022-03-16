package product

import (
	"context"
	"gomies/app/core/entities/product"
	"gomies/app/core/managers/session"
	"gomies/app/core/types/fault"
	"gomies/app/core/types/id"
)

func (w workflow) Create(ctx context.Context, input product.Product) (product.Product, error) {
	const operation = "Workflows.Product.Create"
	ctx = w.transactions.Begin(ctx)
	defer w.transactions.End(ctx)

	_, err := session.DelegateSessionProps(ctx, operation, &input.Entity)
	if err != nil {
		return product.Product{}, fault.Wrap(err, operation)
	}

	if err := input.Validate(); err != nil {
		return product.Product{}, fault.Wrap(err, operation)
	}

	// If there is a category external ID assigned to the input, retrieves its internal ID
	if input.CategoryID == 0 && input.CategoryExternalID != id.Nil && (input.CategoryExternalID != id.External{}) {
		c, err := w.categories.Get(ctx, input.CategoryExternalID)
		if err != nil {
			return product.Product{}, fault.Wrap(err, operation)
		}
		input.CategoryID = c.ID
	}

	input, err = w.products.Create(ctx, input)
	if err != nil {
		return product.Product{}, fault.Wrap(err, operation)
	}

	return input, nil
}
