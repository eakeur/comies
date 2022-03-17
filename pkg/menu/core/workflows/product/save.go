package product

import (
	"context"
	"gomies/pkg/menu/core/entities/category"
	"gomies/pkg/menu/core/entities/product"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/session"
	"gomies/pkg/sdk/types"
)

func (w workflow) Save(ctx context.Context, input product.Product, flag ...types.WritingFlag) (product.Product, error) {
	const operation = "Workflows.Product.Save"
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
	if input.CategoryID == 0 && input.CategoryExternalID != types.Nil && (input.CategoryExternalID != types.External{}) {
		c, err := w.categories.Get(ctx, category.Key{ID: input.CategoryExternalID})
		if err != nil {
			return product.Product{}, fault.Wrap(err, operation)
		}
		input.CategoryID = c.Entity.ID
	}

	input, err = w.products.Save(ctx, input, flag...)
	if err != nil {
		return product.Product{}, fault.Wrap(err, operation)
	}

	return input, nil
}
